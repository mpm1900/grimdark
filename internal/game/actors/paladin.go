package actors

import (
	"grimdark/internal/game"
	"grimdark/internal/game/weapons"

	"github.com/google/uuid"
)

var Paladin = newPaladin()

func newPaladin() game.Class {
	class := game.NewClass()
	class.ID = uuid.MustParse("019f9f6e-520e-74cb-a474-48bdd6ee60cc")
	class.Name = "Paladin"
	class.SpriteURL = "/actors/55_crop.png"
	class.Affinities = map[game.Affinity]struct{}{
		game.Holy:     {},
		game.Physical: {},
	}
	class.Stats = map[game.Stat]float64{
		game.Health:         91,
		game.Speed:          80,
		game.Melee:          134,
		game.Ranged:         80,
		game.Special:        110,
		game.MartialDefense: 95,
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
			weapons.FireTome,
			weapons.Greatsword,
		},
	}

	return class
}
