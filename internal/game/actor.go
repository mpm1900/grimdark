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

	Damage      float64
	Status      Status
	IsAlive     bool
	IsProtected bool

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

		Damage:      0,
		Status:      StatusNone,
		IsAlive:     true,
		IsProtected: false,

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

		Damage:      a.Damage,
		Status:      a.Status,
		IsAlive:     a.IsAlive,
		IsProtected: a.IsProtected,

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
	a.Damage = a.Damage + damage
	if a.Damage < 0 {
		a.Damage = 0
	}

	a.IsAlive = resolved.Stats[Health] > a.Damage
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
			b = 0
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
func (a Actor) GetRemainingHealth() float64 {
	health := a.Stats[Health]
	return health - a.Damage
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
func (a Actor) GetActionByID(action_id uuid.UUID) (Action, bool) {
	return Action{}, false
}

type actorJSON struct {
	ID                 uuid.UUID         `json:"ID"`
	PlayerID           uuid.UUID         `json:"player_ID"`
	PositionID         uuid.UUID         `json:"position_ID"`
	Affinities         []Affinity        `json:"affinities"`
	AffinityDamage     map[Affinity]int  `json:"affinity_damage"`
	AffinityResistance map[Affinity]int  `json:"affinity_resistance"`
	Stats              map[Stat]int      `json:"stats"`
	UnmodifiedStats    map[Stat]int      `json:"unmodified_stats"`
	AppliedEffects     map[uuid.UUID]int `json:"applied_effects"`
	Damage             int               `json:"damage"`
	Status             Status            `json:"status"`
	IsAlive            bool              `json:"is_alive"`
	IsProtected        bool              `json:"is_protected"`
}

func (a Actor) ToJSON(g Game) actorJSON {
	stats := make(map[Stat]int, len(a.Stats))
	unmodified_stats := make(map[Stat]int, len(a.UnmodifiedStats))
	applied_effects := g.AppliedEffects(a.ID)

	for k, v := range a.Stats {
		stats[k] = int(v)
	}
	for k, v := range a.UnmodifiedStats {
		unmodified_stats[k] = int(v)
	}

	return actorJSON{
		ID:                 a.ID,
		PlayerID:           a.PlayerID,
		PositionID:         a.PositionID,
		Affinities:         slices.Collect(maps.Keys(a.Affinities)),
		AffinityDamage:     maps.Clone(a.AffinityDamage),
		AffinityResistance: maps.Clone(a.AffinityResistance),
		Stats:              stats,
		UnmodifiedStats:    unmodified_stats,
		AppliedEffects:     applied_effects,
		Damage:             int(a.Damage),
		Status:             a.Status,
		IsAlive:            a.IsAlive,
		IsProtected:        a.IsProtected,
	}
}
