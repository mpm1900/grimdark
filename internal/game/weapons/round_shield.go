package weapons

import (
	"grimdark/internal/game"
	"grimdark/internal/game/actions"

	"github.com/google/uuid"
)

var RoundShield = game.Weapon{
	Item: game.Item{
		ID:          uuid.MustParse("019fb102-abae-7174-9c46-6aa0fdabdf9a"),
		Name:        "Round Shield",
		Description: "Small, leather-covered round shield, reinforced in critical spots with metal.",
		Effects:     []game.Effect{},
	},
	Actions: []game.Action{
		actions.Protect,
	},
	OffsetStats: map[game.Stat]float64{
		game.Health:         16,
		game.MartialDefense: 16,
	},
	Weight:     1,
	WeaponType: game.WeaponTypeShield,
}
