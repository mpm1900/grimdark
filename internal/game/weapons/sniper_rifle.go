package weapons

import (
	"grimdark/internal/game"
	"grimdark/internal/game/actions"

	"github.com/google/uuid"
)

var RelicRifle = game.Weapon{
	Item: game.Item{
		ID:   uuid.MustParse("019f859c-d25e-724c-aceb-bf82cf5c7389"),
		Name: "Sniper Rifle",
		// Description: "A long rifle that is as much a holy relic as it is a weapon, a divinely-inspired instrument of death crafted by the most sought-after gunsmiths in the world and blessed by saintly figures.",
		Description: "A high-precision, long-range rifle, widely used to pick off high value targets such as officers, sappers and artillery crews. Expensive and rare, they are given to the best marksmen and sharpshooters of the warband.",
		Effects:     []game.Effect{},
	},
	Actions: []game.Action{
		actions.CalledShot,
		actions.CollateralShot,
		actions.Headshot,
		actions.LockOn,
		actions.PiercingShot,
		actions.Protect,
	},
	OffsetStats: map[game.Stat]float64{
		game.Ranged:         32,
		game.Special:        16,
		game.SpecialDefense: 16,
	},
	Weight:     2,
	WeaponType: game.WeaponTypeRifle,
}
