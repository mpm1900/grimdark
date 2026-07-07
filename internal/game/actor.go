package game

import (
	"maps"
	"slices"

	"github.com/google/uuid"
)

type ActorState string
type ActorStatus string
type ActorRace string
type ActorFaction string

const (
	StateGrounded ActorState = "grounded"
)

const (
	StatusNone   ActorStatus = "none"
	StatusBurned ActorStatus = "burned"
)

const (
	RaceHuman ActorRace = "human"
)

const (
	FactionImperium ActorFaction = "imperium"
)

type ActorDef struct {
	ID         uuid.UUID
	Affinities map[Affinity]struct{}
	Effects    []Effect
	Faction    ActorFaction
	Name       string
	Race       ActorRace
	Stats      map[Stat]float64
}

type ActionState struct {
	Cooldown   int
	IsDisabled bool
}
type ActorMeta struct {
	ActiveTurns      int
	InactiveTurns    int
	LastUsedActionID uuid.UUID
	Seen             bool
}

type Actor struct {
	ActorDef
	Level      int
	PlayerID   uuid.UUID
	PositionID uuid.UUID

	Actions []Action
	Item    *HeldItem
	Weapon  *Weapon

	AffinityDamage     map[Affinity]int
	AffinityResistance map[Affinity]int
	AffinityImmunities map[Affinity]float64
	UnmodifiedStats    map[Stat]float64
	Stages             map[Stat]int
	AuxStats           map[Stat]float64
	Stats              map[Stat]float64

	Wounds float64
	State  ActorState
	Status ActorStatus

	IsAlive     bool
	IsBulwark   bool // stops collateral penetration
	IsHidden    bool // cannot be targeted by single-target actions
	IsInsulated bool // is immune from the secondary effects of attacking attacks (ie through AddResultEffects())
	IsProtected bool // protected from actions that check accuracy
	IsStunned   bool // staggered + and cannot queue commands (may not be needed)

	Meta         ActorMeta
	ActionsState map[uuid.UUID]ActionState
}

type actorJSON struct {
	ID                 uuid.UUID            `json:"ID"`
	Name               string               `json:"name"`
	Race               ActorRace            `json:"race"`
	Faction            ActorFaction         `json:"faction"`
	Level              int                  `json:"level"`
	PlayerID           uuid.UUID            `json:"player_ID"`
	PositionID         *uuid.UUID           `json:"position_ID"`
	Actions            []actionJSON         `json:"actions"`
	Weapon             *weaponJSON          `json:"weapon"`
	Item               *HeldItem            `json:"item"`
	Effects            []Effect             `json:"effects"`
	Affinities         []Affinity           `json:"affinities"`
	AffinityDamage     map[Affinity]int     `json:"affinity_damage"`
	AffinityResistance map[Affinity]int     `json:"affinity_resistance"`
	AffinityImmunities map[Affinity]float64 `json:"affinity_immunities"`
	Stats              map[Stat]int         `json:"stats"`
	Stages             map[Stat]int         `json:"stages"`
	UnmodifiedStats    map[Stat]int         `json:"unmodified_stats"`
	ActiveModifiers    []uuid.UUID          `json:"active_modifiers"`
	Wounds             int                  `json:"wounds"`
	Seen               bool                 `json:"-"`
	State              ActorState           `json:"state"`
	Status             ActorStatus          `json:"status"`
	IsActive           bool                 `json:"is_active"`
	IsBulwark          bool                 `json:"is_bulwark"`
	IsAlive            bool                 `json:"is_alive"`
	IsHidden           bool                 `json:"is_hidden"`
	IsInsulated        bool                 `json:"is_insulated"`
	IsProtected        bool                 `json:"is_protected"`
	IsStunned          bool                 `json:"is_stunned"`
}

