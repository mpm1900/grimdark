package game

import (
	"github.com/google/uuid"
)

type Player struct {
	ID   uuid.UUID
	User User
}

type playerJSON struct {
	ID         uuid.UUID `json:"ID"`
	User       User      `json:"user"`
	ActorCount int       `json:"actor_count"`
}

func NewPlayer(user User) Player {
	return Player{
		ID:   user.ID,
		User: user,
	}
}

func (p Player) ToJSON(actor_count int) playerJSON {
	return playerJSON{
		ID:         p.ID,
		User:       p.User,
		ActorCount: actor_count,
	}
}
