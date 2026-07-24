package weapons

import (
	"grimdark/internal/game"
	"grimdark/internal/game/actions"

	"github.com/google/uuid"
)

var FireTome = game.Weapon{
	Item: game.Item{
		ID:          uuid.MustParse("019f90ba-d85c-7a42-b2be-223a79131c2b"),
		Name:        "Fire Tome",
		Description: "A tome of prayers to the god-king.",
		Effects:     []game.Effect{},
	},
	Actions: []game.Action{
		actions.Immolate,
		actions.Protect,
	},
	OffsetStats: map[game.Stat]float64{
		game.Special:        32,
		game.SpecialDefense: 32,
	},
	Weight:     2,
	WeaponType: game.WeaponTypeTome,
}
