package actors

import (
	"grimdark/internal/game"
	"grimdark/internal/game/effects"
	"grimdark/internal/game/weapons"

	"github.com/google/uuid"
)

var Ultramarine = newUltramarine()

func newUltramarine() game.Class {
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
		a.EffectImmunities[effects.Weakened().ID] = struct{}{}
		return a
	})

	class := game.NewClass()
	class.ID = uuid.MustParse("019f5f12-6e78-7eda-b638-980453e3eaba")
	class.Name = "Storm Warden"
	class.SpriteURL = "/img/spm.png"
	class.Affinities = map[game.Affinity]struct{}{
		game.Kinetic: {},
	}
	class.Stats = map[game.Stat]float64{
		game.Health:         100,
		game.Speed:          61,
		game.Melee:          134,
		game.Ranged:         95,
		game.Special:        60,
		game.MartialDefense: 100,
		game.SpecialDefense: 110,
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
			weapons.SlashSword,
			weapons.Greatsword,
			weapons.Pistol,
		},
	}

	return class
}
