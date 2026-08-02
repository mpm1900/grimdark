package actors

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

var All = map[uuid.UUID]game.Class{
	Templar.ID:      Templar,
	Prophet.ID:      Prophet,
	Paladin.ID:      Paladin,
	SisterOfFire.ID: SisterOfFire,
	Cultist.ID:      Cultist,
	Champion.ID:     Champion,
	Rogue.ID:        Rogue,
	Vicar.ID:        Vicar,
	Sensor.ID:       Sensor,
}

func HydrateActorClass(id uuid.UUID) (game.Class, bool) {
	class, ok := All[id]
	return class, ok
}

func HydrateActorConfig(config game.ActorConfig) (*game.Actor, bool) {
	if config.Class == nil {
		return nil, false
	}

	class, ok := HydrateActorClass(*config.Class)
	if !ok {
		return nil, false
	}

	return game.NewActor(class, config), true
}

func ApplyTeamConfig(g *game.Game, playerID uuid.UUID, config game.TeamConfig) {
	for _, a_config := range config.Actors {
		actor, ok := HydrateActorConfig(a_config)
		if ok {
			actor.PlayerID = playerID
			g.AddActors(*actor)
		}
	}
}
