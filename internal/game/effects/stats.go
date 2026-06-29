package effects

import (
	"fmt"
	"grimdark/internal/game"

	"github.com/google/uuid"
)

func StatChangeSourceOnSuccess(g *game.Game, this game.Effect, ctx game.Context) {
	source, ok := g.GetSource(ctx)
	if !ok {
		return
	}

	g.PushLog(game.NewLog(
		"$source$ gained $effect$.",
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
			"$target$ gained $effect$.",
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

var StatUpIDs = map[game.Stat]uuid.UUID{
	game.Melee:          uuid.New(),
	game.Ranged:         uuid.New(),
	game.Special:        uuid.New(),
	game.MartialDefense: uuid.New(),
	game.SpecialDefense: uuid.New(),
	game.Speed:          uuid.New(),
	game.Accuracy:       uuid.New(),
	game.Evasion:        uuid.New(),
}
var StatDownIDs = map[game.Stat]uuid.UUID{
	game.Melee:          uuid.New(),
	game.Ranged:         uuid.New(),
	game.Special:        uuid.New(),
	game.MartialDefense: uuid.New(),
	game.SpecialDefense: uuid.New(),
	game.Speed:          uuid.New(),
	game.Accuracy:       uuid.New(),
	game.Evasion:        uuid.New(),
}

func StatUpSource(stat game.Stat, amount int) game.Effect {
	effect := game.EffectSource(game.EffectPriorityStages, StatChangeActor(stat, amount))
	effect.Name = fmt.Sprintf("%s up", stat)
	effect.CheckSuccess = StatChangeSourceOnSuccess
	effect.ID = StatUpIDs[stat]

	return effect
}
func StatDownSource(stat game.Stat, amount int) game.Effect {
	effect := game.EffectSource(game.EffectPriorityStages, StatChangeActor(stat, -amount))
	effect.Name = fmt.Sprintf("%s down", stat)
	effect.CheckSuccess = StatChangeSourceOnSuccess
	effect.ID = StatDownIDs[stat]

	return effect
}
func StatUpTargets(stat game.Stat, amount int) game.Effect {
	effect := game.EffectTargets(game.EffectPriorityStages, StatChangeActor(stat, amount))
	effect.Name = fmt.Sprintf("%s up", stat)
	effect.CheckSuccess = StatChangeTargetsOnSuccess
	effect.ID = StatUpIDs[stat]

	return effect
}
func StatDownTargets(stat game.Stat, amount int) game.Effect {
	effect := game.EffectTargets(game.EffectPriorityStages, StatChangeActor(stat, -amount))
	effect.Name = fmt.Sprintf("%s down", stat)
	effect.CheckSuccess = StatChangeTargetsOnSuccess
	effect.ID = StatDownIDs[stat]

	return effect
}

func StagesResetWhere(where game.Filter[game.Actor]) game.Effect {
	effect := game.EffectActorsWhere(game.EffectPriorityStagesOverwrite, where, func(g game.Game, a game.Actor, ctx game.Context) game.Actor {
		for stat, _ := range a.Stages {
			a.Stages[stat] = 0
		}
		for aff, _ := range a.AffinityResistance {
			a.AffinityResistance[aff] = 0
		}
		for aff, _ := range a.AffinityDamage {
			a.AffinityDamage[aff] = 0
		}

		return a
	})

	return effect
}

func AuxResetWhere(where game.Filter[game.Actor]) game.Effect {
	effect := game.EffectActorsWhere(game.EffectPriorityAuxOverwrite, where, func(g game.Game, a game.Actor, ctx game.Context) game.Actor {
		for stat, _ := range a.AuxStats {
			a.AuxStats[stat] = 0
		}

		return a
	})

	return effect
}
