package actors

import (
	"grimdark/internal/game"
)

func NewSisterOfBattle() game.ActorDef {
	def := game.NewActorDef()
	def.Name = "Adepta Sororitas"
	def.Affinities = map[game.Affinity]struct{}{
		game.Fire:    {},
		game.Kinetic: {},
	}
	def.Effects = []game.Effect{}
	def.SpriteURL = "/img/sis.png"

	return def
}
