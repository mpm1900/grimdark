package actors

import (
	"grimdark/internal/game"
	"grimdark/internal/game/weapons"

	"github.com/google/uuid"
)

func Warrior() game.Class {
	class := game.NewClass()
	class.ID = uuid.MustParse("01a02548-3fa0-70c8-8006-60225fb8addc")
	class.Name = "Warrior"
	class.SpriteURL = "/actors/134_64128.png"
	class.Affinities = game.Set[game.Affinity]{
		game.Physical: {},
	}
	class.Stats = map[game.Stat]float64{
		game.Health:         85,
		game.Speed:          115,
		game.Melee:          130,
		game.Ranged:         70,
		game.Special:        50,
		game.MartialDefense: 80,
		game.SpecialDefense: 70,
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
