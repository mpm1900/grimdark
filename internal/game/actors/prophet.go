package actors

import (
	"grimdark/internal/game"
	"grimdark/internal/game/weapons"

	"github.com/google/uuid"
)

var Prophet = newProphet()

func newProphet() game.Class {
	class := game.NewClass()
	class.ID = uuid.MustParse("019f5f11-626c-7b69-a0f9-8115db2e2c8c")
	class.Name = "Prophet"
	class.SpriteURL = "/actors/81_crop.png"
	class.Affinities = game.Set[game.Affinity]{
		game.Blood: {},
		game.Holy:  {},
	}
	class.Stats = map[game.Stat]float64{
		game.Health:         75,
		game.Speed:          104,
		game.Melee:          68,
		game.Ranged:         75,
		game.Special:        114,
		game.MartialDefense: 72,
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
			weapons.FireTome,
			weapons.SlashSword,
			weapons.Greatsword,
		},
	}

	return class
}
