package actions

import (
	"grimdark/internal/game"
	"grimdark/internal/game/effects"

	"github.com/google/uuid"
)

var StunningShot = game.Action{
	ID:   uuid.MustParse("019fb10e-e44c-711f-98bb-d6d3aebc0fca"),
	Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
	Config: game.ActionConfig{
		Name:         "Stunning Shot",
		Description:  "Staggers the target. Only usable the first turn out in battle.",
		Affinity:     game.Physical,
		Stat:         game.Ranged,
		Accuracy:     game.P(1.0),
		Power:        20,
		Lifesteal:    0,
		Recoil:       0,
		Hits:         1,
		CritStage:    0,
		CritModifier: 1.5,
		TargetCount:  1,
		Priority:     game.ActionPrioritySpecial,
	},
	Resolve: game.MakeAttack(game.AttackConfig{
		OnSuccessResult: func(g *game.Game, context game.Context, this *game.ActionContext, result game.DamageResult) {
			game.AddResultEffects(
				1,
				effects.StaggerTargets,
			)(g, context, this, result)
		},
	}),
	ValidateContext:  game.ContextTargetLength(1),
	TargetsPredicate: game.CombineFilters(game.ActiveActors, game.OtherActors, game.ActionRange(1)),
	DisabledCheck: func(g *game.Game, source game.Actor) bool {
		if source.Meta.ActiveTurns > 1 {
			return true
		}
		position, ok := g.GetPosition(source.PositionID)
		if !ok {
			return true
		}

		return position.Rank == 0
	},
}
