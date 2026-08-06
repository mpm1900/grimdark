package actions

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

var Warp = game.Action{
	ID:   uuid.MustParse("019fd321-c32b-74df-a765-aca140ca97e7"),
	Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
	Config: game.ActionConfig{
		Name:        "Warp",
		Description: "Exchanges the positions of target enemies.",
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
