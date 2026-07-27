package weapons

import (
	"grimdark/internal/game"
	"grimdark/internal/game/actions"

	"github.com/google/uuid"
)

var Dagger = game.Weapon{
	Item: game.Item{
		ID:          uuid.MustParse("019fa1b1-9e2c-72b3-acc6-aa5afc4684ae"),
		Name:        "Dagger",
		Description: "Forged for warriors sworn to clergy of the healing blood church, these short blades are honed to unsurpassed sharpness and able to slice through plate armor and thick hides.",
		Effects:     []game.Effect{},
	},
	Actions: []game.Action{
		actions.Protect,
		actions.QuickStrike,
		actions.RecklessStrike,
	},
	OffsetStats: map[game.Stat]float64{
		game.Melee: 16,
		game.Speed: 16,
	},
	Weight:     1,
	WeaponType: game.WeaponTypeDagger,
}

// quick attack
// retreating attack
// dagger throw (2 range)
