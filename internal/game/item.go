package game

import (
	"slices"

	"github.com/google/uuid"
)

type Item struct {
	ID          uuid.UUID `json:"ID"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Effects     []Effect  `json:"effects"`
}

func (i Item) Clone() Item {
	return Item{
		ID:          i.ID,
		Name:        i.Name,
		Description: i.Description,
		Effects:     slices.Clone(i.Effects),
	}
}
