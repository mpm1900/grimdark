package items

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

var Rations = rations()

func rations() game.Item {
	effect := game.EffectParent(game.EffectPriorityTriggers, func(g *game.Game, a game.Actor, ctx game.Context) game.Actor {
		return a
	})
	effect.Triggers = append(effect.Triggers, game.Trigger{
		On: game.OnTurnEnd,
		Action: game.Action{
			Config: game.ActionConfig{
				Name: "Rations",
			},
			Resolve: func(g *game.Game, ctx game.Context, this game.ActionContext) []game.Transaction {
				heal_context := game.MakeContextFor(this.Source, this.Source)
				this.Push(game.HealRatioTargets(0.10).Bind(heal_context))
				return this.Done()
			},
		},
	})

	effect.Name = "Rations"
	effect.Description = "On turn end, this actor heals 10% HP."

	return game.Item{
		ID:          uuid.MustParse("019fcab0-c659-779d-a8e3-a440d272561a"),
		Name:        "Rations",
		Description: "On turn end, this actor heals 10% HP.",
		Effects: []game.Effect{
			effect,
		},
	}
}
