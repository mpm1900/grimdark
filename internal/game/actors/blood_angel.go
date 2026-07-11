package actors

import (
	"grimdark/internal/game"
	"grimdark/internal/game/weapons"
)

var BloodAngel = newBloodAngel()

func newBloodAngel() game.Class {
	class := game.NewClass()
	class.Name = "Blood Angel"
	class.SpriteURL = "/img/spm2.png"
	class.Affinities = map[game.Affinity]struct{}{
		game.Arcane:  {},
		game.Kinetic: {},
	}
	class.Stats = map[game.Stat]float64{
		game.Health:         100,
		game.Speed:          61,
		game.Melee:          130,
		game.Ranged:         95,
		game.Special:        90,
		game.MartialDefense: 80,
		game.SpecialDefense: 80,
		game.Accuracy:       1,
		game.Evasion:        1,

		game.CriticalChance: 1,
		game.CriticalDamage: 1,
		game.DamageReflect:  0,
		game.EffectChance:   1,
	}
	class.Effects = []game.Effect{}
	class.Options = game.ClassOptions{
		Items: []game.Item{},
		Weapons: []game.Weapon{
			weapons.SlashSword,
			weapons.BigSword,
		},
	}

	return class
}
