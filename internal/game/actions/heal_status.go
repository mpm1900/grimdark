package actions

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

func HealStatus() game.Action {
	id := uuid.MustParse("01a01ae7-749b-7160-8305-671c900bf86e")

	return game.Action{
		ID:   id,
		Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
		Entity: game.MakeEntity(
			id,
			"Heal Status",
			"Removes the user's status condition.",
		),
		Config: game.ActionConfig{
			Affinity:    game.Holy,
			TargetCount: 0,
		},
		ValidateContext:  game.TrueGameFilter,
		TargetsPredicate: game.NoneActors,
		Resolve: func(g *game.Game, ctx game.Context, this game.ActionContext) []game.Transaction {
			mut := game.RemoveStatusTargets()
			mut_ctx := game.MakeContextFor(this.Source, this.Source)
			this.Push(mut.Bind(mut_ctx))
			return this.Done()
		},
	}
}
