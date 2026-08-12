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

var StatChangeTag game.EffectTag = "stat-change"
var StatUpTag game.EffectTag = "stat-up"
var StatDownTag game.EffectTag = "stat-down"

// consider refactoring these to xyzParent
func StatUpSource(stat game.Stat, amount int) game.Effect {
	effect := game.EffectSource(game.EffectPriorityStages, StatChangeActor(stat, amount))
	effect.ID = StatUpIDs[stat]

	description := fmt.Sprintf("Raises %s by %d stages.", stat, amount)
	if amount == 1 {
		description = fmt.Sprintf("Raises %s stat.", stat)
	}
	effect.Entity = game.MakeEntity(
		effect.ID,
		fmt.Sprintf("%s up", stat),
		description,
	)

	effect.CheckSuccess = game.EffectGainSourceOnSuccess
	effect.SetTag(StatChangeTag)
	effect.SetTag(StatUpTag)

	return effect
}
func StatDownSource(stat game.Stat, amount int) game.Effect {
	effect := game.EffectSource(game.EffectPriorityStages, StatChangeActor(stat, -amount))
	effect.ID = StatDownIDs[stat]

	description := fmt.Sprintf("Lowers %s by %d stages.", stat, amount)
	if amount == 1 {
		description = fmt.Sprintf("Lowers %s stat.", stat)
	}
	effect.Entity = game.MakeEntity(
		effect.ID,
		fmt.Sprintf("%s down", stat),
		description,
	)

	effect.CheckSuccess = game.EffectGainSourceOnSuccess
	effect.SetTag(StatChangeTag)
	effect.SetTag(StatDownTag)

	return effect
}
func StatUpTargets(stat game.Stat, amount int) game.Effect {
	effect := game.EffectTargets(game.EffectPriorityStages, StatChangeActor(stat, amount))
	effect.ID = StatUpIDs[stat]

	description := fmt.Sprintf("Raises %s by %d stages.", stat, amount)
	if amount == 1 {
		description = fmt.Sprintf("Raises %s stat.", stat)
	}
	effect.Entity = game.MakeEntity(
		effect.ID,
		fmt.Sprintf("%s up", stat),
		description,
	)

	effect.CheckSuccess = game.EffectGainTargetsOnSuccess
	effect.SetTag(StatChangeTag)
	effect.SetTag(StatUpTag)

	return effect
}
func StatDownTargets(stat game.Stat, amount int) game.Effect {
	effect := game.EffectTargets(game.EffectPriorityStages, StatChangeActor(stat, -amount))
	effect.ID = StatDownIDs[stat]

	description := fmt.Sprintf("Lowers %s by %d stages.", stat, amount)
	if amount == 1 {
		description = fmt.Sprintf("Lowers %s stat.", stat)
	}
	effect.Entity = game.MakeEntity(
		effect.ID,
		fmt.Sprintf("%s down", stat),
		description,
	)

	effect.CheckSuccess = game.EffectGainTargetsOnSuccess
	effect.SetTag(StatChangeTag)
	effect.SetTag(StatDownTag)

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
	effect.SetTag(StatChangeTag)

	return effect
}
