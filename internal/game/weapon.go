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
	WeaponTypeRifle    WeaponType = "rifle"
)

type Weapon struct {
	Item
	Actions     []Action
	Clip        int
	OffsetStats map[Stat]float64
	Reload      func(g *Game, parent *Actor, slot uuid.UUID)
	Slot        uuid.UUID
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
	clone := Weapon{
		Item:        w.Item.Clone(),
		Actions:     slices.Clone(w.Actions),
		Clip:        w.Clip,
		OffsetStats: maps.Clone(w.OffsetStats),
		Reload:      w.Reload,
		Slot:        w.Slot,
		Weight:      w.Weight,
		WeaponType:  w.WeaponType,
	}
	clone.BindActions()
	return clone
}

func (w *Weapon) BindActions() {
	for i := range w.Actions {
		w.Actions[i].Weapon = w
	}
}

func (w *Weapon) AddAction(a Action) {
	a.Weapon = w
	w.Actions = append(w.Actions, a)
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
