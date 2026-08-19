package weapons

import (
	"grimdark/internal/game"
	"grimdark/internal/game/actions"
	"grimdark/internal/game/effects"

	"github.com/google/uuid"
)

func PrayerStaff() game.Weapon {
	id := uuid.MustParse("019fabb9-0f83-75cb-b0d2-50531a77e321")

	return game.Weapon{
		Item: game.Item{
			ID: id,
			Entity: game.MakeEntity(
				id,
				"Prayer Staff",
				"A prayer staff of the Greater Will, used as a means to better commune with the divne.",
			),
			Effects: []game.Effect{
				effects.Devout(),
			},
		},
		Actions: []game.Action{
			actions.BlessingOfMichael(),
			actions.CommandRepeat(),
			actions.HealingBlessing(),
			actions.HealingPrayer(),
			actions.HealStatus(),
		},
		OffsetStats: map[game.Stat]float64{
			game.SpecialDefense: 32,
		},
		Weight:     1,
		WeaponType: game.WeaponTypeStaff,
	}
}
