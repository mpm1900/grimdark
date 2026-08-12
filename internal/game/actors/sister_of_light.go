package actors

import (
	"grimdark/internal/game"
	"grimdark/internal/game/weapons"

	"github.com/google/uuid"
)

var SisterOfLight = newSisterOfLight()

func newSisterOfLight() game.Class {
	class := game.NewClass()
	class.ID = uuid.MustParse("019fc507-ad97-7179-8049-d0d85eaa8cc5")
	class.Name = "Sister of Light"
	class.SpriteURL = "/actors/236_64128.png"
	class.Affinities = game.Set[game.Affinity]{
		game.Holy: {},
	}
	class.Stats = map[game.Stat]float64{
		game.Health:         255,
		game.Speed:          55,
		game.Melee:          10,
		game.Ranged:         45,
		game.Special:        75,
		game.MartialDefense: 10,
		game.SpecialDefense: 135,
		game.Accuracy:       1,
		game.Evasion:        1,

		game.Actions:        1,
		game.CriticalChance: 1,
		game.CriticalDamage: 1,
		game.DamageReflect:  0,
		game.EffectChance:   1,
	}
	class.Items = 2
	class.Effects = []game.Effect{}
	class.Options = game.ClassOptions{
		Items: []game.Item{},
		Weapons: []game.Weapon{
			weapons.FireTome(),
			weapons.PrayerStaff(),
			weapons.RoundShield(),
		},
	}

	return class
}
