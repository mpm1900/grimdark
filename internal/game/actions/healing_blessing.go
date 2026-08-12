package actions

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

func HealingBlessing() game.Action {
	id := uuid.MustParse("019fb154-1ee5-74bd-8212-58e30a719886")

	return game.Action{
		ID:   id,
		Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
		Entity: game.MakeEntity(
			id,
			"Healing Blessing",
			"Heals all active allies for up-to 25% health.",
		),
		Config: game.ActionConfig{
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
}
