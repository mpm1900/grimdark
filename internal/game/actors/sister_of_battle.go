package actors

import (
	"grimdark/internal/game"
	"grimdark/internal/game/items"
	"grimdark/internal/game/weapons"

	"github.com/google/uuid"
)

var SisterOfBattle = newSisterOfBattle()

func newSisterOfBattle() game.Class {
	class := game.NewClass()
	class.ID = uuid.MustParse("019f5f11-933a-7cd3-bac9-5133bba94c7b")
	class.Name = "Adepta Sororitas"
	class.SpriteURL = "/img/sis.png"
	class.Affinities = map[game.Affinity]struct{}{
		game.Fire:      {},
		game.Lightning: {},
	}
	class.Stats = map[game.Stat]float64{
		game.Health:         80,
		game.Speed:          60,
		game.Melee:          74,
		game.Ranged:         95,
		game.Special:        95,
		game.MartialDefense: 80,
		game.SpecialDefense: 116,
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
		Items: []game.Item{
			items.TestItem,
		},
		Weapons: []game.Weapon{
			weapons.FireTome,
			weapons.Greatsword,
			weapons.Pistol,
		},
	}

	return class
}
