package instance

import (
	"fmt"
	"grimdark/internal/game"
	"grimdark/internal/game/actions"
	"grimdark/internal/game/effects"
	"grimdark/internal/game/weapons"
)

func SetupGame(g *game.Game, user game.User) {
	effect := game.EffectSource(game.EffectPriorityStages, func(g game.Game, a game.Actor, ctx game.Context) game.Actor {
		a.Stages[game.Accuracy] = a.Stages[game.Accuracy] + 1
		a.Stages[game.CriticalChance] = a.Stages[game.CriticalChance] + 1
		a.Stages[game.CriticalDamage] = a.Stages[game.CriticalDamage] - 1
		a.AffinityResistance[game.Kinetic] += 1
		return a
	})
	effect.Triggers = []game.Trigger{
		{
			On:       game.OnDamageRecieve,
			Validate: game.TriggerTargetMatchesModifierParent,
			Action: game.Action{
				Resolve: func(g *game.Game, ctx game.Context, this game.ActionContext) []game.Transaction {
					targets := g.GetTargets(ctx)
					fmt.Println(game.OnDamageRecieve)
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

	bypass := effects.StagesResetWhere(func(g game.Game, a game.Actor, ctx game.Context) bool {
		active_context := g.State().ActiveContext
		if active_context == nil {
			return false
		}

		if active_context.SourceID == ctx.ParentID {
			return active_context.HasTarget(a)
		}

		return false
	})
	bypass.Name = "bypass effect"
	bypass_aux := effects.AuxResetWhere(func(g game.Game, a game.Actor, ctx game.Context) bool {
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
	player.ID = user.ID
	player.User = user
	if len(g.State().Players) > 0 {
		player = g.Base().Players[0]
	}
	max_def := game.NewActorDef()
	max_def.Name = "Max"
	max_def.Affinities = map[game.Affinity]struct{}{
		game.Fire: {},
	}
	max := game.NewActor(player.ID, max_def)
	max.Effects = []game.Effect{bypass, bypass_aux, effects.Intimidate()}
	max.Weapon = &weapons.SlashSword

	katie_def := game.NewActorDef()
	katie_def.Name = "Katie"
	katie_def.Affinities = map[game.Affinity]struct{}{
		game.Cryo:   {},
		game.Arcane: {},
	}
	katie_def.Effects = []game.Effect{effect, effect2}

	katie := game.NewActor(player.ID, katie_def)
	katie.AuxStats[game.Speed] = 10
	if len(g.State().Players) == 0 {
		g.AddPlayers(player)
	}
	katie.Actions = []game.Action{
		actions.SwordsDance,
	}
	katie.Weapon = &weapons.SlashSword
	katie.AffinityImmunities = map[game.Affinity]float64{
		game.Kinetic: 0,
	}

	gabe_def := game.NewActorDef()
	gabe_def.Name = "gabe"
	gabe_def.Affinities = map[game.Affinity]struct{}{
		game.Cryo:   {},
		game.Arcane: {},
	}
	gabe_def.Effects = []game.Effect{effect, effect2}

	gabe := game.NewActor(player.ID, gabe_def)
	if len(g.State().Players) == 0 {
		g.AddPlayers(player)
	}
	gabe.Actions = []game.Action{
		actions.SwordsDance,
	}
	gabe.Weapon = &weapons.SlashSword
	gabe.AffinityImmunities = map[game.Affinity]float64{
		game.Kinetic: 0,
	}

	g.AddActor(max)
	g.AddActor(katie)
	g.AddActor(gabe)
}
