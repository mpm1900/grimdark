package actions

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

func Headshot() game.Action {
	id := uuid.MustParse("019f87f3-2b7b-7726-aaac-129f070adff8")

	return game.Action{
		ID:   id,
		Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
		Entity: game.MakeEntity(
			id,
			"Headshot",
			"This action always results in a critical hit. This action is only usable from 2nd or 3rd position.",
		),
		Config: game.ActionConfig{
			Affinity:     game.Physical,
			Stat:         game.Ranged,
			Power:        55,
			Accuracy:     game.P(0.90),
			CritStage:    3,
			CritModifier: 1.5,
			TargetCount:  1,
			Priority:     0,
		},
		Resolve:          game.MakeAttack(game.AttackConfig{}),
		ValidateContext:  game.ContextTargetLength(1),
		TargetsPredicate: game.CombineFilters(game.ActiveActors, game.OtherActors),
		DisabledCheck: func(g *game.Game, source game.Actor) bool {
			position, ok := g.GetPosition(source.PositionID)
			if !ok {
				return true
			}

			return position.Rank == 0
		},
	}
}
