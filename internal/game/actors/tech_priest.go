package actors

import "grimdark/internal/game"

func NewTechPriest() game.ActorDef {
	def := game.NewActorDef()
	def.Name = "Tech Priest"
	def.Affinities = map[game.Affinity]struct{}{
		game.Kinetic:   {},
		game.Lightning: {},
	}
	def.Effects = []game.Effect{}
	def.SpriteURL = "/img/thp.png"

	return def
}
