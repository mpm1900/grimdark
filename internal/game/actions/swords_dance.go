package actions

import (
	"grimdark/internal/game"
	"grimdark/internal/game/effects"

	"github.com/google/uuid"
)

func SwordsDance() game.Action {
	id := uuid.MustParse("019f0aee-7aae-7efc-b8e7-d514f3ad2b18")

	return game.Action{
		ID:   id,
		Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
		Entity: game.MakeEntity(
			id,
			"Swords Dance",
			"Raises user's Melee and Speed.",
		),
		Config: game.ActionConfig{
			Affinity:    game.Physical,
			TargetCount: 0,
		},
		Resolve: game.AddSourceEffects(
			game.StatusConfig{},
			1,
			effects.StatUpSource(game.Speed, 1),
			effects.StatUpSource(game.Melee, 1),
		),
		ValidateContext:  game.TrueGameFilter,
		TargetsPredicate: game.NoneActors,
		/*
			DisabledCheck: func(g *game.Game, source game.Actor) bool {
				return source.Meta.ActiveTurns > 1
			},
		*/
		ActiveCheck: game.IsDualWielding,
	}
}
