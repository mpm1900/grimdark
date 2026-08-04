package weapons

import "grimdark/internal/game"

var TomeOfSacrifice = game.Weapon{
	Item:        game.Item{},
	Actions:     []game.Action{},
	OffsetStats: map[game.Stat]float64{},
	Weight:      2,
	WeaponType:  game.WeaponTypeTome,
}
