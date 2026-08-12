package actions

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

func BackStrike() game.Action {
	id := uuid.MustParse("019fa451-5fd6-7254-9433-d09977283382")

	return game.Action{
		ID:   id,
		Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
		Entity: game.MakeEntity(
			id,
			"Back Strike",
			"The user moves backwards after attacking. This action is only usable from 1st position.",
		),
		Config: game.ActionConfig{
			Affinity:     game.Physical,
			Stat:         game.Melee,
			Power:        70,
			Accuracy:     game.P(1.0),
			CritStage:    0,
			CritModifier: 1.5,
			TargetCount:  1,
			Range:        game.P(1),
			Priority:     game.ActionPriorityQuick,
		},
		Resolve: game.MakeAttack(game.AttackConfig{
			OnFinally: func(g *game.Game, context game.Context, this *game.ActionContext) {
				this.Push(game.PushSourceBackwards().Bind(context))
			},
		}),
		ValidateContext:  game.ContextTargetLength(1),
		TargetsPredicate: game.CombineFilters(game.ActiveActors, game.OtherActors, game.ActionRange(1)),
		DisabledCheck: func(g *game.Game, source game.Actor) bool {
			position, ok := g.GetPosition(source.PositionID)
			if !ok {
				return true
			}

			return position.Rank != 0
		},
	}
}
