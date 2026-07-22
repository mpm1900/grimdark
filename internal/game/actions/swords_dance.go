package actions

import (
	"grimdark/internal/game"
	"grimdark/internal/game/effects"

	"github.com/google/uuid"
)

var SwordsDance = game.Action{
	ID:   uuid.MustParse("019f0aee-7aae-7efc-b8e7-d514f3ad2b18"),
	Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
	Config: game.ActionConfig{
		Name:        "Swords Dance",
		Description: "Raises user's Melee and Speed.",
		Affinity:    game.Kinetic,
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
	DisabledCheck: func(g *game.Game, source game.Actor) bool {
		return source.Meta.ActiveTurns > 1
	},
	ActiveCheck: func(source game.Actor) bool {
		found := uuid.Nil
		for _, w := range source.Weapons {
			if found == uuid.Nil {
				found = w.Item.ID
				continue
			}

			if w.Item.ID != found {
				return false
			}
		}
		return true
	},
}
