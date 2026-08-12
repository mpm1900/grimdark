package items

import (
	"grimdark/internal/game"
	"grimdark/internal/game/effects"

	"github.com/google/uuid"
)

func cursedBoost(stat game.Stat) game.Effect {
	effect := game.EffectParent(game.EffectPriorityPostStagesStats, func(g *game.Game, a game.Actor, ctx game.Context) game.Actor {
		a.Stats[stat] *= 1.5
		return a
	})

	return effect
}

func cursedItem(effect game.Effect) game.Item {
	id := uuid.New()

	return game.Item{
		ID:     id,
		Entity: effect.Entity,
		Effects: []game.Effect{
			effects.ChoiceLocked(),
			effect,
		},
	}
}

func CursedRing() game.Item {
	effect := cursedBoost(game.Melee)
	effect.Entity = game.MakeEntity(
		effect.ID,
		"Cursed Strength",
		"This actor's Melee is increased by 1.5x.",
	)

	id := uuid.MustParse("019fca3b-2cee-76f8-8f07-b576848c4026")

	item := cursedItem(effect)
	item.ID = id
	item.Entity = game.MakeEntity(
		id,
		"Cursed Ring",
		"This actor's Melee is increased by 1.5x but can only use a single action.",
	)
	return item
}

func CursedBoots() game.Item {
	effect := cursedBoost(game.Speed)
	effect.Entity = game.MakeEntity(
		effect.ID,
		"Cursed Speed",
		"This actor's Speed is increased by 1.5x.",
	)

	id := uuid.MustParse("019fca52-9a89-77f8-a58d-8e5933f5e291")
	item := cursedItem(effect)
	item.ID = id
	item.Entity = game.MakeEntity(
		id,
		"Cursed Boots",
		"This actor's Speed is increased by 1.5x but can only use a single action.",
	)
	return item
}
