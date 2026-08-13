package actors

import (
	"grimdark/internal/game"
	"grimdark/internal/game/effects"
	"grimdark/internal/game/weapons"

	"github.com/google/uuid"
)

func Vicar() game.Class {
	class := game.NewClass()
	class.ID = uuid.MustParse("019fb617-3034-7348-8747-74279a8cd80d")
	class.Name = "Vicar"
	class.SpriteURL = "/actors/344_64128.png"
	class.Affinities = game.Set[game.Affinity]{
		game.Holy: {},
	}
	class.Stats = map[game.Stat]float64{
		game.Health:         60,
		game.Speed:          70,
		game.Melee:          60,
		game.Ranged:         70,
		game.Special:        144,
		game.MartialDefense: 64,
		game.SpecialDefense: 114,
		game.Accuracy:       1,
		game.Evasion:        1,

		game.Actions:        1,
		game.CriticalChance: 1,
		game.CriticalDamage: 1,
		game.DamageReflect:  0,
		game.EffectChance:   1,
	}
	class.Effects = []game.Effect{
		effects.Blithe(),
	}
	class.Options = game.ClassOptions{
		Items: []game.Item{},
		Weapons: []game.Weapon{
			weapons.FireTome(),
			weapons.HandCrossbow(),
			weapons.PrayerStaff(),
			weapons.RoundShield(),
		},
	}

	return class
}
