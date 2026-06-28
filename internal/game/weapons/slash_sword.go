package weapons

import (
	"grimdark/internal/game"
	"grimdark/internal/game/actions"

	"github.com/google/uuid"
)

var SlashSword = game.Weapon{
	ID:      uuid.MustParse("019f0b34-e7a4-78e5-9e60-e4a50ded7a11"),
	Name:    "Slash Sword",
	Actions: []game.Action{actions.Slash},
	AuxStats: map[game.Stat]float64{
		game.Melee: 20,
	},
	Effects: []game.Effect{},
}
