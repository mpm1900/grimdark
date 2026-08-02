package actors

import (
	"grimdark/internal/game"
	"grimdark/internal/game/weapons"

	"github.com/google/uuid"
)

var SisterOfFire = newSisterOfFire()

func newSisterOfFire() game.Class {
	class := game.NewClass()
	class.ID = uuid.MustParse("019f5f11-933a-7cd3-bac9-5133bba94c7b")
	class.Name = "Sister of Fire"
	class.SpriteURL = "/actors/230_crop.png"
	class.Affinities = game.Set[game.Affinity]{
		game.Fire: {},
	}
	class.Stats = map[game.Stat]float64{
		game.Health:         65,
		game.Speed:          110,
		game.Melee:          55,
		game.Ranged:         65,
		game.Special:        130,
		game.MartialDefense: 60,
		game.SpecialDefense: 105,
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
			weapons.HandCrossbow,
			weapons.PrayerStaff,
			weapons.RoundShield,
		},
	}

	return class
}
