package main

import (
	"fmt"
	"grimdark/internal/game"
	"grimdark/internal/game/effects"

	"github.com/google/uuid"
	"github.com/k0kubun/pp/v3"
)

func main() {
	playerID := uuid.New()
	max_def := game.NewActorDef()
	max_def.Affinities = map[game.Affinity]struct{}{
		game.Fire: {},
	}
	max := game.NewActor(playerID, max_def)
	max.Name = "Max"

	katie_def := game.NewActorDef()
	katie_def.Affinities = map[game.Affinity]struct{}{
		game.Cryo:   {},
		game.Arcane: {},
	}
	katie := game.NewActor(playerID, katie_def)
	katie.Name = "Katie"
	katie.Aux[game.MartialDefense] = 10

	effect := game.EffectTargets(game.EffectPriorityStages, func(a game.Actor, ctx game.Context) game.Actor {
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
					return this.Done()
				},
			},
		},
	}

	g := game.NewGame()
	g.AddActors(max, katie)
	g.AddModifiers(effect.Bind(game.MakeContextFor(katie, katie)))

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
		MapContext: func(g game.Game, ctx game.Context, this game.ActionContext) game.Context {
			c := ctx.CloneWithTargets(g.FindActors(game.ActiveActors, ctx))
			return c
		},
	}

	swords_dance := game.Action{
		Config: game.ActionConfig{
			Name:           "Swords Dance",
			BypassAccuracy: true,
		},
		Resolve: func(g game.Game, ctx game.Context, this game.ActionContext) []game.Transaction {
			this.Push(game.AddModifiers(
				effects.StatUpSource(game.Melee, 1).Bind(ctx),
				effects.StatUpSource(game.Speed, 1).Bind(ctx),
			).Bind(game.NewContext()))

			return this.Done()
		},
	}

	g.PushCommand(swords_dance.Bind(game.MakeContextFrom(katie)))
	g.PushCommand(slash.Bind(game.MakeContextFor(max, katie)))
	g.Flush()

	katie, _ = g.GetActor(katie.ID)
	pp.Print(katie)
}
