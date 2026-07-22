package game

import (
	"encoding/json"
	"maps"
	"slices"

	"github.com/google/uuid"
)

type ClassOptions struct {
	Items   []Item
	Weapons []Weapon
}

type Class struct {
	ID         uuid.UUID
	Actions    []Action
	Affinities map[Affinity]struct{}
	Arms       int // the amount of weapons that can be equipped
	Effects    []Effect
	Faction    ActorFaction
	Name       string
	Options    ClassOptions
	Race       ActorRace
	SpriteURL  string
	Stats      map[Stat]float64
	Strength   int // weight of weapons that can be equipped
}

type classOptionsJSON struct {
	Items   []Item       `json:"items"`
	Weapons []weaponJSON `json:"weapons"`
}
type classJSON struct {
	ID         uuid.UUID        `json:"ID"`
	Actions    []actionJSON     `json:"actions"`
	Affinities []Affinity       `json:"affinities"`
	Arms       int              `json:"arms"`
	Effects    []Effect         `json:"effects"`
	Faction    ActorFaction     `json:"faction"`
	Name       string           `json:"name"`
	Options    classOptionsJSON `json:"options"`
	Race       ActorRace        `json:"race"`
	SpriteURL  string           `json:"sprite_url"`
	Stats      map[Stat]float64 `json:"stats"`
	Strength   int              `json:"strength"`
}

func NewClass() Class {
	return Class{
		ID:         uuid.New(),
		Actions:    []Action{},
		Affinities: map[Affinity]struct{}{},
		Arms:       2,
		Effects:    []Effect{},
		Faction:    FactionImperium,
		Name:       "",
		Race:       RaceHuman,
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
			Actions:        1,
			CriticalChance: 1,
			CriticalDamage: 1,
			EffectChance:   1,
		},
		Strength: 2,
	}
}

func (o ClassOptions) Clone() ClassOptions {
	return ClassOptions{
		Items:   slices.Clone(o.Items),
		Weapons: slices.Clone(o.Weapons),
	}
}

func (c Class) Clone() Class {
	return Class{
		ID:         c.ID,
		Actions:    slices.Clone(c.Actions),
		Affinities: maps.Clone(c.Affinities),
		Arms:       c.Arms,
		Effects:    slices.Clone(c.Effects),
		Faction:    c.Faction,
		Options:    c.Options.Clone(),
		Name:       c.Name,
		Race:       c.Race,
		SpriteURL:  c.SpriteURL,
		Stats:      maps.Clone(c.Stats),
		Strength:   c.Strength,
	}
}

// omit options to save on data
func (c Class) CloneForActor() Class {
	clone := c.Clone()
	clone.Options = ClassOptions{}
	return clone
}

func (c Class) ToJSON() classJSON {
	actions := make([]actionJSON, len(c.Actions))
	for i, a := range c.Actions {
		actions[i] = a.ToJSONStatic()
	}
	weapons := make([]weaponJSON, len(c.Options.Weapons))
	for i, w := range c.Options.Weapons {
		weapons[i] = w.ToJSONStatic()
	}
	return classJSON{
		ID:         c.ID,
		Actions:    actions,
		Affinities: slices.Collect(maps.Keys(c.Affinities)),
		Arms:       c.Arms,
		Effects:    c.Effects,
		Faction:    c.Faction,
		Name:       c.Name,
		Options: classOptionsJSON{
			Weapons: weapons,
			Items:   c.Options.Items,
		},
		Race:      c.Race,
		SpriteURL: c.SpriteURL,
		Stats:     c.Stats,
		Strength:  c.Strength,
	}
}

func (c Class) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.ToJSON())
}
