package weapons

import (
	"grimdark/internal/game"
	"grimdark/internal/game/actions"
	"grimdark/internal/game/effects"

	"github.com/google/uuid"
)

func FireTome() game.Weapon {
	id := uuid.MustParse("019f90ba-d85c-7a42-b2be-223a79131c2b")

	return game.Weapon{
		Item: game.Item{
			ID: id,
			Entity: game.MakeEntity(
				id,
				"Fire Tome",
				"A tome of prayers to the First Flame, Uriel.",
			),
			Effects: []game.Effect{
				effects.Elamentalist(),
				effects.AffinityImmunity(game.Fire),
			},
		},
		Actions: []game.Action{
			actions.Firestorm(),
			actions.Ignite(),
			actions.Immolate(),
			actions.Protect(),
			actions.SacredFlame(),
			actions.Wildfire(),
		},

		OffsetStats: map[game.Stat]float64{
			game.Special:        32,
			game.SpecialDefense: 32,
		},
		Weight:     2,
		WeaponType: game.WeaponTypeTome,
	}
}
