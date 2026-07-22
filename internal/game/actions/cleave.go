package actions

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

var Cleave = game.Action{
	ID:   uuid.MustParse("019f8b2e-55cb-7cc5-a83a-12ca212535f9"),
	Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
	Config: game.ActionConfig{
		Name:         "Cleave",
		Description:  "Damages all enemy actors in 1st and 2nd position.",
		Affinity:     game.Kinetic,
		Stat:         game.Melee,
		Accuracy:     game.P(0.90),
		Power:        70,
		Lifesteal:    0,
		Hits:         1,
		CritStage:    0,
		CritModifier: 1.5,
		TargetCount:  0,
	},
	Resolve:          game.MakeAttack(game.AttackConfig{}),
	MapContext:       game.CtxToRangeEnemies(2),
	ValidateContext:  game.ContextTargetLength(0),
	TargetsPredicate: game.CombineFilters(game.ActiveActors, game.Enemies),
	DisabledCheck: func(g *game.Game, source game.Actor) bool {
		position, ok := g.GetPosition(source.PositionID)
		if !ok {
			return true
		}

		return position.Rank != 0
	},
}
