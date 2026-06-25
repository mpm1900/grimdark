package game

import "github.com/google/uuid"

type Player struct {
	ID        uuid.UUID
	Positions map[uuid.UUID]struct{}
}

func NewPlayer() Player {
	return Player{
		ID:        uuid.New(),
		Positions: map[uuid.UUID]struct{}{},
	}
}
