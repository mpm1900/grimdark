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

type ActionState struct {
	Cooldown      int
	CooldownBonus int
	IsDisabled    bool
	PriorityBonus int
	RangeBonus    int
}

type ActorMeta struct {
	ActiveTurns      int
	InactiveTurns    int
	LastUsedActionID uuid.UUID
	Seen             bool
}

type Actor struct {
	ID         uuid.UUID
	Class      Class
	Name       string
	Level      int
	PlayerID   uuid.UUID
	PositionID uuid.UUID

	Item    *Item
	WeaponL *Weapon
	WeaponR *Weapon

	ActionStates       map[uuid.UUID]ActionState
	AffinityDamage     map[Affinity]int
	AffinityResistance map[Affinity]int
	AffinityImmunities map[Affinity]float64
	EffectImmunities   map[uuid.UUID]struct{}
	effectStates       map[uuid.UUID]EffectState
	OffsetStats        map[Stat]float64
	Stages             map[Stat]int
	Stats              map[Stat]float64
	UnmodifiedStats    map[Stat]float64
	Stacks             map[Stack]float64

	State  ActorState
	Status ActorStatus

	IsAlive     bool
	IsBulwark   bool // stops collateral penetration
	IsHidden    bool // cannot be targeted by single-target actions (unimplemented)
	IsInsulated bool // is immune from the secondary effects of attacking attacks (ie through AddResultEffects())
	IsProtected bool // protected from actions that check accuracy
	IsStunned   bool // cannot act and cannot queue commands (may not be needed)

	Meta ActorMeta
}

func NewActor(class Class, config ActorConfig) *Actor {
	var weapon_l *Weapon
	var weapon_r *Weapon
	var item *Item

	for _, w := range class.Options.Weapons {
		if config.WeaponL != nil && w.ID == *config.WeaponL {
			weapon_l = P(w.Clone())
		}
		if config.WeaponR != nil && w.ID == *config.WeaponR {
			weapon_r = P(w.Clone())
		}
	}
	for _, i := range class.Options.Items {
		if slices.Contains(config.Items, i.ID) {
			item = P(i.Clone())
		}
	}
	return &Actor{
		ID:         uuid.New(),
		Class:      class.CloneForActor(),
		Name:       class.Name,
		PlayerID:   uuid.Nil,
		PositionID: uuid.Nil,
		Level:      100,

		WeaponL: weapon_l,
		WeaponR: weapon_r,
		Item:    item,

		AffinityDamage:     map[Affinity]int{},
		AffinityImmunities: map[Affinity]float64{},
		AffinityResistance: map[Affinity]int{},
		EffectImmunities:   map[uuid.UUID]struct{}{},
		OffsetStats:        map[Stat]float64{},
		Stacks: map[Stack]float64{
			Wounds: 0,
		},
		Stages: map[Stat]int{},
		Stats:  maps.Clone(class.Stats),

		State:  StateGrounded,
		Status: StatusNone,

		IsAlive:     true,
		IsBulwark:   false,
		IsHidden:    false,
		IsInsulated: false,
		IsProtected: false,
		IsStunned:   false,

		ActionStates: map[uuid.UUID]ActionState{},
		effectStates: map[uuid.UUID]EffectState{},

		Meta: ActorMeta{
			ActiveTurns:      0,
			InactiveTurns:    0,
			LastUsedActionID: uuid.Nil,
			Seen:             false,
		},
	}
}

func (a *Actor) Clone() *Actor {
	var weapon_l *Weapon = nil
	var weapon_r *Weapon = nil
	var item *Item = nil
	if a.WeaponL != nil {
		clone := a.WeaponL.Clone()
		weapon_l = &clone
	}
	if a.WeaponR != nil {
		clone := a.WeaponR.Clone()
		weapon_r = &clone
	}
	if a.Item != nil {
		clone := a.Item.Clone()
		item = &clone
	}
	return &Actor{
		ID:         a.ID,
		Class:      a.Class.CloneForActor(),
		Name:       a.Name,
		Level:      a.Level,
		PlayerID:   a.PlayerID,
		PositionID: a.PositionID,

		WeaponL: weapon_l,
		WeaponR: weapon_r,
		Item:    item,

		AffinityDamage:     maps.Clone(a.AffinityDamage),
		AffinityResistance: maps.Clone(a.AffinityResistance),
		AffinityImmunities: maps.Clone(a.AffinityImmunities),
		EffectImmunities:   maps.Clone(a.EffectImmunities),
		OffsetStats:        maps.Clone(a.OffsetStats),
		Stacks:             maps.Clone(a.Stacks),
		Stages:             maps.Clone(a.Stages),
		Stats:              maps.Clone(a.Stats),
		UnmodifiedStats:    maps.Clone(a.UnmodifiedStats),

		State:  a.State,
		Status: a.Status,

		IsAlive:     a.IsAlive,
		IsBulwark:   a.IsBulwark,
		IsHidden:    a.IsHidden,
		IsInsulated: a.IsInsulated,
		IsProtected: a.IsProtected,
		IsStunned:   a.IsStunned,

		ActionStates: maps.Clone(a.ActionStates),
		effectStates: maps.Clone(a.effectStates),

		Meta: a.Meta,
	}
}