func NewActorDef() ActorDef {
	return ActorDef{
		ID:         uuid.New(),
		Name:       "",
		Race:       RaceHuman,
		Faction:    FactionImperium,
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
			CriticalChance: 1,
			CriticalDamage: 1,
			EffectChance:   1,
			Range:          0,
		},
		Effects: []Effect{},
	}
}

func (d ActorDef) Clone() ActorDef {
	return ActorDef{
		ID:         d.ID,
		Affinities: maps.Clone(d.Affinities),
		Effects:    slices.Clone(d.Effects),
		Faction:    d.Faction,
		Name:       d.Name,
		Race:       d.Race,
		Stats:      maps.Clone(d.Stats),
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
		Item:    nil,

		AffinityDamage:     map[Affinity]int{},
		AffinityResistance: map[Affinity]int{},
		AffinityImmunities: map[Affinity]float64{},
		Stages:             map[Stat]int{},
		AuxStats:           map[Stat]float64{},
		Stats:              maps.Clone(def.Stats),

		Wounds: 0,
		State:  StateGrounded,
		Status: StatusNone,

		IsAlive:     true,
		IsBulwark:   false,
		IsHidden:    false,
		IsInsulated: false,
		IsProtected: false,
		IsStunned:   false,

		ActionsState: map[uuid.UUID]ActionState{},

		Meta: ActorMeta{
			ActiveTurns:      0,
			InactiveTurns:    0,
			LastUsedActionID: uuid.Nil,
			Seen:             false,
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
		Item:    a.Item,

		AffinityDamage:     maps.Clone(a.AffinityDamage),
		AffinityResistance: maps.Clone(a.AffinityResistance),
		AffinityImmunities: maps.Clone(a.AffinityImmunities),
		UnmodifiedStats:    maps.Clone(a.UnmodifiedStats),
		Stages:             maps.Clone(a.Stages),
		AuxStats:           maps.Clone(a.AuxStats),
		Stats:              maps.Clone(a.Stats),

		Wounds: a.Wounds,
		State:  a.State,
		Status: a.Status,

		IsAlive:     a.IsAlive,
		IsBulwark:   a.IsBulwark,
		IsHidden:    a.IsHidden,
		IsInsulated: a.IsInsulated,
		IsProtected: a.IsProtected,
		IsStunned:   a.IsStunned,

		ActionsState: maps.Clone(a.ActionsState),

		Meta: a.Meta,
	}
}

func mapBaseStat(actor Actor, stat Stat, stats map[Stat]float64, aux float64) float64 {
	base := stats[stat]*2 + aux
	ratio := float64(actor.Level) / 100
	result := (base*ratio + 5)
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
		if _, ok := mappedStats[stat]; !ok {
			continue
		}

		a.Stats[stat] = mapBaseStat(*a, stat, a.Stats, a.getAux(stat))
		a.UnmodifiedStats[stat] = mapBaseStat(*a, stat, a.UnmodifiedStats, 0)
	}
}
func (a *Actor) mapStagedStats() {
	builder := NewStageBuilder(a.Stages)
	stats := builder.ResolveAll(a.Stats)

	builder.Mod = 3
	accuracy := builder.Resolve(Accuracy, a.Stats[Accuracy])
	evasion := builder.Resolve(Evasion, a.Stats[Evasion])
	crit_damage := builder.Resolve(CriticalDamage, a.Stats[CriticalDamage])
	effect_chance := builder.Resolve(EffectChance, a.Stats[EffectChance])

	a.Stats = stats
	a.Stats[Accuracy] = accuracy
	a.Stats[Evasion] = evasion
	a.Stats[CriticalDamage] = crit_damage
	a.Stats[EffectChance] = effect_chance
}

