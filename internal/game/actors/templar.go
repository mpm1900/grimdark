package actors

import (
	"grimdark/internal/game"
	"grimdark/internal/game/weapons"

	"github.com/google/uuid"
)

func Templar() game.Class {
	class := game.NewClass()
	class.ID = uuid.MustParse("019f5f10-f5c1-7fd6-a1df-98644956735e")
	class.Name = "Templar"
	class.SpriteURL = "/actors/386_64128.png"
	class.Affinities = game.Set[game.Affinity]{
		game.Lightning: {},
		game.Physical:  {},
	}
	class.Stats = map[game.Stat]float64{
		game.Health:         90,
		game.Speed:          65,
		game.Melee:          70,
		game.Ranged:         60,
		game.Special:        130,
		game.MartialDefense: 110,
		game.SpecialDefense: 85,
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
			weapons.PrayerStaff(),
			weapons.RoundShield(),
			weapons.ShortBow(),
		},
	}

	return class
}
