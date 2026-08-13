package actors

import (
	"grimdark/internal/game"
	"grimdark/internal/game/effects"
	"grimdark/internal/game/weapons"

	"github.com/google/uuid"
)

func Seeker() game.Class {
	class := game.NewClass()
	class.ID = uuid.MustParse("019fc0cc-6c4f-74ec-99d1-94d976fc96c0")
	class.Name = "Seeker"
	class.SpriteURL = "/actors/418_64128.png"
	class.Affinities = game.Set[game.Affinity]{
		game.Arcane: {},
	}
	class.Stats = map[game.Stat]float64{
		game.Health:         120,
		game.Speed:          60,
		game.Melee:          50,
		game.Ranged:         50,
		game.Special:        110,
		game.MartialDefense: 70,
		game.SpecialDefense: 70,
		game.Accuracy:       1,
		game.Evasion:        1,

		game.Actions:        1,
		game.CriticalChance: 1,
		game.CriticalDamage: 1,
		game.DamageReflect:  0,
		game.EffectChance:   1,
	}
	class.Effects = []game.Effect{
		effects.PriorityFailure(),
	}
	class.Options = game.ClassOptions{
		Items: []game.Item{},
		Weapons: []game.Weapon{
			weapons.FireTome(),
			weapons.TomeOfSacrifice(),
		},
	}

	return class
}
