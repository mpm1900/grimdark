package actions

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

func Cleave() game.Action {
	id := uuid.MustParse("019fe214-985a-775e-9d8f-21eaf072af7d")

	return game.Action{
		ID:   id,
		Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
		Entity: game.MakeEntity(
			id,
			"Cleave",
			"Damages all enemy actors in 1st and 2nd position. This action is only usable in 1st position.",
		),
		Config: game.ActionConfig{
			Affinity:     game.Physical,
			Stat:         game.Melee,
			Accuracy:     game.P(0.90),
			Power:        70,
			CritStage:    0,
			CritModifier: 1.5,
			TargetCount:  2,
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
}
