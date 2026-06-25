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
	).Bind(ctx))
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
		).Bind(ctx))
	}
}

func StatChangeActor(stat game.Stat, amount int) game.Updater[game.Actor] {
	return func(g game.Game, a game.Actor, ctx game.Context) game.Actor {
		a.Stages[stat] += amount
		return a
	}
}

func StatUpSource(stat game.Stat, amount int) game.Effect {
	effect := game.EffectSource(game.EffectPriorityStages, StatChangeActor(stat, amount))
	effect.Name = fmt.Sprintf("%s up", stat)
	effect.CheckSuccess = StatChangeSourceOnSuccess

	return effect
}
func StatDownSource(stat game.Stat, amount int) game.Effect {
	effect := game.EffectSource(game.EffectPriorityStages, StatChangeActor(stat, -amount))
	effect.Name = fmt.Sprintf("%s down", stat)
	effect.CheckSuccess = StatChangeSourceOnSuccess

	return effect
}
func StatUpTargets(stat game.Stat, amount int) game.Effect {
	effect := game.EffectTargets(game.EffectPriorityStages, StatChangeActor(stat, amount))
	effect.Name = fmt.Sprintf("%s up", stat)
	effect.CheckSuccess = StatChangeTargetsOnSuccess

	return effect
}
func StatDownTargets(stat game.Stat, amount int) game.Effect {
	effect := game.EffectTargets(game.EffectPriorityStages, StatChangeActor(stat, -amount))
	effect.Name = fmt.Sprintf("%s down", stat)
	effect.CheckSuccess = StatChangeTargetsOnSuccess

	return effect
}

func StatsResetWhere(where game.Filter[game.Actor]) game.Effect {
	effect := game.EffectActorsWhere(game.EffectPriorityStagesOverwrite, where, func(g game.Game, a game.Actor, ctx game.Context) game.Actor {
		for stat, _ := range a.Stages {
			a.Stages[stat] = 0
		}

		for aff, _ := range a.AffinityResistance {
			a.AffinityResistance[aff] = 0
		}
		return a
	})

	return effect
}
