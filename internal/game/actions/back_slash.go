package actions

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

var BackStrike = game.Action{
	ID:   uuid.MustParse("019fa451-5fd6-7254-9433-d09977283382"),
	Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
	Config: game.ActionConfig{
		Name:         "Back Strike",
		Description:  "The user moves backwards after attacking. This action is only usable from 1st position.",
		Affinity:     game.Physical,
		Stat:         game.Melee,
		Power:        70,
		Accuracy:     game.P(1.0),
		Lifesteal:    0,
		Hits:         1,
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
