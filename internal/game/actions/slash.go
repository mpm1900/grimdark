package actions

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

var Slash = game.Action{
	ID:   uuid.MustParse("019f0aec-8b34-72cc-bbcc-36350e9fa6fb"),
	Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
	Config: game.ActionConfig{
		Name:         "Slash",
		Description:  "",
		Affinity:     game.Physical,
		Stat:         game.Melee,
		Accuracy:     game.P(0.90),
		Power:        90,
		CritStage:    0,
		CritModifier: 1.5,
		TargetCount:  1,
		Range:        game.P(1),
	},
	Resolve:          game.MakeAttack(game.AttackConfig{}),
	ValidateContext:  game.ContextTargetLength(1),
	TargetsPredicate: game.CombineFilters(game.ActiveActors, game.NotSourceActor, game.ActionRange(1)),
}
