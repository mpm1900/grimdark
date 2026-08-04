package items

import (
	"grimdark/internal/game"
	"grimdark/internal/game/effects"

	"github.com/google/uuid"
)

var CursedBoots = cursedBoots()
var CursedRing = cursedRing()

func cursedBoost(stat game.Stat) game.Effect {
	effect := game.EffectParent(game.EffectPriorityPostStagesStats, func(g *game.Game, a game.Actor, ctx game.Context) game.Actor {
		a.Stats[stat] *= 1.5
		return a
	})

	return effect
}

func cursedItem(effect game.Effect) game.Item {
	return game.Item{
		ID:          uuid.New(),
		Name:        effect.Name,
		Description: effect.Description,
		Effects: []game.Effect{
			effects.ChoiceLocked(),
			effect,
		},
	}
}

func cursedRing() game.Item {
	effect := cursedBoost(game.Melee)

	effect.Name = "Cursed Strength"
	effect.Description = "This actor's Melee is increased by 1.5x."
	item := cursedItem(effect)
	item.ID = uuid.MustParse("019fca3b-2cee-76f8-8f07-b576848c4026")
	item.Name = "Cursed Ring"
	item.Description = "This actor's Melee is increased by 1.5x but can only use a single action."
	return item
}

func cursedBoots() game.Item {
	effect := cursedBoost(game.Speed)
	effect.Name = "Cursed Speed"
	effect.Description = "This actor's Speed is increased by 1.5x."
	item := cursedItem(effect)
	item.ID = uuid.MustParse("019fca52-9a89-77f8-a58d-8e5933f5e291")
	item.Name = "Cursed Boots"
	item.Description = "This actor's Speed is increased by 1.5x but can only use a single action."
	return item
}