// mappers
func (a *Actor) mapBaseStat(stat Stat, stats map[Stat]float64, offset float64) float64 {
	base := stats[stat]*2 + offset
	ratio := float64(a.Level) / 100
	result := (base*ratio + 5)
	if stat == Health {
		result += float64(a.Level)
	}
	return result
}
func (a *Actor) getStatOffset(stat Stat) float64 {
	offset, ok := a.OffsetStats[stat]
	if !ok {
		offset = 0
	}
	if a.WeaponL != nil {
		weapon_offset, ok := a.WeaponL.OffsetStats[stat]
		if ok {
			offset += weapon_offset
		}
	}
	if a.WeaponR != nil {
		weapon_offset, ok := a.WeaponR.OffsetStats[stat]
		if ok {
			offset += weapon_offset
		}
	}

	return offset
}
func (a *Actor) mapBaseStats() {
	a.UnmodifiedStats = maps.Clone(a.Stats)

	for stat, _ := range a.Stats {
		if _, ok := mappedStats[stat]; !ok {
			continue
		}

		a.Stats[stat] = a.mapBaseStat(stat, a.Stats, a.getStatOffset(stat))
		a.UnmodifiedStats[stat] = a.mapBaseStat(stat, a.UnmodifiedStats, 0)
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

// mutators
func (a *Actor) ApplyDamage(damage float64, resolved Actor) {
	a.Stacks[Wounds] = a.Stacks[Wounds] + damage
	if a.Stacks[Wounds] < 0 {
		a.Stacks[Wounds] = 0
	}

	a.IsAlive = resolved.Stats[Health] > a.Stacks[Wounds]
}
func (a *Actor) UpdateActionState(action_id uuid.UUID, updater func(ActionState) ActionState) {
	state, ok := a.ActionStates[action_id]
	if !ok {
		a.ActionStates[action_id] = updater(ActionState{})
	} else {
		a.ActionStates[action_id] = updater(state)
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

	for aid, state := range a.ActionStates {
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
		a.Meta.LastUsedActionID = uuid.Nil
	} else {
		a.Meta.Seen = true
	}
	if a.PositionID == uuid.Nil {
		a.Meta.ActiveTurns = 0
	}
	a.PositionID = position_id
}

// getters
func (a *Actor) IsActive() bool {
	return a.PositionID != uuid.Nil
}
func (a *Actor) CanAct() bool {
	return !a.IsStunned
}
func (a *Actor) GetAffinityDamage(affinity Affinity) int {
	base := maps.Clone(a.AffinityDamage)
	for affinity := range a.Class.Affinities {
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
func (a *Actor) GetAffinityResistance(affinity Affinity) int {
	value, ok := a.AffinityResistance[affinity]
	if ok {
		return value
	}

	return 0
}
func (a *Actor) GetEffectiveAffinityResistance(affinity Affinity) int {
	return a.GetAffinityResistance(affinity) - affinity.GetBaseModifier(*a)
}
func (a *Actor) GetRemainingHealth() float64 {
	health := a.Stats[Health]
	return health - a.Stacks[Wounds]
}
func (a *Actor) getWeaponEffects() []Effect {
	effects := map[uuid.UUID]Effect{}
	if a.WeaponL != nil {
		for _, e := range a.WeaponL.Effects {
			effects[e.ID] = e
		}
	}
	if a.WeaponR != nil {
		for _, e := range a.WeaponR.Effects {
			effects[e.ID] = e
		}
	}

	return slices.Collect(maps.Values(effects))
}
func (a *Actor) GetEffects() []Effect {
	effects := slices.Clone(a.Class.Effects)
	effects = append(effects, a.getWeaponEffects()...)
	if a.Item != nil {
		effects = append(effects, a.Item.Effects...)
	}

	for i, e := range effects {
		state, ok := a.effectStates[e.ID]
		if ok {
			effects[i].ApplyState(state)
		}
	}

	effects, _ = a.FilterEffectImmunities(effects)
	return effects
}
func (a *Actor) HasEffectImmunity(tag uuid.UUID) bool {
	_, ok := a.EffectImmunities[tag]
	return ok
}
func (a *Actor) FilterEffectImmunities(effects []Effect) ([]Effect, []Effect) {
	result := []Effect{}
	removed := []Effect{}
	for _, effect := range effects {
		if a.HasEffectImmunity(effect.ID) {
			removed = append(removed, effect)
			continue
		}

		result = append(result, effect)
	}

	return result, removed
}
func (a *Actor) GetModifiers() []Modifier {
	modifiers := []Modifier{}
	for _, effect := range a.GetEffects() {
		if effect.Ready() {
			context := MakeContextFrom(*a)
			context.EffectID = effect.ID
			modifier := effect.Bind(context)
			modifiers = append(modifiers, modifier)
		}
	}

	return modifiers
}
func (a *Actor) GetActions() []Action {
	actions := []Action{}
	seen := map[uuid.UUID]struct{}{}

	addActions := func(next []Action) {
		for _, action := range next {
			if _, ok := seen[action.ID]; ok {
				continue
			}

			seen[action.ID] = struct{}{}
			actions = append(actions, action)
		}
	}

	addActions(a.Class.Actions)
	if a.WeaponL != nil {
		addActions(a.WeaponL.Actions)
	}
	if a.WeaponR != nil {
		addActions(a.WeaponR.Actions)
	}

	addActions(GLOBAL_ACTIONS)
	actions = slices.DeleteFunc(actions, func(action Action) bool {
		return action.ActiveCheck != nil && !action.ActiveCheck(*a)
	})
	return actions
}
func (a *Actor) GetActionByID(action_id uuid.UUID) (Action, bool) {
	for _, action := range a.GetActions() {
		if action.ID == action_id {
			return action, true
		}
	}

	return Action{}, false
}
func (a *Actor) Targetable() bool {
	if a.IsActive() {
		return !a.IsHidden
	}

	return true
}
