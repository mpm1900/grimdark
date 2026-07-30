package actions

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

var Poke = game.Action{
	ID:   uuid.MustParse("019fb005-52e8-7224-95ca-b6aae73e4723"),
	Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
	Config: game.ActionConfig{
		Name:         "Poke",
		Description:  "",
		Affinity:     game.Physical,
		Stat:         game.Melee,
		Accuracy:     game.P(0.90),
		Power:        80,
		Lifesteal:    0,
		Hits:         1,
		Cooldown:     0,
		CritStage:    0,
		CritModifier: 1.5,
		TargetCount:  1,
		Range:        game.P(2),
	},
	Resolve:          game.MakeAttack(game.AttackConfig{}),
	ValidateContext:  game.ContextTargetLength(1),
	TargetsPredicate: game.CombineFilters(game.ActiveActors, game.NotSourceActor, game.ActionRange(2)),
}
