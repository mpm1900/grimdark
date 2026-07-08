package game

import (
	"maps"
	"slices"

	"github.com/google/uuid"
)

type WeaponType string

type Weapon struct {
	Item
	Actions    []Action
	AuxStats   map[Stat]float64
	Hands      int
	WeaponType WeaponType
}

type weaponJSON struct {
	ID          uuid.UUID        `json:"ID"`
	Actions     []actionJSON     `json:"actions"`
	AuxStats    map[Stat]float64 `json:"aux_stats"`
	Description string           `json:"description"`
	Effects     []Effect         `json:"effects"`
	Hands       int              `json:"hands"`
	Name        string           `json:"name"`
	WeaponType  WeaponType       `json:"weapon_type"`
}

func (w Weapon) Clone() Weapon {
	return Weapon{
		Item:       w.Item.Clone(),
		Actions:    slices.Clone(w.Actions),
		AuxStats:   maps.Clone(w.AuxStats),
		Hands:      w.Hands,
		WeaponType: w.WeaponType,
	}
}

func (w Weapon) ToJSON(g Game, source Actor) weaponJSON {
	actions := make([]actionJSON, len(w.Actions))
	for i, action := range w.Actions {
		actions[i] = action.ToJSON(g, source)
	}

	return weaponJSON{
		ID:          w.ID,
		Actions:     actions,
		AuxStats:    w.AuxStats,
		Description: w.Description,
		Effects:     w.Effects,
		Hands:       w.Hands,
		Name:        w.Name,
		WeaponType:  w.WeaponType,
	}
}
