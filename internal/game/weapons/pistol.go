package weapons

import (
	"grimdark/internal/game"
	"grimdark/internal/game/actions"
	"grimdark/internal/game/effects"

	"github.com/google/uuid"
)

var Pistol = game.Weapon{
	Item: game.Item{
		ID:          uuid.MustParse("019f5319-dffb-7d06-b6f3-af8dca62bffd"),
		Name:        "Pistol",
		Description: "A bolter pistol. Pew pew.",
		Effects:     []game.Effect{effects.AuraOfWeakness},
	},
	Actions: []game.Action{actions.Slash, actions.Blast, actions.SwordsDance, actions.EdictOfSpeed},
	OffsetStats: map[game.Stat]float64{
		game.Ranged: 20,
	},
	Weight:     1,
	WeaponType: game.WeaponTypePistol,
}
