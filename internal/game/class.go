package game

import (
	"maps"
	"slices"

	"github.com/google/uuid"
)

type Class struct {
	ID         uuid.UUID
	Actions    []Action
	Affinities map[Affinity]struct{}
	Effects    []Effect
	Faction    ActorFaction
	Name       string
	Race       ActorRace
	SpriteURL  string
	Stats      map[Stat]float64
}

func NewClass() Class {
	return Class{
		ID:         uuid.New(),
		Name:       "",
		Race:       RaceHuman,
		Faction:    FactionImperium,
		Actions:    []Action{},
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
		},
		Effects: []Effect{},
	}
}

func (c Class) Clone() Class {
	return Class{
		ID:         c.ID,
		Actions:    slices.Clone(c.Actions),
		Affinities: maps.Clone(c.Affinities),
		Effects:    slices.Clone(c.Effects),
		Faction:    c.Faction,
		Name:       c.Name,
		Race:       c.Race,
		SpriteURL:  c.SpriteURL,
		Stats:      maps.Clone(c.Stats),
	}
}
