package actions

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

func HealingPrayer() game.Action {
	id := uuid.MustParse("019fae66-768a-736e-b01e-3ed0699d3f19")

	return game.Action{
		ID:   id,
		Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
		Entity: game.MakeEntity(
			id,
			"Healing Prayer",
			"Heals the user for up-to 50% health.",
		),
		Config: game.ActionConfig{
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
	}
}
