package actors

import (
	"grimdark/internal/game"
	"grimdark/internal/game/effects"

	"github.com/google/uuid"
)

func NewUltramarine() game.ActorDef {
	bypass := effects.StagesResetWhere(func(g game.Game, a game.Actor, ctx game.Context) bool {
		active_context := g.State().ActiveContext
		if active_context == nil {
			return false
		}

		if active_context.ActionID != uuid.Nil && active_context.SourceID == ctx.ParentID {
			return active_context.HasTarget(a)
		}

		return false
	})
	bypass.Name = "bypass effect"

	def := game.NewActorDef()
	def.Name = "Ultramarine"
	def.Affinities = map[game.Affinity]struct{}{
		game.Kinetic: {},
	}
	def.Effects = []game.Effect{bypass, effects.Intimidate()}
	def.SpriteURL = "/img/spm.png"

	return def
}
