package effects

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

var AuraOfWeakness = auraOfWeakness()

func auraOfWeakness() game.Effect {
	effect := game.EffectSource(game.EffectPriorityTriggers, func(g *game.Game, a game.Actor, ctx game.Context) game.Actor {
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
				mutation := game.AddModifiers(Weakened().Bind(ctx))
				this.Push(mutation.Bind(ctx))
				return this.Done()
			},
		},
	})

	effect.Name = "Aura of Weakness"
	return effect
}

func meleeDown(g *game.Game, a game.Actor, ctx game.Context) game.Actor {
	a.Stats[game.Melee] = a.Stats[game.Melee] * 0.75
	return a
}
func Weakened() game.Effect {
	effect := game.EffectActorsActiveOther(game.EffectPriorityPostStagesStats, meleeDown)
	effect.ID = uuid.MustParse("019f58b9-4edd-7c60-bccd-08ff88120a5b")
	effect.Name = "Weakened"
	effect.CheckSuccess = game.EffectGainWhereOnSuccess(
		game.OtherActors,
	)
	return effect
}
