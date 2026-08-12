package weapons

import (
	"grimdark/internal/game"
	"grimdark/internal/game/actions"

	"github.com/google/uuid"
)

func TomeOfSacrifice() game.Weapon {
	id := uuid.MustParse("019fe756-d14d-749e-8625-7231f53e168d")

	return game.Weapon{
		Item: game.Item{
			ID: id,
			Entity: game.MakeEntity(
				id,
				"Tome of Sacrifice",
				"",
			),
			Effects: []game.Effect{},
		},
		Actions: []game.Action{
			actions.ArcaneRitual(),
			actions.Warp(),
		},
		OffsetStats: map[game.Stat]float64{},
		Weight:      2,
		WeaponType:  game.WeaponTypeTome,
	}
}
