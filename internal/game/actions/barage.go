package actions

import (
	"grimdark/internal/game"
	"grimdark/internal/game/effects"

	"github.com/google/uuid"
)

func Barage() game.Action {
	id := uuid.MustParse("019fc4ad-43b4-7648-be4b-6d6444f33aa2")

	return game.Action{
		ID:   id,
		Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
		Entity: game.MakeEntity(
			id,
			"Barage",
			"Damages all enemy actors. If there are no misses, this action also lowers each target's Evasion stats.",
		),
		Config: game.ActionConfig{
			Affinity:     game.Physical,
			Stat:         game.Ranged,
			Accuracy:     game.P(0.80),
			Power:        40,
			CritStage:    0,
			CritModifier: 1.5,
			TargetCount:  3,
		},
		Resolve: game.MakeAttack(game.AttackConfig{
			OnSuccess: func(g *game.Game, context game.Context, this *game.ActionContext) {
				targets := g.GetTargets(context)
				for _, t := range targets {
					game.AddEffectsTarget(
						1.0,
						t,
						effects.StatDownTargets(game.Evasion, 1),
					)(g, context, this)
				}
			},
		}),
		MapContext:       game.CtxToAllEnemies(),
		ValidateContext:  game.ContextTargetLength(0),
		TargetsPredicate: game.NoneActors,
		DisabledCheck: func(g *game.Game, source game.Actor) bool {
			position, ok := g.GetPosition(source.PositionID)
			if !ok {
				return true
			}

			return position.Rank == 0
		},
	}
}
