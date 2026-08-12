package actions

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

func Warp() game.Action {
	id := uuid.MustParse("019fd321-c32b-74df-a765-aca140ca97e7")

	return game.Action{
		ID:   id,
		Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
		Entity: game.MakeEntity(
			id,
			"Warp",
			"Exchanges the positions of target enemies.",
		),
		Config: game.ActionConfig{
			Priority:    game.ActionPriorityDelayed,
			TargetCount: 2,
			Affinity:    game.Arcane,
		},
		ValidateContext:  game.ContextTargetLength(2),
		TargetsPredicate: game.CombineFilters(game.Enemies, game.ActiveActors),
		Resolve: func(g *game.Game, ctx game.Context, this game.ActionContext) []game.Transaction {
			var prev *game.Actor
			for _, target := range g.GetTargets(ctx) {
				if prev != nil {
					this.Push(game.SwapPositions(*prev, target).Bind(game.NewContext()))
				}

				prev = &target
			}
			return this.Done()
		},
	}
}
