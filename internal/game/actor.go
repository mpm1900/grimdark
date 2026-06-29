package game

import (
	"maps"
	"slices"

	"github.com/google/uuid"
)

type ActorState string
type ActorStatus string

const (
	StateGrounded ActorState = "grounded"
)

const (
	StatusNone   ActorStatus = "none"
	StatusBurned ActorStatus = "burned"
)

type ActorDef struct {
	ID         uuid.UUID
	Name       string
	Affinities map[Affinity]struct{}
	Stats      map[Stat]float64
	Effects    []Effect
}

type ActorMeta struct {
	Active_turns   int
	Inactive_turns int
}

type Actor struct {
	ActorDef
	Level      int
	PlayerID   uuid.UUID
	PositionID uuid.UUID

	Actions []Action
	Weapon  *Weapon

	AffinityDamage     map[Affinity]int
	AffinityResistance map[Affinity]int
	AffinityImmunities map[Affinity]float64
	UnmodifiedStats    map[Stat]float64
	Stages             map[Stat]int
	AuxStats           map[Stat]float64
	Stats              map[Stat]float64

	Wounds  float64
	Augment Augment
	State   ActorState
	Status  ActorStatus

	IsAlive     bool
	IsHidden    bool // cannot be targeted by single-target actions
	IsProtected bool
	IsStaggered bool // cannot act
	IsStunned   bool // cannot act, and cannot queue commands

	meta ActorMeta
}

type actorJSON struct {
	ID                 uuid.UUID        `json:"ID"`
	Name               string           `json:"name"`
	Level              int              `json:"level"`
	PlayerID           uuid.UUID        `json:"player_ID"`
	PositionID         *uuid.UUID       `json:"position_ID"`
	Actions            []actionJSON     `json:"actions"`
	Weapon             *weaponJSON      `json:"weapon"`
	Effects            []Effect         `json:"effects"`
	Affinities         []Affinity       `json:"affinities"`
	AffinityDamage     map[Affinity]int `json:"affinity_damage"`
	AffinityResistance map[Affinity]int `json:"affinity_resistance"`
	Stats              map[Stat]int     `json:"stats"`
	Stages             map[Stat]int     `json:"stages"`
	UnmodifiedStats    map[Stat]int     `json:"unmodified_stats"`
	ActiveModifiers    []uuid.UUID      `json:"active_modifiers"`
	Wounds             int              `json:"wounds"`
	Augment            Augment          `json:"augment"`
	State              ActorState       `json:"state"`
	Status             ActorStatus      `json:"status"`
	IsActive           bool             `json:"is_active"`
	IsAlive            bool             `json:"is_alive"`
	IsHidden           bool             `json:"is_hidden"`
	IsProtected        bool             `json:"is_protected"`
	IsStaggered        bool             `json:"is_staggered"`
	IsStunned          bool             `json:"is_stunned"`
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

		Actions: []Action{},
		Weapon:  nil,

		AffinityDamage:     map[Affinity]int{},
		AffinityResistance: map[Affinity]int{},
		AffinityImmunities: map[Affinity]float64{},
		Stages:             map[Stat]int{},
		AuxStats:           map[Stat]float64{},
		Stats:              maps.Clone(def.Stats),

		Wounds:  0,
		Augment: AugmentDefault,
		State:   StateGrounded,
		Status:  StatusNone,

		IsAlive:     true,
		IsHidden:    false,
		IsProtected: false,
		IsStaggered: false,
		IsStunned:   false,

		meta: ActorMeta{
			Active_turns:   0,
			Inactive_turns: 0,
		},
	}
}

func (a Actor) Clone() Actor {
	var weapon *Weapon = nil
	if a.Weapon != nil {
		clone := a.Weapon.Clone()
		weapon = &clone
	}
	return Actor{
		ActorDef:   a.ActorDef.Clone(),
		Level:      a.Level,
		PlayerID:   a.PlayerID,
		PositionID: a.PositionID,

		Actions: slices.Clone(a.Actions),
		Weapon:  weapon,

		AffinityDamage:     maps.Clone(a.AffinityDamage),
		AffinityResistance: maps.Clone(a.AffinityResistance),
		AffinityImmunities: maps.Clone(a.AffinityImmunities),
		UnmodifiedStats:    maps.Clone(a.UnmodifiedStats),
		Stages:             maps.Clone(a.Stages),
		AuxStats:           maps.Clone(a.AuxStats),
		Stats:              maps.Clone(a.Stats),

		Wounds:  a.Wounds,
		Augment: a.Augment,
		State:   a.State,
		Status:  a.Status,

		IsAlive:     a.IsAlive,
		IsHidden:    a.IsHidden,
		IsProtected: a.IsProtected,
		IsStaggered: a.IsStaggered,
		IsStunned:   a.IsStunned,

		meta: a.meta,
	}
}

