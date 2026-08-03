package actions

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

var Wildfire = game.Action{
	ID:   uuid.MustParse("019f9c87-eecc-7785-8b0f-22743dda1ad1"),
	Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
	Config: game.ActionConfig{
		Name:         "Wildfire",
		Description:  "Hits all active enemies. This action is not usable from 1st position.",
		Affinity:     game.Fire,
		Stat:         game.Special,
		Accuracy:     game.P(0.9),
		Power:        60,
		CritStage:    0,
		CritModifier: 1.5,
		TargetCount:  0,
	},
	Resolve:          game.MakeAttack(game.AttackConfig{}),
	ValidateContext:  game.ContextTargetLength(0),
	TargetsPredicate: game.NoneActors,
	MapContext:       game.CtxToAllEnemies(),
	DisabledCheck: func(g *game.Game, source game.Actor) bool {
		position, ok := g.GetPosition(source.PositionID)
		if !ok {
			return true
		}

		return position.Rank == 0
	},
}
