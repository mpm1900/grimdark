package effects

import (
	"grimdark/internal/game"
)

var Intimidate = intimidate()

func intimidate() game.Effect {
	effect := game.EffectSource(game.EffectPriorityTriggers, func(g *game.Game, a game.Actor, ctx game.Context) game.Actor {
		return a
	})
	effect.Triggers = append(effect.Triggers, game.Trigger{
		On:       game.OnActorEnter,
		Validate: game.TriggerSourceMatchesModifierParent,
		Action: game.Action{
			Config: game.ActionConfig{
				Name: "Intimidate",
			},
			Resolve: func(g *game.Game, ctx game.Context, this game.ActionContext) []game.Transaction {
				other_actors := g.FindActors(game.CombineFilters(game.ActiveActors, game.NotSourceActor), ctx)
				for _, target := range other_actors {
					target_ctx := game.MakeModifierContext(this.Source, target)
					mutation := game.AddModifiers(StatDownTargets(game.Special, 1).Bind(target_ctx))
					this.Push(mutation.Bind(game.NewContext()))
				}

				return this.Done()
			},
		},
	})

	effect.Name = "Intimidate"
	effect.Description = "When this actor enters, the attack stat of other active actors is lowered by one stage."
	return effect
}
