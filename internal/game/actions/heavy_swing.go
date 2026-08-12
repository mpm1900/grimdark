package actions

import (
	"grimdark/internal/game"
	"grimdark/internal/game/effects"

	"github.com/google/uuid"
)

func HeavySwing() game.Action {
	id := uuid.MustParse("019f8b1e-d47f-7863-981f-bf17defd2135")

	return game.Action{
		ID:   id,
		Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
		Entity: game.MakeEntity(
			id,
			"Heavy Swing",
			"This action lowers the user's Martial Defense stat. This action is only usable from 1st position.",
		),
		Config: game.ActionConfig{
			Affinity:     game.Physical,
			Stat:         game.Melee,
			Power:        110,
			Accuracy:     game.P(0.9),
			CritStage:    0,
			CritModifier: 1.5,
			TargetCount:  1,
			Range:        game.P(1),
			Priority:     game.ActionPriorityDefault,
		},
		Resolve: game.MakeAttack(game.AttackConfig{
			OnSuccessResult: func(g *game.Game, context game.Context, this *game.ActionContext, result game.DamageResult) {
				mod_ctx := game.MakeModifierContext(this.Source, this.Source)
				this.Push(game.AddModifiers(effects.StatDownSource(game.MartialDefense, 1).Bind(mod_ctx)).Bind(mod_ctx))
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
