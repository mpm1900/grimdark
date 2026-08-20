package actors

import (
	"grimdark/internal/game"
	"grimdark/internal/game/weapons"

	"github.com/google/uuid"
)

func Knight() game.Class {
	class := game.NewClass()
	class.ID = uuid.MustParse("01a01d8b-a4ea-711e-a671-e65a62925cdf")
	class.Name = "Knight"
	class.SpriteURL = "/actors/352_64128.png"
	class.Affinities = game.Set[game.Affinity]{
		game.Physical: {},
	}
	class.Stats = map[game.Stat]float64{
		game.Health:         100,
		game.Speed:          70,
		game.Melee:          120,
		game.Ranged:         80,
		game.Special:        60,
		game.MartialDefense: 110,
		game.SpecialDefense: 110,
		game.Accuracy:       1,
		game.Evasion:        1,

		game.Actions:        1,
		game.CriticalChance: 1,
		game.CriticalDamage: 1,
		game.DamageReflect:  0,
		game.EffectChance:   1,
	}
	class.Effects = []game.Effect{}
	class.Options = game.ClassOptions{
		Items: []game.Item{},
		Weapons: []game.Weapon{
			weapons.Greatsword(),
			weapons.HandCrossbow(),
			weapons.RoundShield(),
			weapons.ShortBow(),
			weapons.Spear(),
		},
	}

	return class
}
