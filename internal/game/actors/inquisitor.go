package actors

import (
	"grimdark/internal/game"
	"grimdark/internal/game/weapons"

	"github.com/google/uuid"
)

var Inquisitor = newInquisitor()

func newInquisitor() game.Class {
	class := game.NewClass()
	class.ID = uuid.MustParse("019ff406-1e84-7618-a62e-91df5634c5f6")
	class.Name = "Inquisitor"
	class.SpriteURL = "/actors/389_64128.png"
	class.Affinities = game.Set[game.Affinity]{
		game.Holy:  {},
		game.Blood: {},
	}
	class.Stats = map[game.Stat]float64{
		game.Health:         70,
		game.Speed:          125,
		game.Melee:          70,
		game.Ranged:         100,
		game.Special:        135,
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
			weapons.Dagger(),
			weapons.HandCrossbow(),
			weapons.RoundShield(),
			weapons.Spear(),
			weapons.TomeOfSacrifice(),
			weapons.FireTome(),
			weapons.PrayerStaff(),
		},
	}

	return class
}
