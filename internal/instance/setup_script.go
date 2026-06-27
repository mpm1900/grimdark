package instance

import (
	"fmt"
	"grimdark/internal/game"
	"grimdark/internal/game/effects"

	"github.com/google/uuid"
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

	slash := game.Action{
		ID: uuid.New(),
		Config: game.ActionConfig{
			Name:         "Slash",
			Affinity:     game.Kinetic,
			Stat:         game.Melee,
			Accuracy:     game.P(0.98),
			Power:        70,
			Lifesteal:    0.12,
			Hits:         1,
			CritChance:   0.05,
			CritModifier: 1.5,
		},
		Resolve: game.BasicAttack(game.AttackConfig{
			OnSuccessResult: func(g game.Game, context game.Context, this *game.ActionContext, result game.DamageResult) {
				game.AddResultEffects(
					0.5,
					effects.StatDownTargets(game.Speed, 1),
				)(g, context, this, result)
				game.AddResultEffects(
					0.5,
					effects.StaggerTargets(),
				)(g, context, this, result)
			},
		}),
		ValidateContext:  game.ContextTargetLength(1),
		TargetsPredicate: game.CombineFilters(game.ActiveActors, game.Enemies),
	}
	swords_dance := game.Action{
		ID: uuid.New(),
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
	max.Actions = []game.Action{slash}

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
	katie.Actions = []game.Action{slash, swords_dance}

	g.AddActor(max)
	g.AddActor(katie)

	temp_player, _ := g.GetPlayer(player.ID)
	g.SetPosition(max.ID, temp_player.GetOpenPosition())
	temp_player, _ = g.GetPlayer(player.ID)
	g.SetPosition(katie.ID, temp_player.GetOpenPosition())
}
