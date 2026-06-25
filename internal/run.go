package main

import (
	"fmt"
	"grimdark/internal/game"
	"grimdark/internal/game/effects"
)

func main() {
	effect := game.EffectSource(game.EffectPriorityStages, func(g game.Game, a game.Actor, ctx game.Context) game.Actor {
		a.Stages[game.Evasion] = a.Stages[game.Evasion] + 1
		// a.AffinityImmunities[game.Kinetic] = 0
		a.AffinityResistance[game.Kinetic] += 1
		return a
	})
	effect.Triggers = []game.Trigger{
		{
			On:       game.OnDamageRecieve,
			Validate: game.TriggerTargetMatchesModifierParent,
			Action: game.Action{
				Resolve: func(g *game.Game, ctx game.Context, this game.ActionContext) []game.Transaction {
					fmt.Println("ON DAMAGE TRIGGER:")
					return this.Done()
				},
			},
		},
	}

	bypass := effects.StatsResetWhere(func(g game.Game, a game.Actor, ctx game.Context) bool {
		active_context := g.State().ActiveContext
		if active_context == nil {
			return false
		}

		if active_context.SourceID == ctx.ParentID {
			return active_context.HasTarget(a)
		}

		return false
	})

	player := game.NewPlayer()
	max_def := game.NewActorDef()
	max_def.Name = "Max"
	max_def.Affinities = map[game.Affinity]struct{}{
		game.Fire: {},
	}
	max := game.NewActor(player.ID, max_def)
	max.Effects = []game.Effect{bypass}

	katie_def := game.NewActorDef()
	katie_def.Name = "Katie"
	katie_def.Affinities = map[game.Affinity]struct{}{
		game.Cryo:   {},
		game.Arcane: {},
	}
	katie_def.Effects = []game.Effect{effect}

	katie := game.NewActor(player.ID, katie_def)
	katie.Aux[game.MartialDefense] = 10

	g := game.NewGame()
	g.AddPlayers(player)
	g.AddActor(max)
	g.AddActor(katie)

	temp_player, _ := g.GetPlayer(player.ID)
	g.SetPosition(max.ID, temp_player.GetOpenPosition())
	temp_player, _ = g.GetPlayer(player.ID)
	g.SetPosition(katie.ID, temp_player.GetOpenPosition())

	slash := game.Action{
		Config: game.ActionConfig{
			Name:         "Slash",
			Affinity:     game.Kinetic,
			Stat:         game.Melee,
			Accuracy:     game.P(0.98),
			Power:        70,
			Recoil:       0.12,
			Hits:         1,
			CritChance:   0.05,
			CritModifier: 1.5,
		},
		Resolve: game.BasicAttack(game.AttackConfig{
			OnSuccessResult: game.AddResultEffects(
				0.5,
				effects.StatDownTargets(game.Speed, 1),
			),
		}),
		ValidateContext:  game.ContextTargetLength(1),
		TargetsPredicate: game.CombineFilters(game.ActiveActors, game.Enemies),
	}
	swords_dance := game.Action{
		Config: game.ActionConfig{
			Name:     "Swords Dance",
			Affinity: game.Kinetic,
		},
		Resolve: game.AddSourceEffects(
			1,
			effects.StatUpSource(game.Speed, 1),
			effects.StatUpSource(game.Melee, 1),
		),
		ValidateContext:  game.TrueGameFilter,
		TargetsPredicate: game.NoneActors,
	}

	g.PushCommand(swords_dance.Bind(game.MakeContextFrom(katie)))
	g.PushCommand(slash.Bind(game.MakeContextFor(max, katie)))
	g.PushCommand(slash.Bind(game.MakeContextFor(katie, max)))
	g.Flush()

	katie, _ = g.GetActor(katie.ID)
	max, _ = g.GetActor(max.ID)

	fmt.Println(katie.AffinityResistance)
	fmt.Println(max.AffinityResistance)
}
