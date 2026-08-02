package weapons

import (
	"grimdark/internal/game"
	"grimdark/internal/game/actions"
	"grimdark/internal/game/effects"

	"github.com/google/uuid"
)

var FireTome = game.Weapon{
	Item: game.Item{
		ID:          uuid.MustParse("019f90ba-d85c-7a42-b2be-223a79131c2b"),
		Name:        "Fire Tome",
		Description: "A tome of prayers to the first flame, Gabriel.",
		Effects: []game.Effect{
			effects.Elamentalist(),
			effects.AffinityImmunity(game.Fire),
		},
	},
	Actions: []game.Action{
		actions.ArcaneRitual,
		actions.Firestorm,
		actions.Ignite,
		actions.Immolate,
		actions.Protect,
		actions.SacredFlame,
		actions.Wildfire,
	},

	OffsetStats: map[game.Stat]float64{
		game.Special:        32,
		game.SpecialDefense: 32,
	},
	Weight:     2,
	WeaponType: game.WeaponTypeTome,
}
