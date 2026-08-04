package weapons

import (
	"grimdark/internal/game"
	"grimdark/internal/game/actions"
	"grimdark/internal/game/effects"

	"github.com/google/uuid"
)

func LongBow() game.Weapon {
	weapon := game.Weapon{
		Item: game.Item{
			ID:          uuid.MustParse("019f859c-d25e-724c-aceb-bf82cf5c7389"),
			Name:        "Long Bow",
			Description: "A high-precision, long-range bow, widely used to pick off high value targets. Expensive and rare, they are given to the best marksmen and sharpshooters.",
			Effects: []game.Effect{
				effects.OtherEye,
			},
		},
		Actions: []game.Action{
			actions.CalledShot,
			actions.CollateralShot,
			actions.Headshot,
			actions.LockOn,
			actions.PiercingShot,
			actions.RetreatingShot,
		},
		OffsetStats: map[game.Stat]float64{
			game.Ranged:         32,
			game.Special:        16,
			game.SpecialDefense: 16,
		},
		Weight:     2,
		WeaponType: game.WeaponTypeLongBow,
	}

	return weapon
}
