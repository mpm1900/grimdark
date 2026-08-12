package actions

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

func RetreatingShot() game.Action {
	id := uuid.MustParse("019fce3e-14f8-7569-92c0-54a77231bbf4")

	return game.Action{
		ID:   id,
		Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
		Entity: game.MakeEntity(
			id,
			"Retreating Shot",
			"The user retreats after a successful attack. This action is only usable from 3rd position.",
		),
		Config: game.ActionConfig{
			Affinity:     game.Physical,
			Stat:         game.Ranged,
			Power:        50,
			Accuracy:     game.P(0.9),
			CritStage:    0,
			CritModifier: 1.5,
			TargetCount:  1,
			Range:        game.P(1),
		},
		Resolve: game.MakeAttack(game.AttackConfig{
			OnSuccess: func(g *game.Game, context game.Context, this *game.ActionContext) {
				this.Push(game.SetPositionSource(uuid.Nil).Bind(context))
			},
		}),
		ValidateContext:  game.ContextTargetLength(1),
		TargetsPredicate: game.CombineFilters(game.ActiveActors, game.OtherActors),
		DisabledCheck: func(g *game.Game, source game.Actor) bool {
			position, ok := g.GetPosition(source.PositionID)
			if !ok {
				return true
			}

			ctx := game.MakeContextFrom(source)
			inactive_allies := g.FindActors(game.CombineFilters(
				game.InactiveActors,
				game.AliveActors,
				game.Allies,
			), ctx)

			if len(inactive_allies) == 0 {
				return true
			}

			return position.Rank != 2
		},
	}
}
