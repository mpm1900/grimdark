package actors

import (
	"grimdark/internal/game"
)

func NewSisterOfBattle() game.Class {
	class := game.NewClass()
	class.Name = "Adepta Sororitas"
	class.Affinities = map[game.Affinity]struct{}{
		game.Fire:    {},
		game.Kinetic: {},
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
		game.CriticalChance: 1,
		game.CriticalDamage: 1,
		game.EffectChance:   1,
	}
	class.Effects = []game.Effect{}
	class.SpriteURL = "/img/sis.png"

	return class
}
