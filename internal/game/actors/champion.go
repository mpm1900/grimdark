package actors

import (
	"grimdark/internal/game"
	"grimdark/internal/game/effects"
	"grimdark/internal/game/weapons"

	"github.com/google/uuid"
)

var Champion = newChampion()

func newChampion() game.Class {
	_ = effects.ProtectedWhere(func(g *game.Game, a game.Actor, ctx game.Context) bool {
		active_context := g.State().ActiveContext
		if active_context == nil {
			return false
		}

		parent, ok := g.GetParent(ctx)
		if !ok {
			return false
		}

		if active_context.ActionID == uuid.Nil {
			return false
		}

		if !active_context.HasTarget(parent) {
			return false
		}

		source, ok := g.GetSource(*active_context)
		if !ok {
			return false
		}

		action, ok := source.GetActionByID(active_context.ActionID)
		if !ok {
			return false
		}

		return action.Config.Affinity == game.Psychic
	})

	weakness_immune := game.EffectSource(game.EffectPriorityImmunities, func(g *game.Game, a game.Actor, ctx game.Context) game.Actor {
		a.EffectImmunities.Push(effects.Weakened().ID)
		return a
	})

	class := game.NewClass()
	class.ID = uuid.MustParse("019f5f12-6e78-7eda-b638-980453e3eaba")
	class.Name = "Champion"
	class.SpriteURL = "/actors/402_crop.png"
	class.Affinities = game.Set[game.Affinity]{
		game.Physical: {},
	}
	class.Stats = map[game.Stat]float64{
		game.Health:         110,
		game.Speed:          90,
		game.Melee:          115,
		game.Ranged:         100,
		game.Special:        50,
		game.MartialDefense: 90,
		game.SpecialDefense: 80,
		game.Accuracy:       1,
		game.Evasion:        1,

		game.Actions:        1,
		game.CriticalChance: 1,
		game.CriticalDamage: 1,
		game.DamageReflect:  0,
		game.EffectChance:   1,
	}
	class.Effects = []game.Effect{weakness_immune}
	class.Options = game.ClassOptions{
		Items: []game.Item{},
		Weapons: []game.Weapon{
			weapons.Greatsword,
			weapons.HandCrossbow,
			weapons.RoundShield,
			weapons.ShortBow,
			weapons.Spear,
		},
	}

	return class
}
