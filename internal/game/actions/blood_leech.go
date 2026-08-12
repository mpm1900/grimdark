package actions

import (
	"grimdark/internal/game"
	"grimdark/internal/game/effects"

	"github.com/google/uuid"
)

func BloodLeech() game.Action {
	id := uuid.MustParse("019f8185-3375-74dc-8812-bfa2ea14259d")

	return game.Action{
		ID:   id,
		Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
		Entity: game.MakeEntity(
			id,
			"Blood Leech",
			"Applies Leeched to target.",
		),
		Config: game.ActionConfig{
			Affinity:    game.Poison,
			TargetCount: 1,
		},
		Resolve: func(g *game.Game, ctx game.Context, this game.ActionContext) []game.Transaction {
			modifier_context := game.NewContext()
			modifier_context.AddModifierTarget(this.Source)
			resolve := game.AddTargetsEffects(game.StatusConfig{}, modifier_context, effects.Leeched)
			return resolve(g, ctx, this)
		},
		ValidateContext:  game.ContextTargetLength(1),
		TargetsPredicate: game.OtherActors,
	}
}
