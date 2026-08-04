package weapons

import (
	"grimdark/internal/game"
	"grimdark/internal/game/actions"

	"github.com/google/uuid"
)

var Spear = game.Weapon{
	Item: game.Item{
		ID:          uuid.MustParse("019faff0-7f1d-7755-86d2-b24cae061810"),
		Name:        "Spear",
		Description: "Standard spear used commonly by soldiers. Long reach, and often used with shields.",
		Effects:     []game.Effect{},
	},
	Actions: []game.Action{
		actions.BackStrike,
		actions.Cleave,
		actions.Poke,
		actions.Latch,
		actions.Sharpen,
	},
	OffsetStats: map[game.Stat]float64{
		game.Melee:  16,
		game.Health: 16,
	},
	Weight:     1,
	WeaponType: game.WeaponTypeSpear,
}
