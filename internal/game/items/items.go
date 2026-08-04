package items

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

var Global = map[uuid.UUID]game.Item{
	CursedRing.ID: CursedRing,
}

func HydrateGlobalItem(id uuid.UUID) (game.Item, bool) {
	item, ok := Global[id]
	return item, ok
}
