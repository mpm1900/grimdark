package actions

import (
	"grimdark/internal/game"
	"grimdark/internal/game/effects"

	"github.com/google/uuid"
)

var BlessingOfMichael = game.Action{
	ID:   uuid.MustParse("019fae7e-3b36-73cb-b387-d3d798c67d29"),
	Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
	Config: game.ActionConfig{
		Name:        "Blessing of Michael",
		Description: "Raises the target's Melee and Martial Defensse.",
		Affinity:    game.Holy,
		TargetCount: 1,
	},
	Resolve: func(g *game.Game, ctx game.Context, this game.ActionContext) []game.Transaction {
		resolve := game.AddTargetsEffects(
			game.StatusConfig{},
			ctx,
			effects.StatUpTargets(game.MartialDefense, 1),
			effects.StatUpTargets(game.Melee, 1),
		)
		return resolve(g, ctx, this)
	},
	ValidateContext:  game.ContextTargetLength(1),
	TargetsPredicate: game.CombineFilters(game.Allies, game.ActiveActors),
}
