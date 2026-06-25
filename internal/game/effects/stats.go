package effects

import (
	"fmt"
	"grimdark/internal/game"
)

func StatChangeSourceOnSuccess(g *game.Game, this game.Effect, ctx game.Context) {
	source, ok := g.GetSource(ctx)
	if !ok {
		return
	}

	g.PushLog(game.NewLog(
		"$source$ gained $effect$",
		map[string]string{
			"$source$": source.Name,
			"$effect$": this.Name,
		},
	), ctx)
}
func StatChangeTargetsOnSuccess(g *game.Game, this game.Effect, ctx game.Context) {
	targets := g.GetTargets(ctx)
	for _, target := range targets {
		g.PushLog(game.NewLog(
			"$target$ gained $effect$",
			map[string]string{
				"$target$": target.Name,
				"$effect$": this.Name,
			},
		), ctx)
	}
}

func StatUpSource(stat game.Stat, amount int) game.Effect {
	effect := game.EffectSource(game.EffectPriorityStages, func(a game.Actor, ctx game.Context) game.Actor {
		a.Stages[stat] += amount
		return a
	})

	effect.Name = fmt.Sprintf("%s up", stat)
	effect.OnSuccess = StatChangeSourceOnSuccess

	return effect
}
func StatDownSource(stat game.Stat, amount int) game.Effect {
	effect := game.EffectSource(game.EffectPriorityStages, func(a game.Actor, ctx game.Context) game.Actor {
		a.Stages[stat] -= amount
		return a
	})

	effect.Name = fmt.Sprintf("%s down", stat)
	effect.OnSuccess = StatChangeSourceOnSuccess

	return effect
}
func StatUpTargets(stat game.Stat, amount int) game.Effect {
	effect := game.EffectTargets(game.EffectPriorityStages, func(g game.Game, a game.Actor, ctx game.Context) game.Actor {
		a.Stages[stat] += amount
		return a
	})

	effect.Name = fmt.Sprintf("%s up", stat)
	effect.OnSuccess = StatChangeTargetsOnSuccess

	return effect
}
func StatDownTargets(stat game.Stat, amount int) game.Effect {
	effect := game.EffectTargets(game.EffectPriorityStages, func(g game.Game, a game.Actor, ctx game.Context) game.Actor {
		a.Stages[stat] -= amount
		return a
	})

	effect.Name = fmt.Sprintf("%s down", stat)
	effect.OnSuccess = StatChangeTargetsOnSuccess

	return effect
}
