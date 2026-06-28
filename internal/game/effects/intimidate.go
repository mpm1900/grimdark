package effects

import (
	"grimdark/internal/game"
)

func Intimidate() game.Effect {
	effect := game.EffectSource(game.EffectPriorityTriggers, func(g game.Game, a game.Actor, ctx game.Context) game.Actor {
		return a
	})
	effect.Triggers = append(effect.Triggers, game.Trigger{
		On:       game.OnActorEnter,
		Validate: game.TriggerSourceMatchesModifierParent,
		Action: game.Action{
			Resolve: func(g *game.Game, ctx game.Context, this game.ActionContext) []game.Transaction {
				modifiers := []game.Modifier{}
				other_actors := g.FindActors(game.CombineFilters(game.ActiveActors, game.NotSourceActor), ctx)
				for _, target := range other_actors {
					target_ctx := game.MakeContextFor(this.Source, target)
					modifiers = append(modifiers, StatDownTargets(game.Special, 1).Bind(target_ctx))
				}

				mutation := game.AddModifiers(modifiers...)
				this.Push(mutation.Bind(game.NewContext()))

				return this.Done()
			},
		},
	})
	effect.Name = "Intimidate"

	return effect
}
