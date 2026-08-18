package effects

import (
	"grimdark/internal/game"
	"slices"

	"github.com/google/uuid"
)

func ChoiceLocked() game.Effect {
	effect := game.EffectTargets(game.EffectPriorityActionState, func(g *game.Game, a game.Actor, ctx game.Context) game.Actor {
		struggle_ID := game.Struggle().ID
		for _, action := range a.GetActions() {
			if a.Meta.LastUsedActionID == uuid.Nil {
				continue
			}
			if action.ID == a.Meta.LastUsedActionID {
				continue
			}

			a.UpdateActionState(action.ID, func(as game.ActionState) game.ActionState {
				as.IsDisabled = action.ID != struggle_ID && !slices.Contains(action.Tags, game.ATSystem)
				return as
			})
		}
		return a
	})
	effect.Entity = game.MakeEntity(
		effect.ID,
		"Choice Locked",
		"This actor must repeat their last used action.",
	)
	effect.CheckSuccess = game.EffectGainTargetsOnSuccess
	effect.CheckFailure = func(g *game.Game, e game.Effect, ctx game.Context) {
		g.PushLogMeta(game.NewLog(
			"$effect$ failed.",
			map[string]string{
				"$effect$": e.Entity.Name,
			},
		).Bind(game.NewContext()))
	}
	effect.Check = func(g *game.Game, ctx game.Context) bool {
		for _, t := range g.GetTargets(ctx) {
			if t.Meta.LastUsedActionID == uuid.Nil {
				return false
			}
		}
		return true
	}

	return effect
}