func (a *Actor) ApplyDamage(damage float64, resolved Actor) {
	a.Wounds = a.Wounds + damage
	if a.Wounds < 0 {
		a.Wounds = 0
	}

	a.IsAlive = resolved.Stats[Health] > a.Wounds
}
func (a *Actor) UpdateActionState(action_id uuid.UUID, updater func(ActionState) ActionState) {
	state, ok := a.ActionsState[action_id]
	if !ok {
		a.ActionsState[action_id] = updater(ActionState{})
	} else {
		a.ActionsState[action_id] = updater(state)
	}
}
func (a *Actor) SetActionCooldown(action_id uuid.UUID, cooldown int) {
	// +1 is for semantics, since cooldowns are decremented at turn end,
	// a one-turn cooldown needs a value of 2.
	a.UpdateActionState(action_id, func(s ActionState) ActionState {
		s.Cooldown = cooldown + 1
		return s
	})
}
func (a *Actor) NextTurn() {
	if a.IsActive() {
		a.Meta.ActiveTurns++
	} else {
		a.Meta.InactiveTurns++
	}

	for aid, state := range a.ActionsState {
		if state.Cooldown > 0 {
			a.UpdateActionState(aid, func(s ActionState) ActionState {
				s.Cooldown--
				return s
			})
		}
	}
}
func (a *Actor) SetPosition(position_id uuid.UUID) {
	if position_id == uuid.Nil {
		a.Meta.InactiveTurns = 0
	} else {
		a.Meta.Seen = true
	}
	if a.PositionID == uuid.Nil {
		a.Meta.ActiveTurns = 0
	}
	a.PositionID = position_id
	a.Meta.LastUsedActionID = uuid.Nil
}

func (a Actor) IsActive() bool {
	return a.PositionID != uuid.Nil
}
func (a Actor) CanAct() bool {
	return !a.IsStunned
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
	if a.Item != nil {
		effects = append(effects, a.Item.Effects...)
	}
	return effects
}
func (a Actor) GetModifiers() []Modifier {
	modifiers := []Modifier{}
	for _, effect := range a.GetEffects() {
		if effect.Ready() {
			context := MakeContextFrom(a)
			context.EffectID = effect.ID
			modifier := effect.Bind(context)
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
	return slices.DeleteFunc(actions, func(action Action) bool {
		if action.ActiveCheck == nil {
			return false
		}

		return !action.ActiveCheck(a)
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
	return a.Meta
}

func (a Actor) ToJSON(g Game) actorJSON {
	stats := make(map[Stat]int, len(a.Stats))
	unmodified_stats := make(map[Stat]int, len(a.UnmodifiedStats))
	active_modifiers := slices.Collect(maps.Keys(g.AppliedModifiers(a.ID)))

	for stat, v := range a.Stats {
		if _, ok := percentStats[stat]; ok {
			v = v * 100
		}
		stats[stat] = int(v)
	}
	for stat, v := range a.UnmodifiedStats {
		if _, ok := percentStats[stat]; ok {
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
		Faction:            a.Faction,
		Race:               a.Race,
		Level:              a.Level,
		PlayerID:           a.PlayerID,
		PositionID:         NilifyUUID(a.PositionID),
		Actions:            actions,
		Weapon:             weapon,
		Item:               a.Item,
		Effects:            a.GetEffects(),
		Affinities:         slices.Collect(maps.Keys(a.Affinities)),
		AffinityDamage:     affinity_damage,
		AffinityResistance: affinity_resistance,
		AffinityImmunities: a.AffinityImmunities,
		Stats:              stats,
		Stages:             maps.Clone(a.Stages),
		UnmodifiedStats:    unmodified_stats,
		ActiveModifiers:    active_modifiers,
		Wounds:             int(a.Wounds),
		Seen:               a.Meta.Seen,
		State:              a.State,
		Status:             a.Status,
		IsActive:           a.IsActive(),
		IsBulwark:          a.IsBulwark,
		IsAlive:            a.IsAlive,
		IsHidden:           a.IsHidden,
		IsInsulated:        a.IsInsulated,
		IsProtected:        a.IsProtected,
		IsStunned:          a.IsStunned,
	}
}
