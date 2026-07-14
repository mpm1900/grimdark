package game

import (
	"maps"
	"slices"

	"github.com/google/uuid"
)

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
	Stacks             map[Stack]int        `json:"stacks"`
	State              ActorState           `json:"state"`
	Stats              map[Stat]int         `json:"stats"`
	Stages             map[Stat]int         `json:"stages"`
	Status             ActorStatus          `json:"status"`
	UnmodifiedStats    map[Stat]int         `json:"unmodified_stats"`
	WeaponL            *weaponJSON          `json:"weapon_l"`
	WeaponR            *weaponJSON          `json:"weapon_r"`
	Wounds             int                  `json:"wounds"`
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
		Stacks:             a.Stacks,
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
