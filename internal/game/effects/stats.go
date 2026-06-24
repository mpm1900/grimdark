package effects

import (
	"fmt"
	"grimdark/internal/game"
)

func SourceModifierLog(g game.Game, this game.Effect, ctx game.Context) (game.Log, bool) {
	source, ok := g.GetSource(ctx)
	if !ok {
		return game.Log{}, false
	}

	return game.NewLog(
		"$source$ gained $effect$",
		map[string]string{
			"$source$": source.Name,
			"$effect$": this.Name,
		},
	), true
}

func StatUpSource(stat game.Stat, amount int) game.Effect {
	effect := game.EffectSource(game.EffectPriorityStages, func(a game.Actor, ctx game.Context) game.Actor {
		a.Stages[stat] += amount
		return a
	})

	effect.Name = fmt.Sprintf("%s up", stat)
	effect.GetLog = SourceModifierLog

	return effect
}
