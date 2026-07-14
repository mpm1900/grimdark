package game

import (
	"maps"
	"slices"

	"github.com/google/uuid"
)

type WeaponType string

const (
	WeaponTypeSword    WeaponType = "sword"
	WeaponTypeBigSword WeaponType = "big-sword"
	WeaponTypePistol   WeaponType = "pistol"
)

type Weapon struct {
	Item
	Actions     []Action
	OffsetStats map[Stat]float64
	Weight      int
	WeaponType  WeaponType
}

type weaponJSON struct {
	ID          uuid.UUID        `json:"ID"`
	Actions     []actionJSON     `json:"actions"`
	Description string           `json:"description"`
	Effects     []Effect         `json:"effects"`
	Weight      int              `json:"weight"`
	Name        string           `json:"name"`
	OffsetStats map[Stat]float64 `json:"offset_stats"`
	WeaponType  WeaponType       `json:"weapon_type"`
}

func (w Weapon) Clone() Weapon {
	return Weapon{
		Item:        w.Item.Clone(),
		Actions:     slices.Clone(w.Actions),
		OffsetStats: maps.Clone(w.OffsetStats),
		Weight:      w.Weight,
		WeaponType:  w.WeaponType,
	}
}

func (w Weapon) ToJSON(g *Game, source Actor) weaponJSON {
	actions := make([]actionJSON, len(w.Actions))
	for i, action := range w.Actions {
		actions[i] = action.ToJSON(g, source)
	}

	return weaponJSON{
		ID:          w.ID,
		Actions:     actions,
		OffsetStats: w.OffsetStats,
		Description: w.Description,
		Effects:     w.Effects,
		Weight:      w.Weight,
		Name:        w.Name,
		WeaponType:  w.WeaponType,
	}
}

func (w Weapon) ToJSONStatic() weaponJSON {
	actions := make([]actionJSON, len(w.Actions))
	for i, action := range w.Actions {
		actions[i] = action.ToJSONStatic()
	}

	return weaponJSON{
		ID:          w.ID,
		Actions:     actions,
		OffsetStats: w.OffsetStats,
		Description: w.Description,
		Effects:     w.Effects,
		Weight:      w.Weight,
		Name:        w.Name,
		WeaponType:  w.WeaponType,
	}
}
