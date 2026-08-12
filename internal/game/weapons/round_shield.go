package weapons

import (
	"grimdark/internal/game"
	"grimdark/internal/game/actions"

	"github.com/google/uuid"
)

func RoundShield() game.Weapon {
	id := uuid.MustParse("019fb102-abae-7174-9c46-6aa0fdabdf9a")

	return game.Weapon{
		Item: game.Item{
			ID: id,
			Entity: game.MakeEntity(
				id,
				"Round Shield",
				"Small, leather-covered round shield, reinforced in critical spots with metal.",
			),
			Effects: []game.Effect{},
		},
		Actions: []game.Action{
			actions.ArmorUp(),
			actions.Protect(),
		},
		OffsetStats: map[game.Stat]float64{
			game.Health:         16,
			game.MartialDefense: 16,
		},
		Weight:     1,
		WeaponType: game.WeaponTypeShield,
	}
}
