package game

import (
	"maps"
	"slices"

	"github.com/google/uuid"
)

type Weapon struct {
	ID       uuid.UUID
	Actions  []Action
	AuxStats map[Stat]float64
	Effects  []Effect
	Name     string
}

type weaponJSON struct {
	ID       uuid.UUID        `json:"ID"`
	Actions  []actionJSON     `json:"actions"`
	AuxStats map[Stat]float64 `json:"aux_stats"`
	Effects  []Effect         `json:"effects"`
	Name     string           `json:"name"`
}

func (w Weapon) Clone() Weapon {
	return Weapon{
		ID:       w.ID,
		Actions:  slices.Clone(w.Actions),
		AuxStats: maps.Clone(w.AuxStats),
		Effects:  slices.Clone(w.Effects),
		Name:     w.Name,
	}
}

func (w Weapon) ToJSON(g Game, source Actor) weaponJSON {
	actions := make([]actionJSON, len(w.Actions))
	for i, action := range w.Actions {
		actions[i] = action.ToJSON(g, source)
	}

	return weaponJSON{
		ID:       w.ID,
		Actions:  actions,
		AuxStats: w.AuxStats,
		Effects:  w.Effects,
		Name:     w.Name,
	}
}
