package game

import (
	"maps"
	"slices"

	"github.com/google/uuid"
)

type Status string

const (
	StatusNone   Status = "none"
	StatusBurned Status = "burned"
)

type ActorDef struct {
	ID         uuid.UUID
	Name       string
	Affinities map[Affinity]struct{}
	Stats      map[Stat]float64
	Effects    []Effect
}

type actormeta struct {
	active_turns   int
	inactive_turns int
}

type Actor struct {
	ActorDef
	Level      int
	PlayerID   uuid.UUID
	PositionID uuid.UUID

	AffinityDamage     map[Affinity]int
	AffinityResistance map[Affinity]int
	AffinityImmunities map[Affinity]float64
	UnmodifiedStats    map[Stat]float64
	Stages             map[Stat]int
	Aux                map[Stat]float64
	Stats              map[Stat]float64

	Wounds float64
	Status Status

	IsAlive     bool
	IsProtected bool
	IsStaggered bool // cannot act
	IsStunned   bool // cannot act, and cannot queue commands

	meta actormeta
}

func NewActorDef() ActorDef {
	return ActorDef{
		ID:         uuid.New(),
		Name:       "",
		Affinities: map[Affinity]struct{}{},
		Stats: map[Stat]float64{
			Health:         100,
			Speed:          100,
			Melee:          100,
			Ranged:         100,
			Special:        100,
			MartialDefense: 100,
			SpecialDefense: 100,
			Accuracy:       1,
			Evasion:        1,
		},
		Effects: []Effect{},
	}
}

func (d ActorDef) Clone() ActorDef {
	return ActorDef{
		ID:         d.ID,
		Name:       d.Name,
		Affinities: maps.Clone(d.Affinities),
		Stats:      maps.Clone(d.Stats),
		Effects:    slices.Clone(d.Effects),
	}
}

func NewActor(playerID uuid.UUID, def ActorDef) Actor {
	return Actor{
		ActorDef:   def.Clone(),
		PlayerID:   playerID,
		PositionID: uuid.Nil,
		Level:      100,

		AffinityDamage:     map[Affinity]int{},
		AffinityResistance: map[Affinity]int{},
		AffinityImmunities: map[Affinity]float64{},
		Stages:             map[Stat]int{},
		Aux:                map[Stat]float64{},
		Stats:              maps.Clone(def.Stats),

		Wounds: 0,
		Status: StatusNone,

		IsAlive:     true,
		IsProtected: false,
		IsStaggered: false,
		IsStunned:   false,

		meta: actormeta{
			active_turns:   0,
			inactive_turns: 0,
		},
	}
}

func (a Actor) Clone() Actor {
	return Actor{
		ActorDef:   a.ActorDef.Clone(),
		Level:      a.Level,
		PlayerID:   a.PlayerID,
		PositionID: a.PositionID,

		AffinityDamage:     maps.Clone(a.AffinityDamage),
		AffinityResistance: maps.Clone(a.AffinityResistance),
		AffinityImmunities: maps.Clone(a.AffinityImmunities),
		UnmodifiedStats:    maps.Clone(a.UnmodifiedStats),
		Stages:             maps.Clone(a.Stages),
		Aux:                maps.Clone(a.Aux),
		Stats:              maps.Clone(a.Stats),

		Wounds:      a.Wounds,
		Status:      a.Status,
		IsAlive:     a.IsAlive,
		IsProtected: a.IsProtected,
		IsStaggered: a.IsStaggered,
		IsStunned:   a.IsStunned,

		meta: a.meta,
	}
}

func mapBaseStat(actor Actor, stat Stat, stats map[Stat]float64, aux float64) float64 {
	base := stats[stat]*2 + aux
	ratio := float64(actor.Level) / 100
	result := base*ratio + 5
	if stat == Health {
		result += float64(actor.Level)
	}
	return result
}
func (a *Actor) mapBaseStats() {
	a.UnmodifiedStats = maps.Clone(a.Stats)

	for stat, _ := range a.Stats {
		if stat == Accuracy || stat == Evasion {
			continue
		}

		aux, ok := a.Aux[stat]
		if !ok {
			aux = 0
		}

		a.Stats[stat] = mapBaseStat(*a, stat, a.Stats, aux)
		a.UnmodifiedStats[stat] = mapBaseStat(*a, stat, a.UnmodifiedStats, 0)
	}
}

func (a *Actor) ApplyDamage(damage float64, resolved Actor) {
	a.Wounds = a.Wounds + damage
	if a.Wounds < 0 {
		a.Wounds = 0
	}

	a.IsAlive = resolved.Stats[Health] > a.Wounds
}
func (a *Actor) IncrementTurns() {
	if a.IsActive() {
		a.meta.active_turns++
	} else {
		a.meta.inactive_turns++
	}
}

