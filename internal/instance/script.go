package instance

import (
	"grimdark/internal/game"
	"grimdark/internal/game/effects"
)

func TestGame(g *game.Game) {
	max := g.FindActors(func(g game.Game, a game.Actor, ctx game.Context) bool {
		return a.Name == "Max"
	}, game.NewContext())[0]
	katie := g.FindActors(func(g game.Game, a game.Actor, ctx game.Context) bool {
		return a.Name == "Katie"
	}, game.NewContext())[0]

	slash := game.Action{
		Config: game.ActionConfig{
			Name:         "Slash",
			Affinity:     game.Kinetic,
			Stat:         game.Melee,
			Accuracy:     game.P(0.98),
			Power:        70,
			Lifesteal:    0.12,
			Hits:         1,
			CritChance:   0.05,
			CritModifier: 1.5,
		},
		Resolve: game.BasicAttack(game.AttackConfig{
			OnSuccessResult: func(g game.Game, context game.Context, this *game.ActionContext, result game.DamageResult) {
				game.AddResultEffects(
					0.5,
					effects.StatDownTargets(game.Speed, 1),
				)(g, context, this, result)
				/*
					game.AddResultEffects(
						0.5,
						effects.StaggerTargets(),
					)(g, context, this, result)
				*/
			},
		}),
		ValidateContext:  game.ContextTargetLength(1),
		TargetsPredicate: game.CombineFilters(game.ActiveActors, game.Enemies),
	}
	swords_dance := game.Action{
		Config: game.ActionConfig{
			Name:     "Swords Dance",
			Affinity: game.Kinetic,
		},
		Resolve: game.AddSourceEffects(
			1,
			effects.StatUpSource(game.Speed, 1),
			effects.StatUpSource(game.Melee, 1),
		),
		ValidateContext:  game.TrueGameFilter,
		TargetsPredicate: game.NoneActors,
	}

	g.PushCommand(swords_dance.Bind(game.MakeContextFrom(katie)))
	g.PushCommand(slash.Bind(game.MakeContextFor(max, katie)))
	g.PushCommand(slash.Bind(game.MakeContextFor(katie, max)))
}
