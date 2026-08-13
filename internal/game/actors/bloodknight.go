package actors

import (
	"grimdark/internal/game"
	"grimdark/internal/game/effects"
	"grimdark/internal/game/weapons"

	"github.com/google/uuid"
)

func Bloodknight() game.Class {
	class := game.NewClass()
	class.ID = uuid.MustParse("019fd8d3-7807-73c5-9311-c69db259a283")
	class.Name = "Bloodknight"
	class.SpriteURL = "/actors/39_64128.png"
	class.Affinities = game.Set[game.Affinity]{
		game.Physical: {},
		game.Blood:    {},
	}
	class.Stats = map[game.Stat]float64{
		game.Health:         99,
		game.Speed:          83,
		game.Melee:          125,
		game.Ranged:         45,
		game.Special:        109,
		game.MartialDefense: 81,
		game.SpecialDefense: 83,
		game.Accuracy:       1,
		game.Evasion:        1,

		game.Actions:        1,
		game.CriticalChance: 1,
		game.CriticalDamage: 1,
		game.DamageReflect:  0,
		game.EffectChance:   1,
	}
	class.Effects = []game.Effect{
		effects.Intimidate,
	}
	class.Options = game.ClassOptions{
		Items: []game.Item{},
		Weapons: []game.Weapon{
			weapons.Greatsword(),
			weapons.HandCrossbow(),
			weapons.RoundShield(),
			weapons.Spear(),
			weapons.TomeOfSacrifice(),
		},
	}

	return class
}
