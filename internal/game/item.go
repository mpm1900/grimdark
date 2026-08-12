package game

import (
	"slices"

	"github.com/google/uuid"
)

type Item struct {
	ID      uuid.UUID `json:"ID"`
	Entity  Entity    `json:"entity"`
	Effects []Effect  `json:"effects"`
}

func (i Item) Clone() Item {
	return Item{
		ID:      i.ID,
		Effects: slices.Clone(i.Effects),
	}
}
