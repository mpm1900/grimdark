package actors

import (
	"grimdark/internal/game"
	"grimdark/internal/game/weapons"

	"github.com/google/uuid"
)

func SisterOfSacrifice() game.Class {
	class := game.NewClass()
	class.ID = uuid.MustParse("019fd030-170a-768c-8fcd-4eb7375a0d4b")
	class.Name = "Sister of Sacrifice"
	class.SpriteURL = "/actors/212_crop.png"
	class.Affinities = game.Set[game.Affinity]{
		game.Blood: {},
	}
	class.Stats = map[game.Stat]float64{
		game.Health:         55,
		game.Speed:          100,
		game.Melee:          100,
		game.Ranged:         65,
		game.Special:        100,
		game.MartialDefense: 80,
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
			weapons.FireTome(),
			weapons.PrayerStaff(),
			weapons.RoundShield(),
			weapons.TomeOfSacrifice(),
		},
	}

	return class
}
