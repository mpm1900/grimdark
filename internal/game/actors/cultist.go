package actors

import (
	"grimdark/internal/game"
	"grimdark/internal/game/weapons"

	"github.com/google/uuid"
)

var Cultist = newCultist()

func newCultist() game.Class {
	class := game.NewClass()
	class.ID = uuid.MustParse("019f5f12-29e1-7cc9-bfeb-468df5c53990")
	class.Name = "Cultist"
	class.SpriteURL = "/actors/373_crop.png"
	class.Affinities = map[game.Affinity]struct{}{
		game.Arcane: {},
		game.Blood:  {},
	}
	class.Stats = map[game.Stat]float64{
		game.Health:         76,
		game.Speed:          108,
		game.Melee:          71,
		game.Ranged:         108,
		game.Special:        108,
		game.MartialDefense: 71,
		game.SpecialDefense: 71,
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
			weapons.SlashSword,
			weapons.Greatsword,
			weapons.LongBow(),
			weapons.Dagger,
		},
	}

	return class
}
