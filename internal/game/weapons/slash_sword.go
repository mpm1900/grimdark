package weapons

import (
	"grimdark/internal/game"
	"grimdark/internal/game/actions"

	"github.com/google/uuid"
)

var SlashSword = game.Weapon{
	ID:      uuid.MustParse("019f0b34-e7a4-78e5-9e60-e4a50ded7a11"),
	Name:    "Slash Sword",
	Actions: []game.Action{actions.Slash, actions.Blast},
	AuxStats: map[game.Stat]float64{
		game.Melee: 20,
	},
	Description: "A large sword, designed for slashing.",
	Effects:     []game.Effect{},
	Hands:       2,
	WeaponType:  "Sword",
}
