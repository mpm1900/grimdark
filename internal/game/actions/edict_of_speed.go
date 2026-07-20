package actions

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

var EdictOfSpeed = game.Action{
	ID:   uuid.MustParse("019f80f4-fa77-74ea-888c-2cc10b0d2051"),
	Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
	Config: game.ActionConfig{
		Name:        "Edict Of Speed",
		Description: "Doubles team's Speed for 5 turns.",
		Affinity:    game.Lightning,
		TargetCount: 0,
	},
	Resolve: game.AddGlobalEffects(
		game.StatusConfig{},
		1,
		edictOfSpeedEffect(),
	),
	ValidateContext:  game.TrueGameFilter,
	TargetsPredicate: game.NoneActors,
	ActiveCheck: func(source game.Actor) bool {
		if source.WeaponL != nil && source.WeaponR != nil {
			return source.WeaponL.ID == source.WeaponR.ID
		}
		return false
	},
}

func speedUp(g *game.Game, a game.Actor, ctx game.Context) game.Actor {
	a.Stats[game.Speed] = a.Stats[game.Speed] * 2
	return a
}
func edictOfSpeedEffect() game.Effect {
	effect := game.EffectAllies(game.EffectPriorityPostStagesStats, speedUp)
	effect.ID = uuid.MustParse("019f80fe-f95c-7214-93f9-3415769f408b")
	effect.Name = "Edict Of Speed"
	effect.Description = "Speed x2.0."
	effect.Duration = game.P(6)
	effect.CheckSuccess = game.EffectGainWhereOnSuccess(
		game.Allies,
	)
	return effect
}
