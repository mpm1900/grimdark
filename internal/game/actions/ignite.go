package actions

import (
	"grimdark/internal/game"
	"grimdark/internal/game/effects"

	"github.com/google/uuid"
)

var Ignite = game.Action{
	ID:   uuid.MustParse("019fb674-25a3-7361-99b0-c8d999f9caa4"),
	Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
	Config: game.ActionConfig{
		Name:        "Ignite",
		Description: "Burns target.",
		Affinity:    game.Fire,
		Stat:        game.Special,
		Accuracy:    game.P(0.80),
		TargetCount: 1,
	},
	Resolve: func(g *game.Game, ctx game.Context, this game.ActionContext) []game.Transaction {
		resolve := game.AddTargetsEffects(game.StatusConfig{}, game.NewContext(), effects.Burned())
		return resolve(g, ctx, this)
	},
	ValidateContext:  game.ContextTargetLength(1),
	TargetsPredicate: game.CombineFilters(game.OtherActors, game.ActiveActors),
}
