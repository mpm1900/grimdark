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
	effect.Entity = game.MakeEntity(
		effect.ID,
		"Firestorm",
		"On turn end, all non-fire actors take 12% damage.",
	)
	effect.Triggers = append(effect.Triggers, game.Trigger{
		On: game.OnTurnEnd,
		Action: game.Action{
			Entity: effect.Entity,
			Resolve: func(g *game.Game, ctx game.Context, this game.ActionContext) []game.Transaction {
				targets := g.FindActors(firestormWhere, ctx)
				for _, target := range targets {
					dmg_context := game.MakeContextFor(this.Source, target)
					health := target.Stats[game.Health]
					amount := health * 0.12
					this.Push(game.DamageTargets(amount, false, false).Bind(dmg_context))
				}
				return this.Done()
			},
		},
	})
	effect.Duration = game.P(5)
	effect.CheckSuccess = game.EffectGainWhereOnSuccess(
		firestormWhere,
	)
	return effect
}

func Firestorm() game.Action {
	id := uuid.MustParse("019f9ccc-1f0c-7bc6-abf1-f8b38503c3dc")

	return game.Action{
		ID:   id,
		Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
		Entity: game.MakeEntity(
			id,
			"Firestorm",
			fmt.Sprintf("Applies Firestorm to the battlefield for 5 turns. (%s)", firestorm().Entity.Description),
		),
		Config: game.ActionConfig{
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
}
