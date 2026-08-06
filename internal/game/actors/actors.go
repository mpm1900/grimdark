package actors

import (
	"grimdark/internal/game"
	"grimdark/internal/game/items"
	"maps"
	"slices"

	"github.com/google/uuid"
)

var All = map[uuid.UUID]game.Class{
	Templar.ID:           Templar,
	Prophet.ID:           Prophet,
	Paladin.ID:           Paladin,
	SisterOfFire.ID:      SisterOfFire,
	SisterOfLight.ID:     SisterOfLight,
	SisterOfSacrifice.ID: SisterOfSacrifice,
	Cultist.ID:           Cultist,
	Champion.ID:          Champion,
	Rogue.ID:             Rogue,
	Vicar.ID:             Vicar,
	Sensor.ID:            Sensor,
	Bloodknight.ID:       Bloodknight,
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

	class.Options.Items = append(class.Options.Items, slices.Collect(maps.Values(items.Global))...)
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
