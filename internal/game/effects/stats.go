package effects

import (
	"fmt"
	"grimdark/internal/game"

	"github.com/google/uuid"
)

func StatChangeActor(stat game.Stat, amount int) game.Updater[game.Actor] {
	return func(g *game.Game, a game.Actor, ctx game.Context) game.Actor {
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
	if amount == 1 {
		effect.Description = fmt.Sprintf("Raises %s by 1 stage.", stat)
	} else {
		effect.Description = fmt.Sprintf("Raises %s by %d stages.", stat, amount)
	}
	effect.CheckSuccess = game.EffectGainSourceOnSuccess
	effect.ID = StatUpIDs[stat]

	return effect
}
func StatDownSource(stat game.Stat, amount int) game.Effect {
	effect := game.EffectSource(game.EffectPriorityStages, StatChangeActor(stat, -amount))
	effect.Name = fmt.Sprintf("%s down", stat)
	if amount == 1 {
		effect.Description = fmt.Sprintf("Lowers %s by 1 stage.", stat)
	} else {
		effect.Description = fmt.Sprintf("Lowers %s by %d stages.", stat, amount)
	}
	effect.CheckSuccess = game.EffectGainSourceOnSuccess
	effect.ID = StatDownIDs[stat]

	return effect
}
func StatUpTargets(stat game.Stat, amount int) game.Effect {
	effect := game.EffectTargets(game.EffectPriorityStages, StatChangeActor(stat, amount))
	effect.Name = fmt.Sprintf("%s up", stat)
	if amount == 1 {
		effect.Description = fmt.Sprintf("This actor's %s stat is raised by 1 stage.", stat)
	} else {
		effect.Description = fmt.Sprintf("This actor's %s stat is raised by %d stages.", stat, amount)
	}
	effect.CheckSuccess = game.EffectGainTargetsOnSuccess
	effect.ID = StatUpIDs[stat]

	return effect
}
func StatDownTargets(stat game.Stat, amount int) game.Effect {
	effect := game.EffectTargets(game.EffectPriorityStages, StatChangeActor(stat, -amount))
	effect.Name = fmt.Sprintf("%s down", stat)
	if amount == 1 {
		effect.Description = fmt.Sprintf("This actor's %s stat is lowered by 1 stage.", stat)
	} else {
		effect.Description = fmt.Sprintf("This actor's %s stat is lowered by %d stages.", stat, amount)
	}
	effect.CheckSuccess = game.EffectGainTargetsOnSuccess
	effect.ID = StatDownIDs[stat]

	return effect
}

func StagesResetWhere(where game.Filter[game.Actor]) game.Effect {
	effect := game.EffectActorsWhere(game.EffectPriorityStagesOverwrite, where, func(g *game.Game, a game.Actor, ctx game.Context) game.Actor {
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
