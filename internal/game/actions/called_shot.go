package actions

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

func CalledShot() game.Action {
	id := uuid.MustParse("019f85ac-e872-7fd0-9511-a7df176f402f")

	return game.Action{
		ID:   id,
		Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
		Entity: game.MakeEntity(
			id,
			"Called Shot",
			"This action cannot miss but is -1 priority. This action is only usable from the 3rd position.",
		),
		Config: game.ActionConfig{
			Affinity:     game.Physical,
			Stat:         game.Ranged,
			Power:        95,
			CritStage:    0,
			CritModifier: 2,
			TargetCount:  1,
			Priority:     game.ActionPriorityDelayed,
		},
		Resolve:          game.MakeAttack(game.AttackConfig{}),
		ValidateContext:  game.ContextTargetLength(1),
		TargetsPredicate: game.CombineFilters(game.ActiveActors, game.OtherActors),
		DisabledCheck: func(g *game.Game, source game.Actor) bool {
			position, ok := g.GetPosition(source.PositionID)
			if !ok {
				return true
			}

			return position.Rank != 2
		},
	}
}
