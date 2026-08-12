package items

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

var Global = []game.Item{
	CorruptedNecklace(),
	CursedBoots(),
	CursedRing(),
	Rations(),
}

func HydrateGlobalItem(id uuid.UUID) (game.Item, bool) {
	for _, item := range Global {
		if item.ID == id {
			return item, true
		}
	}

	return game.Item{}, false
}
