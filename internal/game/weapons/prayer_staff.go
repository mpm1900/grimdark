package weapons

import (
	"grimdark/internal/game"
	"grimdark/internal/game/actions"

	"github.com/google/uuid"
)

var PrayerStaff = game.Weapon{
	Item: game.Item{
		ID:          uuid.MustParse("019fabb9-0f83-75cb-b0d2-50531a77e321"),
		Name:        "Prayer Staff",
		Description: "A prayer staff of the Greater Will, used as a means to better commune with the divne.",
		Effects:     []game.Effect{},
	},
	Actions: []game.Action{
		actions.Protect,
	},
	OffsetStats: map[game.Stat]float64{
		game.Health:         32,
		game.SpecialDefense: 32,
	},
	Weight:     2,
	WeaponType: game.WeaponTypeTome,
}
