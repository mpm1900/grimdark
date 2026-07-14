package actors

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

var All = map[uuid.UUID]game.Class{
	BloodAngel.ID:     BloodAngel,
	NecronWarrior.ID:  NecronWarrior,
	SisterOfBattle.ID: SisterOfBattle,
	TechPriest.ID:     TechPriest,
	Ultramarine.ID:    Ultramarine,
}

func HydrateActorClass(id uuid.UUID) (game.Class, bool) {
	class, ok := All[id]
	return class, ok
}

func HydrateActorConfig(config game.ActorConfig) (game.Actor, bool) {
	if config.Class == nil {
		return game.Actor{}, false
	}

	class, ok := HydrateActorClass(*config.Class)
	if !ok {
		return game.Actor{}, false
	}

	return game.NewActor(class, config), true
}

func ApplyTeamConfig(g *game.Game, playerID uuid.UUID, config game.TeamConfig) {
	for _, a_config := range config.Actors {
		actor, ok := HydrateActorConfig(a_config)
		if ok {
			actor.PlayerID = playerID
			g.AddActors(actor)
		}
	}
}
