package weapons

import (
	"grimdark/internal/game"
	"grimdark/internal/game/actions"

	"github.com/google/uuid"
)

var Greatsword = game.Weapon{
	Item: game.Item{
		ID:          uuid.MustParse("019f4a69-3324-70f6-80e7-aded9c2c1f13"),
		Name:        "Greatsword",
		Description: "A large sword, used when bullets fail to stop quick or well-armoured targets. The strikes from these weapons can easily lop off limbs and heads.",
		Effects:     []game.Effect{},
	},
	Actions: []game.Action{
		actions.Charge,
		actions.Cleave,
		actions.Execute,
		actions.HeavySwing,
	},
	OffsetStats: map[game.Stat]float64{
		game.Melee:          32,
		game.MartialDefense: 32,
	},
	Weight:     2,
	WeaponType: game.WeaponTypeBigSword,
}
