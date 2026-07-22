package actions

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

var Charge = game.Action{
	ID:   uuid.MustParse("019f8aa6-b01a-79e0-b1ac-19b12f3e6e95"),
	Tags: []game.ActionTag{game.ATActor, game.ATWeapon, game.ATMovement},
	Config: game.ActionConfig{
		Name:         "Charge",
		Description:  "Charges forward to 1st position to attack. This action is only usable from 3rd or 2nd position.",
		Affinity:     game.Kinetic,
		Stat:         game.Melee,
		Power:        90,
		Accuracy:     game.P(1.0),
		Lifesteal:    0,
		Hits:         1,
		CritStage:    0,
		CritModifier: 1.5,
		TargetCount:  1,
		Priority:     game.ActionPriorityDefault,
	},
	Resolve: game.MakeAttack(game.AttackConfig{
		BeforeAttack: func(g *game.Game, context game.Context, this *game.ActionContext) {
			this.Push(game.PushSourceToFront().Bind(context))
		},
	}),
	ValidateContext:  game.ContextTargetLength(1),
	TargetsPredicate: game.CombineFilters(game.Enemies, game.PositionRank(0)),
	DisabledCheck: func(g *game.Game, source game.Actor) bool {
		position, ok := g.GetPosition(source.PositionID)
		if !ok {
			return true
		}

		return position.Rank == 0
	},
}
