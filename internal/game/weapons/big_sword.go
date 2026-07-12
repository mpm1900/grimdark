package weapons

import (
	"grimdark/internal/game"
	"grimdark/internal/game/actions"
	"grimdark/internal/game/effects"

	"github.com/google/uuid"
)

var BigSword = game.Weapon{
	Item: game.Item{
		ID:          uuid.MustParse("019f4a69-3324-70f6-80e7-aded9c2c1f13"),
		Name:        "Big Sword With a Long Name",
		Description: "A large sword, designed for slashing.",
		Effects:     []game.Effect{effects.Intimidate},
	},
	Actions: []game.Action{actions.Slash, actions.Blast, actions.SwordsDance},
	AuxStats: map[game.Stat]float64{
		game.Melee: 20,
	},
	Hands:      2,
	WeaponType: game.WeaponTypeBigSword,
}
