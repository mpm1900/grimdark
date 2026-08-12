package actions

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

func QuickStrike() game.Action {
	id := uuid.MustParse("019fe214-39d0-7128-8a10-e0b117ff5dd2")

	return game.Action{
		ID:   id,
		Tags: []game.ActionTag{game.ATActor, game.ATWeapon, game.ATConditional},
		Entity: game.MakeEntity(
			id,
			"Quick Strike",
			"This action has +1 priority. This action is only usable from 1st position.",
		),
		Config: game.ActionConfig{
			Affinity:     game.Physical,
			Stat:         game.Melee,
			Power:        55,
			Accuracy:     game.P(1.0),
			CritStage:    0,
			CritModifier: 1.5,
			TargetCount:  1,
			Range:        game.P(1),
			Priority:     game.ActionPriorityQuick,
		},
		Resolve:          game.MakeAttack(game.AttackConfig{}),
		ValidateContext:  game.ContextTargetLength(1),
		TargetsPredicate: game.CombineFilters(game.ActiveActors, game.OtherActors, game.ActionRange(1)),
		ActiveCheck:      game.IsDualWielding,
		DisabledCheck: func(g *game.Game, source game.Actor) bool {
			position, ok := g.GetPosition(source.PositionID)
			if !ok {
				return true
			}

			return position.Rank != 0
		},
	}
}
