package actors

import "grimdark/internal/game"

func NewBloodAngel() game.ActorDef {
	def := game.NewActorDef()
	def.Name = "Blood Angel"
	def.Affinities = map[game.Affinity]struct{}{
		game.Arcane:  {},
		game.Kinetic: {},
	}
	def.Effects = []game.Effect{}
	def.SpriteURL = "/img/spm2.png"

	return def
}
