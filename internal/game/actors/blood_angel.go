package actors

import (
	"grimdark/internal/game"
	"grimdark/internal/game/weapons"

	"github.com/google/uuid"
)

var BloodAngel = newBloodAngel()

func newBloodAngel() game.Class {
	class := game.NewClass()
	class.ID = uuid.MustParse("019f5f10-f5c1-7fd6-a1df-98644956735e")
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
			weapons.BigSword,
		},
	}

	return class
}