func (a Actor) IsActive() bool {
	return a.PositionID != uuid.Nil
}
func (a Actor) GetAffinityDamage(affinity Affinity) int {
	base := maps.Clone(a.AffinityDamage)
	for affinity := range a.Affinities {
		b, ok := base[affinity]
		if !ok {
			continue
		}

		base[affinity] = b + 1
	}

	value, ok := base[affinity]
	if ok {
		return value
	}

	return 0
}
func (a Actor) GetAffinityResistance(affinity Affinity) int {
	value, ok := a.AffinityResistance[affinity]
	if ok {
		return value
	}

	return 0
}
func (a Actor) GetEffectiveAffinityResistance(affinity Affinity) int {
	return a.GetAffinityResistance(affinity) - affinity.GetBaseModifier(a)
}
func (a Actor) GetRemainingHealth() float64 {
	health := a.Stats[Health]
	return health - a.Wounds
}
func (a Actor) GetModifiers() []Modifier {
	modifiers := []Modifier{}
	for _, effect := range a.Effects {
		if effect.Ready() {
			modifier := effect.Bind(MakeContextFrom(a))
			modifiers = append(modifiers, modifier)
		}
	}

	return modifiers
}
func (a Actor) GetActions() []Action {
	return []Action{}
}
func (a Actor) GetActionByID(action_id uuid.UUID) (Action, bool) {
	return Action{}, false
}

type actorJSON struct {
	ID                 uuid.UUID        `json:"ID"`
	Name               string           `json:"name"`
	Level              int              `json:"level"`
	PlayerID           uuid.UUID        `json:"player_ID"`
	PositionID         *uuid.UUID       `json:"position_ID"`
	Affinities         []Affinity       `json:"affinities"`
	AffinityDamage     map[Affinity]int `json:"affinity_damage"`
	AffinityResistance map[Affinity]int `json:"affinity_resistance"`
	Stats              map[Stat]int     `json:"stats"`
	Stages             map[Stat]int     `json:"stages"`
	UnmodifiedStats    map[Stat]int     `json:"unmodified_stats"`
	AppliedModifiers   []uuid.UUID      `json:"applied_modifiers"`
	Wounds             int              `json:"wounds"`
	Status             Status           `json:"status"`
	IsActive           bool             `json:"is_active"`
	IsAlive            bool             `json:"is_alive"`
	IsProtected        bool             `json:"is_protected"`
	IsStaggered        bool             `json:"is_staggered"`
	IsStunned          bool             `json:"is_stunned"`
}

func (a Actor) ToJSON(g Game) actorJSON {
	stats := make(map[Stat]int, len(a.Stats))
	unmodified_stats := make(map[Stat]int, len(a.UnmodifiedStats))
	applied_modifiers := slices.Collect(maps.Keys(g.AppliedModifiers(a.ID)))

	for stat, v := range a.Stats {
		if stat == Accuracy || stat == Evasion {
			v = v * 100
		}
		stats[stat] = int(v)
	}
	for stat, v := range a.UnmodifiedStats {
		if stat == Accuracy || stat == Evasion {
			v = v * 100
		}
		unmodified_stats[stat] = int(v)
	}

	position_id := &a.PositionID
	if a.PositionID == uuid.Nil {
		position_id = nil
	}

	affinity_resistance := maps.Clone(a.AffinityResistance)
	for affinity := range AFFINITY_MATRIX {
		resistance := a.GetEffectiveAffinityResistance(affinity)
		_, has_explicit_resistance := a.AffinityResistance[affinity]
		if resistance != 0 || has_explicit_resistance {
			affinity_resistance[affinity] = resistance
		} else {
			delete(affinity_resistance, affinity)
		}
	}

	affinity_damage := maps.Clone(a.AffinityDamage)
	for affinity := range a.Affinities {
		base, ok := affinity_damage[affinity]
		if !ok {
			affinity_damage[affinity] = 1
		} else {
			affinity_damage[affinity] = base + 1
		}
	}

	return actorJSON{
		ID:                 a.ID,
		Name:               a.Name,
		Level:              a.Level,
		PlayerID:           a.PlayerID,
		PositionID:         position_id,
		Affinities:         slices.Collect(maps.Keys(a.Affinities)),
		AffinityDamage:     affinity_damage,
		AffinityResistance: affinity_resistance,
		Stats:              stats,
		Stages:             maps.Clone(a.Stages),
		UnmodifiedStats:    unmodified_stats,
		AppliedModifiers:   applied_modifiers,
		Wounds:             int(a.Wounds),
		Status:             a.Status,
		IsActive:           a.IsActive(),
		IsAlive:            a.IsAlive,
		IsProtected:        a.IsProtected,
		IsStaggered:        a.IsStaggered,
		IsStunned:          a.IsStunned,
	}
}
