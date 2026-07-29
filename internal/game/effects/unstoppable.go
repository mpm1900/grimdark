package effects

import "grimdark/internal/game"

func Unstoppable() game.Effect {
	effect := game.EffectSource(game.EffectPriorityTriggers, func(g *game.Game, a game.Actor, ctx game.Context) game.Actor {
		return a
	})
	effect.Triggers = append(effect.Triggers, game.Trigger{
		On:       game.OnDamageRecieve,
		Validate: game.TriggerTargetMatchesModifierParent,
		Action: game.Action{
			Config: game.ActionConfig{
				Name: "Unstoppable",
			},
			Resolve: func(g *game.Game, ctx game.Context, this game.ActionContext) []game.Transaction {
				mutation := game.AddModifiers(StatUpTargets(game.MartialDefense, 1).Bind(ctx))
				this.Push(mutation.Bind(game.NewContext()))

				return this.Done()
			},
		},
	})
	effect.Name = "Unstoppable"
	effect.Description = "When this actor takes damage, they gain +1 to their Defense stat stage."

	return effect
}
