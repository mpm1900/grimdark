package actions

import (
	"grimdark/internal/game"
	"grimdark/internal/game/effects"

	"github.com/google/uuid"
)

func Protect() game.Action {
	id := uuid.MustParse("019f87f2-fd37-72b4-ab5d-bf23e5830011")

	return game.Action{
		ID:   id,
		Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
		Entity: game.MakeEntity(
			id,
			"Protect",
			"User is protected from attacks and actions.",
		),
		Config: game.ActionConfig{
			Affinity:    game.Physical,
			TargetCount: 0,
			Priority:    game.ActionPriorityProtect,
			Cooldown:    1,
		},
		Resolve: game.AddSourceEffects(
			game.StatusConfig{},
			1,
			effects.ProtectedSource(),
		),
		ValidateContext:  game.TrueGameFilter,
		TargetsPredicate: game.NoneActors,
	}
}