func mapBaseStat(actor Actor, stat Stat, stats map[Stat]float64, aux float64) float64 {
	base := stats[stat]*2 + aux
	ratio := float64(actor.Level) / 100
	result := (base*ratio + 5) * actor.Augment.GetMultiplier(stat)
	if stat == Health {
		result += float64(actor.Level)
	}
	return result
}
func (a *Actor) getAux(stat Stat) float64 {
	aux, ok := a.AuxStats[stat]
	if !ok {
		aux = 0
	}
	if a.Weapon != nil {
		weapon_aux, ok := a.Weapon.AuxStats[stat]
		if ok {
			aux += weapon_aux
		}
	}

	return aux
}
func (a *Actor) mapBaseStats() {
	a.UnmodifiedStats = maps.Clone(a.Stats)

	for stat, _ := range a.Stats {
		// accuracy and evasion are un-mapped
		if stat == Accuracy || stat == Evasion {
			continue
		}

		a.Stats[stat] = mapBaseStat(*a, stat, a.Stats, a.getAux(stat))
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
		a.meta.Active_turns++
	} else {
		a.meta.Inactive_turns++
	}
}
func (a *Actor) SetPosition(position_id uuid.UUID) {
	a.PositionID = position_id
	if position_id == uuid.Nil {
		a.meta.Inactive_turns = 0
	} else {
		a.meta.Active_turns = 0
	}
}

func (a Actor) IsActive() bool {
	return a.PositionID != uuid.Nil
}
func (a Actor) CanAct() bool {
	return !a.IsStaggered && !a.IsStunned
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
func (a Actor) GetEffects() []Effect {
	effects := slices.Clone(a.Effects)
	if a.Weapon != nil {
		effects = append(effects, a.Weapon.Effects...)
	}
	return effects
}
func (a Actor) GetModifiers() []Modifier {
	modifiers := []Modifier{}
	for _, effect := range a.GetEffects() {
		if effect.Ready() {
			modifier := effect.Bind(MakeContextFrom(a))
			modifiers = append(modifiers, modifier)
		}
	}

	return modifiers
}
func (a Actor) GetActions() []Action {
	actions := slices.Clone(a.Actions)
	if a.Weapon != nil {
		actions = append(actions, a.Weapon.Actions...)
	}

	actions = append(actions, GLOBAL_ACTIONS...)
	return slices.DeleteFunc(actions, func(a Action) bool {
		return !a.IsActive
	})
}
func (a Actor) GetActionByID(action_id uuid.UUID) (Action, bool) {
	for _, action := range a.GetActions() {
		if action.ID == action_id {
			return action, true
		}
	}

	return Action{}, false
}
func (a Actor) Targetable() bool {
	if a.IsActive() {
		return !a.IsHidden
	}

	return true
}
func (a Actor) GetMeta() ActorMeta {
	return a.meta
}

func (a Actor) ToJSON(g Game) actorJSON {
	stats := make(map[Stat]int, len(a.Stats))
	unmodified_stats := make(map[Stat]int, len(a.UnmodifiedStats))
	active_modifiers := slices.Collect(maps.Keys(g.AppliedModifiers(a.ID)))

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

	actor_actions := a.GetActions()
	actions := make([]actionJSON, len(actor_actions))
	for i, action := range actor_actions {
		actions[i] = action.ToJSON(g, a)
	}

	var weapon *weaponJSON = nil
	if a.Weapon != nil {
		clone := a.Weapon.Clone().ToJSON(g, a)
		weapon = &clone
	}

	return actorJSON{
		ID:                 a.ID,
		Name:               a.Name,
		Level:              a.Level,
		PlayerID:           a.PlayerID,
		PositionID:         NilifyUUID(a.PositionID),
		Actions:            actions,
		Weapon:             weapon,
		Effects:            a.GetEffects(),
		Affinities:         slices.Collect(maps.Keys(a.Affinities)),
		AffinityDamage:     affinity_damage,
		AffinityResistance: affinity_resistance,
		Stats:              stats,
		Stages:             maps.Clone(a.Stages),
		UnmodifiedStats:    unmodified_stats,
		ActiveModifiers:    active_modifiers,
		Wounds:             int(a.Wounds),
		Augment:            a.Augment,
		State:              a.State,
		Status:             a.Status,
		IsActive:           a.IsActive(),
		IsAlive:            a.IsAlive,
		IsHidden:           a.IsHidden,
		IsProtected:        a.IsProtected,
		IsStaggered:        a.IsStaggered,
		IsStunned:          a.IsStunned,
	}
}
