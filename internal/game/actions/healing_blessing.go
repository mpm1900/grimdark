package actions

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

var HealingBlessing = game.Action{
	ID:   uuid.MustParse("019fb154-1ee5-74bd-8212-58e30a719886"),
	Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
	Config: game.ActionConfig{
		Name:        "Healing Blessing",
		Description: "Heals all active allies for up-to 25% health.",
		Affinity:    game.Holy,
		TargetCount: 0,
		Uses:        game.P(2),
	},
	Resolve: func(g *game.Game, ctx game.Context, this game.ActionContext) []game.Transaction {
		mut := game.HealRatioTargets(0.25)
		this.Push(mut.Bind(ctx))
		return this.Done()
	},
	MapContext:       game.CtxToAllies(),
	ValidateContext:  game.TrueGameFilter,
	TargetsPredicate: game.NoneActors,
}
