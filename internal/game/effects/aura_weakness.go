package effects

import (
	"grimdark/internal/game"
)

var AuraOfWeakness = auraOfWeakness()

func auraOfWeakness() game.Effect {
	effect := game.EffectSource(game.EffectPriorityTriggers, func(g game.Game, a game.Actor, ctx game.Context) game.Actor {
		return a
	})
	effect.Triggers = append(effect.Triggers, game.Trigger{
		On:       game.OnActorEnter,
		Validate: game.TriggerSourceMatchesModifierParent,
		Action: game.Action{
			Config: game.ActionConfig{
				Name: "Aura of Weakness",
			},
			Resolve: func(g *game.Game, ctx game.Context, this game.ActionContext) []game.Transaction {
				effect := game.EffectActorsActiveOther(game.EffectPriorityPostStagesStats, func(g game.Game, a game.Actor, ctx game.Context) game.Actor {
					a.Stats[game.Melee] = a.Stats[game.Melee] * 0.75
					return a
				})
				effect.Name = "Weakened"
				effect.CheckSuccess = func(g *game.Game, e game.Effect, ctx game.Context) {
					targets := g.FindActors(game.CombineFilters(game.ActiveActors, game.NotSourceActor), ctx)
					for _, target := range targets {
						log_ctx := game.MakeContextFor(this.Source, target)
						log_ctx.EffectID = effect.ID
						g.PushLogMeta(game.NewLog(
							"$target$ gained $effect$.",
							game.EffectTermsTarget(target, effect),
						).Bind(log_ctx))
					}
				}
				mutation := game.AddModifiers(effect.Bind(ctx))
				this.Push(mutation.Bind(game.NewContext()))
				return this.Done()
			},
		},
	})

	effect.Name = "Aura of Weakness"
	return effect
}
