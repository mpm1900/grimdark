package items

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

var CorruptedNecklace = corruptedNecklace()

func corruptedNecklace() game.Item {
	effect := game.EffectParent(game.EffectPriorityPostStagesStats, func(g *game.Game, a game.Actor, ctx game.Context) game.Actor {
		a.Stats[game.Melee] *= 1.3
		a.Stats[game.Ranged] *= 1.3
		a.Stats[game.Special] *= 1.3
		return a
	})
	effect.Triggers = append(effect.Triggers, game.Trigger{
		On: game.OnTurnEnd,
		Action: game.Action{
			Config: game.ActionConfig{
				Name: "Corruption",
			},
			Resolve: func(g *game.Game, ctx game.Context, this game.ActionContext) []game.Transaction {
				dmg_context := game.MakeContextFor(this.Source, this.Source)
				health := this.Source.Stats[game.Health]
				amount := health * 0.10
				this.Push(game.DamageTargets(amount, false, false).Bind(dmg_context))
				return this.Done()
			},
		},
	})

	effect.Name = "Corrupted Necklace"
	effect.Description = "Increases Melee, Ranged, and Special stats by 1.3x. On turn end, this actor loses 10% HP."

	return game.Item{
		ID:          uuid.MustParse("019fca79-fdc7-757d-9214-5a4952a86358"),
		Name:        "Corrupted Necklace",
		Description: "Increases Melee, Ranged, and Special stats by 1.3x. On turn end, this actor loses 10% HP.",
		Effects: []game.Effect{
			effect,
		},
	}
}
