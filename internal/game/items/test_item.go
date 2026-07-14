package items

import (
	"fmt"
	"grimdark/internal/game"

	"github.com/google/uuid"
)

func testItem() game.Item {
	effect := game.EffectSource(game.EffectPriorityStages, func(g *game.Game, a game.Actor, ctx game.Context) game.Actor {
		a.Stages[game.Accuracy] = a.Stages[game.Accuracy] + 1
		a.Stages[game.CriticalChance] = a.Stages[game.CriticalChance] + 1
		a.Stages[game.CriticalDamage] = a.Stages[game.CriticalDamage] - 1
		a.AffinityResistance[game.Kinetic] += 1
		return a
	})
	effect.Triggers = []game.Trigger{
		{
			On:       game.OnModifierAdd,
			Validate: game.TriggerTargetMatchesModifierParent,
			Action: game.Action{
				Resolve: func(g *game.Game, ctx game.Context, this game.ActionContext) []game.Transaction {
					targets := g.GetTargets(ctx)
					mod, ok := g.GetModifier(ctx.ModifierID)
					if ok {
						fmt.Println(mod.Payload.Name, mod.ID, mod.Payload.ID)
						// this.Push(game.RemoveModifier(mod).Bind(ctx))
						// this.Push(game.ConsumeItem().Bind(ctx))

					}
					fmt.Println(game.OnModifierAdd)
					for _, t := range targets {
						fmt.Println(t.Name)
					}

					return this.Done()
				},
			},
		},
	}
	effect.Name = "test effect"

	effect2 := game.EffectSource(game.EffectPriorityPostStagesStats, func(g *game.Game, a game.Actor, ctx game.Context) game.Actor {
		a.Stats[game.Melee] = a.Stats[game.Melee] / 2
		// this below is the wrong priority, testing only
		a.Status = game.StatusBurned
		return a
	})
	effect2.Name = "burned"

	item := game.Item{
		ID:          uuid.MustParse("019f433b-dda3-7184-99af-c83e1d744977"),
		Name:        "test item",
		Description: "a test item",
		Effects:     []game.Effect{effect, effect2},
	}

	return item
}

var TestItem = testItem()
