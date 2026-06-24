package main

import (
	"fmt"
	"grimdark/internal/game"

	"github.com/google/uuid"
	"github.com/k0kubun/pp/v3"
)

func main() {
	playerID := uuid.New()
	max_def := game.NewActorDef()
	max_def.Affinities = []game.Affinity{game.Fire}
	max := game.NewActor(playerID, max_def)
	max.Name = "Max"

	katie_def := game.NewActorDef()
	katie_def.Affinities = []game.Affinity{game.Cryo, game.Arcane}
	katie := game.NewActor(playerID, katie_def)
	katie.Name = "Katie"
	katie.Aux[game.MartialDefense] = 10

	effect := game.EffectTargets(game.EffectPriorityStages, func(a game.Actor, ctx game.Context) game.Actor {
		a.Stages[game.Melee] = a.Stages[game.Melee] + 1
		a.Stages[game.MartialDefense] = a.Stages[game.MartialDefense] + 1
		a.Stages[game.Speed] = a.Stages[game.Speed] - 1
		a.Stages[game.Evasion] = a.Stages[game.Evasion] + 1
		a.AffinityResistance[game.Kinetic] = a.AffinityResistance[game.Kinetic] + 1
		return a
	})

	effect.Triggers = []game.Trigger{
		{
			On:       game.OnDamageRecieve,
			Validate: game.TriggerTargetMatchesModifierParent,
			Action: game.Action{
				Resolve: func(g game.Game, ctx game.Context, this game.ActionContext) []game.Transaction {
					fmt.Println("hur dur trigger doing things")
					return []game.Transaction{}
				},
			},
		},
	}

	g := game.NewGame()
	context := game.MakeContextFor(max, katie)
	g.AddActors(max, katie)

	slash := game.Action{
		Config: game.ActionConfig{
			Name:         "Slash",
			Affinity:     game.Kinetic,
			Stat:         game.Melee,
			Power:        70,
			Accuracy:     80,
			CritChance:   5,
			CritModifier: 1.5,
		},
		Resolve: game.BasicAttack,
	}

	swords_dance := game.Action{
		Config: game.ActionConfig{
			Name:           "Swords Dance",
			BypassAccuracy: true,
		},
		Resolve: func(g game.Game, ctx game.Context, this game.ActionContext) []game.Transaction {
			this.Push(game.AddModifiers(effect.Bind(ctx)).Bind(game.NewContext()))

			return this.Done()
		},
	}

	g.PushCommand(swords_dance.Bind(game.MakeContextFor(katie, katie)))
	g.PushCommand(slash.Bind(context))
	g.Flush()

	katie, _ = g.GetActor(katie.ID)
	pp.Print(katie)
}
