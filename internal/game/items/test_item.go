package items

import (
	"fmt"
	"grimdark/internal/game"

	"github.com/google/uuid"
)

func TestItem() game.HeldItem {
	effect := game.EffectSource(game.EffectPriorityStages, func(g game.Game, a game.Actor, ctx game.Context) game.Actor {
		a.Stages[game.Accuracy] = a.Stages[game.Accuracy] + 1
		a.Stages[game.CriticalChance] = a.Stages[game.CriticalChance] + 1
		a.Stages[game.CriticalDamage] = a.Stages[game.CriticalDamage] - 1
		a.AffinityResistance[game.Kinetic] += 1
		return a
	})
	effect.Triggers = []game.Trigger{
		{
			On:       game.OnTurnEnd,
			Validate: game.TriggerModifierParentIsActive,
			Action: game.Action{
				Resolve: func(g *game.Game, ctx game.Context, this game.ActionContext) []game.Transaction {
					targets := g.GetTargets(ctx)
					fmt.Println(game.OnTurnEnd)
					for _, t := range targets {
						fmt.Println(t.Name)
					}

					return this.Done()
				},
			},
		},
	}
	effect.Name = "test effect"

	effect2 := game.EffectSource(game.EffectPriorityPostStagesStats, func(g game.Game, a game.Actor, ctx game.Context) game.Actor {
		a.Stats[game.Melee] = a.Stats[game.Melee] / 2
		// this below is the wrong priority, testing only
		a.Status = game.StatusBurned
		return a
	})
	effect2.Name = "burned"

	item := game.HeldItem{
		Item: game.Item{
			ID:          uuid.New(),
			Name:        "test item",
			Description: "a test item",
		},
		Effects: []game.Effect{effect, effect2},
	}

	return item
}
