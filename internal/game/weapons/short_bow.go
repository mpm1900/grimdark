package weapons

import (
	"grimdark/internal/game"
	"grimdark/internal/game/actions"

	"github.com/google/uuid"
)

func ShortBow() game.Weapon {
	weapon := game.Weapon{
		Item: game.Item{
			ID:          uuid.MustParse("019faf80-b98e-7500-8be1-80949bde79f1"),
			Name:        "Short Bow",
			Description: "A mid-range bow, versitile and used to deal damage over multiple hits and multiple targets.",
			Effects:     []game.Effect{},
		},
		Actions: []game.Action{
			actions.DoubleShot,
			actions.FiftyFifty,
			actions.LockOn,
			actions.PinDown,
			actions.Protect,
			actions.SpreadShot,
		},
		OffsetStats: map[game.Stat]float64{
			game.Ranged:         32,
			game.Special:        16,
			game.SpecialDefense: 16,
		},
		Weight:     2,
		WeaponType: game.WeaponTypeShortBow,
	}

	return weapon
}
