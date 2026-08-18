package actors

import (
	"grimdark/internal/game"
	"grimdark/internal/game/items"

	"github.com/google/uuid"
)

var All = []game.Class{
	Templar(),
	Prophet(),
	Paladin(),
	SisterOfFire(),
	SisterOfLight(),
	// SisterOfSacrifice(),
	Cultist(),
	Champion(),
	Rogue(),
	Vicar(),
	Seeker(),
	Bloodknight(),
	Inquisitor(),
	Wanderer(),
}

func HydrateActorClass(id uuid.UUID) (game.Class, bool) {
	for _, c := range All {
		if c.ID == id {
			return c, true
		}
	}

	return game.Class{}, false
}

func HydrateActorConfig(config game.ActorConfig) (*game.Actor, bool) {
	if config.Class == nil {
		return nil, false
	}

	class, ok := HydrateActorClass(*config.Class)
	if !ok {
		return nil, false
	}

	class.Options.Items = append(class.Options.Items, items.Global...)
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
