package game

import (
	"github.com/google/uuid"
)

type Player struct {
	ID   uuid.UUID `json:"ID"`
	User User      `json:"user"`
}

func NewPlayer(id uuid.UUID) Player {
	return Player{
		ID: id,
	}
}
