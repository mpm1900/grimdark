package actors

import (
	"grimdark/internal/game"
	"grimdark/internal/game/weapons"

	"github.com/google/uuid"
)

var NecronWarrior = newNecronWarrior()

func newNecronWarrior() game.Class {
	class := game.NewClass()
	class.ID = uuid.MustParse("019f5f11-626c-7b69-a0f9-8115db2e2c8c")
	class.Name = "Necron Warrior"
	class.SpriteURL = "/img/nec.png"
	class.Affinities = map[game.Affinity]struct{}{
		game.Arcane:    {},
		game.Lightning: {},
	}
	class.Stats = map[game.Stat]float64{
		game.Health:         100,
		game.Speed:          75,
		game.Melee:          75,
		game.Ranged:         75,
		game.Special:        75,
		game.MartialDefense: 90,
		game.SpecialDefense: 90,
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
			weapons.SlashSword,
			weapons.Greatsword,
		},
	}

	return class
}
