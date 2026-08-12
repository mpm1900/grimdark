package items

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

func Rations() game.Item {
	id := uuid.MustParse("019fcab0-c659-779d-a8e3-a440d272561a")
	entity := game.MakeEntity(
		id,
		"Rations",
		"On turn end, this actor heals 10% HP.",
	)
	effect := game.EffectParent(game.EffectPriorityTriggers, func(g *game.Game, a game.Actor, ctx game.Context) game.Actor {
		return a
	})
	effect.Triggers = append(effect.Triggers, game.Trigger{
		On: game.OnTurnEnd,
		Validate: func(g *game.Game, t_context, m_context game.Context) bool {
			parent, ok := g.GetParent(m_context)
			if !ok {
				return false
			}

			return parent.Stacks[game.Wounds] != 0
		},
		Action: game.Action{
			Entity: entity,
			Resolve: func(g *game.Game, ctx game.Context, this game.ActionContext) []game.Transaction {
				heal_context := game.MakeContextFor(this.Source, this.Source)
				this.Push(game.HealRatioTargets(0.10).Bind(heal_context))
				return this.Done()
			},
		},
	})
	effect.Entity = entity

	return game.Item{
		ID:     id,
		Entity: entity,
		Effects: []game.Effect{
			effect,
		},
	}
}
