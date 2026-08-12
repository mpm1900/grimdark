package weapons

import (
	"grimdark/internal/game"
	"grimdark/internal/game/actions"

	"github.com/google/uuid"
)

func HandCrossbow() game.Weapon {
	id := uuid.MustParse("019f5319-dffb-7d06-b6f3-af8dca62bffd")

	return game.Weapon{
		Item: game.Item{
			ID: id,
			Entity: game.MakeEntity(
				id,
				"Hand Crossbow",
				"A compact, one-handed mechanical weapon commonly utilized by rogues, spies, and scoundrels for silent or sudden strikes.",
			),
			Effects: []game.Effect{},
		},
		Actions: []game.Action{
			actions.BurstFire(),
			actions.FiftyFifty(),
			actions.CalledShot(),
		},
		OffsetStats: map[game.Stat]float64{
			game.Ranged: 20,
		},
		Weight:     1,
		WeaponType: game.WeaponTypeCrossBow,
	}
}
