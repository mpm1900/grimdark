package actions

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

var HealingPrayer = game.Action{
	ID:   uuid.MustParse("019fae66-768a-736e-b01e-3ed0699d3f19"),
	Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
	Config: game.ActionConfig{
		Name:        "Healing Prayer",
		Description: "Heals the user for up-to 50% health.",
		Affinity:    game.Holy,
		TargetCount: 0,
		Uses:        game.P(2),
	},
	Resolve: func(g *game.Game, ctx game.Context, this game.ActionContext) []game.Transaction {
		mut := game.HealRatioTargets(0.50)
		heal_ctx := game.MakeContextFor(this.Source, this.Source)
		this.Push(mut.Bind(heal_ctx))
		return this.Done()
	},
	ValidateContext:  game.TrueGameFilter,
	TargetsPredicate: game.NoneActors,
	ActiveCheck:      game.IsDualWielding,
}
