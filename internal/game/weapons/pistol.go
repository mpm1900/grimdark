package weapons

import (
	"grimdark/internal/game"
	"grimdark/internal/game/actions"

	"github.com/google/uuid"
)

var Pistol = game.Weapon{
	Item: game.Item{
		ID:          uuid.MustParse("019f5319-dffb-7d06-b6f3-af8dca62bffd"),
		Name:        "Pistol",
		Description: "A bolter pistol. Pew pew.",
		Effects:     []game.Effect{},
	},
	Actions: []game.Action{actions.Slash, actions.Blast, actions.SwordsDance},
	AuxStats: map[game.Stat]float64{
		game.Ranged: 20,
	},
	Hands:      1,
	WeaponType: game.WeaponTypePistol,
}
