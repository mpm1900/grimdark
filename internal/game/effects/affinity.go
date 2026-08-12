package effects

import (
	"fmt"
	"grimdark/internal/game"
)

func AffinityImmunity(affinity game.Affinity) game.Effect {
	effect := game.EffectParent(game.EffectPriorityImmunities, func(g *game.Game, a game.Actor, ctx game.Context) game.Actor {
		a.AffinityImmunities[affinity] = 0
		return a
	})
	effect.Entity = game.MakeEntity(
		effect.ID,
		fmt.Sprintf("%s Immunity", affinity),
		fmt.Sprintf("Immunity to all %s actions and damage.", affinity),
	)

	return effect
}
