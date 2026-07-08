package actors

import "grimdark/internal/game"

func NecronWarrior() game.Class {
	class := game.NewClass()
	class.Name = "Necron Warrior"
	class.Affinities = map[game.Affinity]struct{}{
		game.Arcane:    {},
		game.Lightning: {},
	}
	class.Stats = map[game.Stat]float64{
		game.Health:         100,
		game.Speed:          75,
		game.Melee:          75,
		game.Ranged:         75,
		game.Special:        75,
		game.MartialDefense: 90,
		game.SpecialDefense: 90,
		game.Accuracy:       1,
		game.Evasion:        1,
		game.CriticalChance: 1,
		game.CriticalDamage: 1,
		game.EffectChance:   1,
	}
	class.Effects = []game.Effect{}
	class.SpriteURL = "/img/nec.png"

	return class
}
