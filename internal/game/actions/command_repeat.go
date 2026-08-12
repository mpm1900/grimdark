package actions

import (
	"grimdark/internal/game"
	"grimdark/internal/game/effects"

	"github.com/google/uuid"
)

func repeat() game.Effect {
	effect := effects.ChoiceLocked()
	effect.Duration = game.P(5)

	return effect
}

func CommandRepeat() game.Action {
	id := uuid.MustParse("019fb3c4-dccb-756a-9d75-3feb46ccba19")

	return game.Action{
		ID:   id,
		Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
		Entity: game.MakeEntity(
			id,
			"Command: Repeat",
			"Target must repeat their last used action for 5 turns.",
		),
		Config: game.ActionConfig{
			Affinity:    game.Psychic,
			TargetCount: 1,
		},
		Resolve: func(g *game.Game, ctx game.Context, this game.ActionContext) []game.Transaction {
			resolve := game.AddTargetsEffects(game.StatusConfig{}, ctx, repeat())
			return resolve(g, ctx, this)
		},
		ValidateContext:  game.ContextTargetLength(1),
		TargetsPredicate: game.OtherActors,
	}
}
