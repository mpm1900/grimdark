package weapons

import (
	"grimdark/internal/game"
	"grimdark/internal/game/actions"
)

var TomeOfSacrifice = game.Weapon{
	Item: game.Item{
		Name:    "Tome of Sacrifice",
		Effects: []game.Effect{},
	},
	Actions: []game.Action{
		actions.ArcaneRitual,
		actions.Warp,
	},
	OffsetStats: map[game.Stat]float64{},
	Weight:      2,
	WeaponType:  game.WeaponTypeTome,
}
