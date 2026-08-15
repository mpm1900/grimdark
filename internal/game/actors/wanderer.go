package actors

import (
	"grimdark/internal/game"
	"grimdark/internal/game/weapons"

	"github.com/google/uuid"
)

func Wanderer() game.Class {
	class := game.NewClass()
	class.ID = uuid.MustParse("01a00690-8b87-7633-a3c0-4337811c4eea")
	class.Name = "Wanderer"
	class.SpriteURL = "/actors/394_64128.png"
	class.Affinities = game.Set[game.Affinity]{
		game.Physical: {},
	}
	class.Stats = map[game.Stat]float64{
		game.Health:         80,
		game.Speed:          100,
		game.Melee:          100,
		game.Ranged:         100,
		game.Special:        80,
		game.MartialDefense: 100,
		game.SpecialDefense: 100,
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
			weapons.Dagger(),
			weapons.FireTome(),
			weapons.Greatsword(),
			weapons.HandCrossbow(),
			weapons.RoundShield(),
			weapons.ShortBow(),
		},
	}

	return class
}
