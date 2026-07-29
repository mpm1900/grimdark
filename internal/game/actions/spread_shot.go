package actions

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

var SpreadShot = game.Action{
	ID:   uuid.MustParse("019f8b2e-55cb-7cc5-a83a-12ca212535f9"),
	Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
	Config: game.ActionConfig{
		Name:         "Spread Shot",
		Description:  "Damages all enemy actors. This action is only usable in 2nd position.",
		Affinity:     game.Physical,
		Stat:         game.Ranged,
		Accuracy:     game.P(0.90),
		Power:        40,
		Lifesteal:    0,
		Hits:         1,
		CritStage:    0,
		CritModifier: 1.5,
		TargetCount:  3,
	},
	Resolve:          game.MakeAttack(game.AttackConfig{}),
	MapContext:       game.CtxToAllEnemies(),
	ValidateContext:  game.ContextTargetLength(0),
	TargetsPredicate: game.NoneActors,
	DisabledCheck: func(g *game.Game, source game.Actor) bool {
		position, ok := g.GetPosition(source.PositionID)
		if !ok {
			return true
		}

		return position.Rank != 1
	},
}
