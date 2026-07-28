package actors

import (
	"grimdark/internal/game"
	"grimdark/internal/game/weapons"

	"github.com/google/uuid"
)

var Rogue = newRogue()

func newRogue() game.Class {
	class := game.NewClass()
	class.ID = uuid.MustParse("019f9fc5-5d62-7c24-9f2e-56f651b15020")
	class.Name = "Rogue"
	class.SpriteURL = "/actors/348_crop.png"
	class.Affinities = game.Set[game.Affinity]{
		game.Poison:   {},
		game.Physical: {},
	}
	class.Stats = map[game.Stat]float64{
		game.Health:         80,
		game.Speed:          120,
		game.Melee:          100,
		game.Ranged:         120,
		game.Special:        40,
		game.MartialDefense: 60,
		game.SpecialDefense: 80,
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
			weapons.Dagger,
		},
	}

	return class
}
