package items

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

func CorruptedNecklace() game.Item {
	id := uuid.MustParse("019fca79-fdc7-757d-9214-5a4952a86358")
	entity := game.MakeEntity(
		id,
		"Corrupted Necklace",
		"Increases Melee, Ranged, and Special stats by 1.3x. On turn end, this actor loses 10% HP.",
	)
	effect := game.EffectParent(game.EffectPriorityPostStagesStats, func(g *game.Game, a game.Actor, ctx game.Context) game.Actor {
		a.Stats[game.Melee] *= 1.3
		a.Stats[game.Ranged] *= 1.3
		a.Stats[game.Special] *= 1.3
		return a
	})
	effect.Triggers = append(effect.Triggers, game.Trigger{
		On: game.OnTurnEnd,
		Action: game.Action{
			Entity: entity,
			Resolve: func(g *game.Game, ctx game.Context, this game.ActionContext) []game.Transaction {
				dmg_context := game.MakeContextFor(this.Source, this.Source)
				health := this.Source.Stats[game.Health]
				amount := health * 0.10
				this.Push(game.DamageTargets(amount, false, false).Bind(dmg_context))
				return this.Done()
			},
		},
	})
	effect.Entity = entity

	return game.Item{
		ID:     id,
		Entity: entity,
		Effects: []game.Effect{
			effect,
		},
	}
}
