package actions

import (
	"grimdark/internal/game"
	"grimdark/internal/game/effects"

	"github.com/google/uuid"
)

var SwordsDance = game.Action{
	ID: uuid.MustParse("019f0aee-7aae-7efc-b8e7-d514f3ad2b18"),
	Config: game.ActionConfig{
		Name:     "Swords Dance",
		Affinity: game.Kinetic,
	},
	Resolve: game.AddSourceEffects(
		1,
		effects.StatUpSource(game.Speed, 1),
		effects.StatUpSource(game.Melee, 1),
	),
	ValidateContext:  game.TrueGameFilter,
	TargetsPredicate: game.NoneActors,
}
