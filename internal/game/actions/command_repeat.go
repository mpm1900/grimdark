package actions

import (
	"grimdark/internal/game"
	"slices"

	"github.com/google/uuid"
)

func repeat() game.Effect {
	effect := game.EffectTargets(game.EffectPriorityActions, func(g *game.Game, a game.Actor, ctx game.Context) game.Actor {
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
	effect.Name = "Command: Repeat"
	effect.Description = "This actor must repeat their last used action."
	effect.Duration = game.P(5)
	effect.CheckSuccess = game.EffectGainTargetsOnSuccess
	effect.CheckFailure = func(g *game.Game, e game.Effect, ctx game.Context) {
		g.PushLogMeta(game.NewLog(
			"$effect$ failed.",
			map[string]string{
				"$effect$": e.Name,
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

var CommandRepeat = game.Action{
	ID:   uuid.MustParse("019fb3c4-dccb-756a-9d75-3feb46ccba19"),
	Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
	Config: game.ActionConfig{
		Name:        "Command: Repeat",
		Description: "Target must repeat their last used action for 5 turns.",
		Affinity:    game.Psychic,
		TargetCount: 1,
	},
	Resolve: func(g *game.Game, ctx game.Context, this game.ActionContext) []game.Transaction {
		resolve := game.AddTargetsEffects(game.StatusConfig{}, ctx, repeat())
		return resolve(g, ctx, this)
	},
	ValidateContext:  game.ContextTargetLength(1),
	TargetsPredicate: game.OtherActors,
}
