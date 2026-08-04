package actions

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

var Latch = game.Action{
	ID:   uuid.MustParse("019fcb35-92d0-7292-b5fe-2d20a9ec1818"),
	Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
	Config: game.ActionConfig{
		Name:        "Latch",
		Description: "Pulls target to 1st position. This action is only usable from first position and has -1 priority.",
		Affinity:    game.Physical,
		Stat:        game.Ranged,
		Accuracy:    game.P(0.7),
		TargetCount: 1,
		Priority:    game.ActionPriorityDelayed,
	},
	Resolve: func(g *game.Game, ctx game.Context, this game.ActionContext) []game.Transaction {
		source := g.GetSourceAction(ctx)
		for _, target := range g.GetTargets(ctx) {
			pull_ctx := game.MakeContextFor(source, target)
			this.Push(game.PushTargetsToFront().Bind(pull_ctx))
		}

		return this.Done()
	},
	ValidateContext:  game.ContextTargetLength(1),
	TargetsPredicate: game.CombineFilters(game.Enemies, game.ActiveActors, game.NotPositionRank(0)),
}
