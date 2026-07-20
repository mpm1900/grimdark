package effects

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

var Leeched = leeched()

func leeched() game.Effect {
	effect := game.EffectParent(game.EffectPriorityTriggers, func(g *game.Game, a game.Actor, ctx game.Context) game.Actor {
		return a
	})
	effect.Triggers = append(effect.Triggers, game.Trigger{
		On: game.OnTurnEnd,
		Validate: func(g *game.Game, t_context, m_context game.Context) bool {
			return true
		},
		Action: game.Action{
			Config: game.ActionConfig{
				Name: "Leeched",
			},
			Resolve: func(g *game.Game, ctx game.Context, this game.ActionContext) []game.Transaction {
				parent, ok := g.GetParent(ctx)
				if !ok {
					return this.Done()
				}

				modifier, ok := g.GetModifier(ctx.ModifierID)
				if !ok {
					return this.Done()
				}

				health := parent.Stats[game.Health]
				amount := health * 0.12
				dmg_context := game.MakeContextFor(parent, parent)
				this.Push(game.DamageTargets(amount, false).Bind(dmg_context))

				for _, target := range g.GetTargets(modifier.Context) {
					heal_context := game.MakeContextFor(parent, target)
					this.Push(game.DamageTargets(-amount, false).Bind(heal_context))
				}

				return this.Done()
			},
		},
	})

	effect.ID = uuid.MustParse("019f8182-e430-7fd3-92d6-6b46385d03a8")
	effect.Name = "Leeched"
	effect.Description = "On turn end, this actor loses 12% of their max Health to another actor."

	return effect
}
