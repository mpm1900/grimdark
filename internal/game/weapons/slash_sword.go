package weapons

import (
	"grimdark/internal/game"
	"grimdark/internal/game/actions"

	"github.com/google/uuid"
)

var SlashSword = game.Weapon{
	Item: game.Item{
		ID:          uuid.MustParse("019f0b34-e7a4-78e5-9e60-e4a50ded7a11"),
		Name:        "Slash Sword",
		Description: "A large sword, designed for slashing.",
		Effects:     []game.Effect{},
	},
	Actions: []game.Action{actions.Slash, actions.Blast, actions.SwordsDance},
	OffsetStats: map[game.Stat]float64{
		game.Melee: 20,
	},
	Weight:     1,
	WeaponType: game.WeaponTypeSword,
}
