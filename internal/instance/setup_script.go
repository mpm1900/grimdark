package instance

import (
	"fmt"
	"grimdark/internal/game"
	"grimdark/internal/game/effects"
)

func SetupGame(g *game.Game) {
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
					fmt.Println(game.OnDamageRecieve)
					return this.Done()
				},
			},
		},
	}
	effect.Name = "test effect"

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
	if len(g.State().Players) > 0 {
		player = g.Base().Players[0]
	}
	max_def := game.NewActorDef()
	max_def.Name = "Max"
	max_def.Affinities = map[game.Affinity]struct{}{
		game.Fire: {},
	}
	max := game.NewActor(player.ID, max_def)
	max.Effects = []game.Effect{bypass, bypass_aux}

	katie_def := game.NewActorDef()
	katie_def.Name = "Katie"
	katie_def.Affinities = map[game.Affinity]struct{}{
		game.Cryo:   {},
		game.Arcane: {},
	}
	katie_def.Effects = []game.Effect{effect}

	katie := game.NewActor(player.ID, katie_def)
	katie.Aux[game.Speed] = 10
	if len(g.State().Players) == 0 {
		g.AddPlayers(player)
	}

	g.AddActor(max)
	g.AddActor(katie)

	temp_player, _ := g.GetPlayer(player.ID)
	g.SetPosition(max.ID, temp_player.GetOpenPosition())
	temp_player, _ = g.GetPlayer(player.ID)
	g.SetPosition(katie.ID, temp_player.GetOpenPosition())
}
