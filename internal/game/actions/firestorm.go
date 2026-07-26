package actions

import (
	"fmt"
	"grimdark/internal/game"

	"github.com/google/uuid"
)

func firestormWhere(g *game.Game, a game.Actor, ctx game.Context) bool {
	_, ok := a.Class.Affinities[game.Fire]
	return !ok && game.ActiveActors(g, a, ctx)
}

func firestorm() game.Effect {
	effect := game.EffectActorsWhere(
		game.EffectPriorityTriggers,
		firestormWhere,
		func(g *game.Game, a game.Actor, ctx game.Context) game.Actor {
			return a
		},
	)

	effect.Triggers = append(effect.Triggers, game.Trigger{
		On: game.OnTurnEnd,
		Action: game.Action{
			Config: game.ActionConfig{
				Name: "Firestorm",
			},
			Resolve: func(g *game.Game, ctx game.Context, this game.ActionContext) []game.Transaction {
				targets := g.FindActors(firestormWhere, ctx)

				for _, target := range targets {
					dmg_context := game.MakeContextFor(this.Source, target)
					health := target.Stats[game.Health]
					amount := health * 0.12
					this.Push(game.DamageTargets(amount, false).Bind(dmg_context))
				}

				return this.Done()
			},
		},
	})

	effect.Name = "Firestorm"
	effect.Description = "On turn end, all non-fire actors take 12% damage."
	effect.Duration = game.P(5)
	effect.CheckSuccess = game.EffectGainWhereOnSuccess(
		firestormWhere,
	)
	return effect
}

var Firestorm = game.Action{
	ID:   uuid.MustParse("019f9ccc-1f0c-7bc6-abf1-f8b38503c3dc"),
	Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
	Config: game.ActionConfig{
		Name:        "Firestorm",
		Description: fmt.Sprintf("Applies Firestorm to the battlefield for 5 turns. (%s)", firestorm().Description),
		Affinity:    game.Fire,
		TargetCount: 0,
		Cooldown:    5,
	},
	Resolve: game.AddGlobalEffects(
		game.StatusConfig{},
		1,
		firestorm(),
	),
	ValidateContext:  game.TrueGameFilter,
	TargetsPredicate: game.NoneActors,
}
