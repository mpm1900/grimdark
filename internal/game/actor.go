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

	ActionsState       map[uuid.UUID]ActionState
	AffinityDamage     map[Affinity]int
	AffinityResistance map[Affinity]int
	AffinityImmunities map[Affinity]float64
	OffsetStats        map[Stat]float64
	Stages             map[Stat]int
	Stats              map[Stat]float64
	UnmodifiedStats    map[Stat]float64

	State  ActorState
	Status ActorStatus
	Wounds float64

	IsAlive     bool
	IsBulwark   bool // stops collateral penetration
	IsHidden    bool // cannot be targeted by single-target actions
	IsInsulated bool // is immune from the secondary effects of attacking attacks (ie through AddResultEffects())
	IsProtected bool // protected from actions that check accuracy
	IsStunned   bool // staggered + and cannot queue commands (may not be needed)

	Meta ActorMeta
}

type actorJSON struct {
	ID                 uuid.UUID            `json:"ID"`
	Actions            []actionJSON         `json:"actions"`
	ActiveModifiers    []uuid.UUID          `json:"active_modifiers"`
	Affinities         []Affinity           `json:"affinities"`
	AffinityDamage     map[Affinity]int     `json:"affinity_damage"`
	AffinityImmunities map[Affinity]float64 `json:"affinity_immunities"`
	AffinityResistance map[Affinity]int     `json:"affinity_resistance"`
	Effects            []Effect             `json:"effects"`
	Faction            ActorFaction         `json:"faction"`
	IsActive           bool                 `json:"is_active"`
	IsAlive            bool                 `json:"is_alive"`
	IsBulwark          bool                 `json:"is_bulwark"`
	IsHidden           bool                 `json:"is_hidden"`
	IsInsulated        bool                 `json:"is_insulated"`
	IsPlayer           bool                 `json:"is_player"`
	IsProtected        bool                 `json:"is_protected"`
	IsStunned          bool                 `json:"is_stunned"`
	Item               *Item                `json:"item"`
	Level              int                  `json:"level"`
	Name               string               `json:"name"`
	PlayerID           uuid.UUID            `json:"player_ID"`
	PositionID         *uuid.UUID           `json:"position_ID"`
	Race               ActorRace            `json:"race"`
	Seen               bool                 `json:"-"`
	SpriteURL          string               `json:"sprite_url"`
	State              ActorState           `json:"state"`
	Stats              map[Stat]int         `json:"stats"`
	Stages             map[Stat]int         `json:"stages"`
	Status             ActorStatus          `json:"status"`
	UnmodifiedStats    map[Stat]int         `json:"unmodified_stats"`
	WeaponL            *weaponJSON          `json:"weapon_l"`
	WeaponR            *weaponJSON          `json:"weapon_r"`
	Wounds             int                  `json:"wounds"`
}

func NewActor(class Class, config ActorConfig) Actor {
	var weapon_l *Weapon
	var weapon_r *Weapon
	var item *Item

	for _, w := range class.Options.Weapons {
		if w.ID == config.WeaponL {
			weapon_l = P(w.Clone())
		}
		if w.ID == config.WeaponR {
			weapon_r = P(w.Clone())
		}
	}
	for _, i := range class.Options.Items {
		if slices.Contains(config.Items, i.ID) {
			item = P(i.Clone())
		}
	}
	return Actor{
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
		AffinityResistance: map[Affinity]int{},
		AffinityImmunities: map[Affinity]float64{},
		Stages:             map[Stat]int{},
		OffsetStats:        map[Stat]float64{},
		Stats:              maps.Clone(class.Stats),

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
	var weapon_l *Weapon = nil
	var weapon_r *Weapon = nil
	if a.WeaponL != nil {
		clone := a.WeaponL.Clone()
		weapon_l = &clone
	}
	if a.WeaponR != nil {
		clone := a.WeaponR.Clone()
		weapon_r = &clone
	}
	return Actor{
		ID:         a.ID,
		Class:      a.Class.CloneForActor(),
		Name:       a.Name,
		Level:      a.Level,
		PlayerID:   a.PlayerID,
		PositionID: a.PositionID,

		WeaponL: weapon_l,
		WeaponR: weapon_r,
		Item:    a.Item,

		AffinityDamage:     maps.Clone(a.AffinityDamage),
		AffinityResistance: maps.Clone(a.AffinityResistance),
		AffinityImmunities: maps.Clone(a.AffinityImmunities),
		UnmodifiedStats:    maps.Clone(a.UnmodifiedStats),
		Stages:             maps.Clone(a.Stages),
		OffsetStats:        maps.Clone(a.OffsetStats),
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

// mappers
func (a *Actor) mapBaseStat(stat Stat, stats map[Stat]float64, aux float64) float64 {
	base := stats[stat]*2 + aux
	ratio := float64(a.Level) / 100
	result := (base*ratio + 5)
	if stat == Health {
		result += float64(a.Level)
	}
	return result
}
func (a *Actor) getStatOffset(stat Stat) float64 {
	aux, ok := a.OffsetStats[stat]
	if !ok {
		aux = 0
	}
	if a.WeaponL != nil {
		weapon_aux, ok := a.WeaponL.AuxStats[stat]
		if ok {
			aux += weapon_aux
		}
	}
	if a.WeaponR != nil {
		weapon_aux, ok := a.WeaponR.AuxStats[stat]
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
func (a Actor) IsActive() bool {
	return a.PositionID != uuid.Nil
}
func (a Actor) CanAct() bool {
	return !a.IsStunned
}
func (a Actor) GetAffinityDamage(affinity Affinity) int {
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
	effects := slices.Clone(a.Class.Effects)
	if a.WeaponL != nil {
		effects = append(effects, a.WeaponL.Effects...)
	}
	if a.WeaponR != nil {
		effects = append(effects, a.WeaponR.Effects...)
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

	addActions(GLOBAL_ACTIONS)
	addActions(a.Class.Actions)
	if a.WeaponL != nil {
		addActions(a.WeaponL.Actions)
	}
	if a.WeaponR != nil {
		addActions(a.WeaponR.Actions)
	}

	return slices.DeleteFunc(actions, func(action Action) bool {
		return action.ActiveCheck != nil && !action.ActiveCheck(a)
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

// json
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
	for affinity := range a.Class.Affinities {
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

	var weapon_l *weaponJSON = nil
	var weapon_r *weaponJSON = nil
	if a.WeaponL != nil {
		clone := a.WeaponL.ToJSON(g, a)
		weapon_l = &clone
	}
	if a.WeaponR != nil {
		clone := a.WeaponR.ToJSON(g, a)
		weapon_r = &clone
	}

	return actorJSON{
		ID:                 a.ID,
		Name:               a.Name,
		Faction:            a.Class.Faction,
		Race:               a.Class.Race,
		Level:              a.Level,
		PlayerID:           a.PlayerID,
		PositionID:         NilifyUUID(a.PositionID),
		Actions:            actions,
		WeaponL:            weapon_l,
		WeaponR:            weapon_r,
		Item:               a.Item,
		Effects:            a.GetEffects(),
		Affinities:         slices.Collect(maps.Keys(a.Class.Affinities)),
		AffinityDamage:     affinity_damage,
		AffinityResistance: affinity_resistance,
		AffinityImmunities: a.AffinityImmunities,
		Stats:              stats,
		Stages:             maps.Clone(a.Stages),
		UnmodifiedStats:    unmodified_stats,
		ActiveModifiers:    active_modifiers,
		Wounds:             int(a.Wounds),
		Seen:               a.Meta.Seen,
		SpriteURL:          a.Class.SpriteURL,
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
