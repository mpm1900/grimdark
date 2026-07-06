package game

import "github.com/google/uuid"

type Item struct {
	ID          uuid.UUID `json:"ID"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type HeldItem struct {
	Item
	Effects []Effect `json:"effects"`
}
