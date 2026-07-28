package actors

import (
	"grimdark/internal/game"
	"grimdark/internal/game/weapons"

	"github.com/google/uuid"
)

var Templar = newTemplar()

func newTemplar() game.Class {
	class := game.NewClass()
	class.ID = uuid.MustParse("019f5f10-f5c1-7fd6-a1df-98644956735e")
	class.Name = "Templar"
	class.SpriteURL = "/actors/386_crop.png"
	class.Affinities = game.Set[game.Affinity]{
		game.Lightning: {},
		game.Physical:  {},
	}
	class.Stats = map[game.Stat]float64{
		game.Health:         75,
		game.Speed:          65,
		game.Melee:          130,
		game.Ranged:         60,
		game.Special:        110,
		game.MartialDefense: 60,
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
			weapons.SlashSword,
			weapons.Greatsword,
		},
	}

	return class
}
