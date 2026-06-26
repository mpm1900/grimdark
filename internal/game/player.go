package game

import (
	"github.com/google/uuid"
)

type Player struct {
	ID        uuid.UUID               `json:"ID"`
	User      User                    `json:"user"`
	Positions map[uuid.UUID]uuid.UUID `json:"positions"`
}

func NewPlayer() Player {
	return Player{
		ID: uuid.New(),
		Positions: map[uuid.UUID]uuid.UUID{
			uuid.New(): uuid.Nil,
			uuid.New(): uuid.Nil,
		},
	}
}

func (p Player) GetOpenPosition() uuid.UUID {
	for pos, actor_id := range p.Positions {
		if actor_id == uuid.Nil {
			return pos
		}
	}

	return uuid.Nil
}
