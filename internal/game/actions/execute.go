package actions

import (
	"grimdark/internal/game"
	"grimdark/internal/game/effects"

	"github.com/google/uuid"
)

var Execute = game.Action{
	ID:   uuid.MustParse("019f8b18-358d-73b6-91a4-5ae8f516b113"),
	Tags: []game.ActionTag{game.ATActor, game.ATWeapon, game.ATMovement},
	Config: game.ActionConfig{
		Name:         "Execute",
		Description:  "Deals massive damage. User is stunned next turn unless this action killed the target. This action is only usable from 1st position.",
		Affinity:     game.Kinetic,
		Stat:         game.Melee,
		Power:        120,
		Accuracy:     game.P(0.9),
		Lifesteal:    0,
		Hits:         1,
		CritStage:    0,
		CritModifier: 1.5,
		TargetCount:  1,
		Range:        game.P(1),
		Priority:     game.ActionPriorityDefault,
	},
	Resolve: game.MakeAttack(game.AttackConfig{
		OnSuccessResult: func(g *game.Game, context game.Context, this *game.ActionContext, result game.DamageResult) {
			result_health := result.Target.GetRemainingHealth()
			if result.Damage < result_health {
				stun_ctx := game.MakeModifierContext(this.Source, this.Source)
				this.Push(game.AddModifiers(effects.StunTargets.Bind(stun_ctx)).Bind(stun_ctx))
			}
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
