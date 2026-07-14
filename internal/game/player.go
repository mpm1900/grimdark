package game

import (
	"github.com/google/uuid"
)

type Player struct {
	ID    uuid.UUID
	Ready bool
	User  User
}

type playerJSON struct {
	ID         uuid.UUID `json:"ID"`
	Ready      bool      `json:"ready"`
	User       User      `json:"user"`
	ActorCount int       `json:"actor_count"`
}

func NewPlayer(user User) *Player {
	return &Player{
		ID:    user.ID,
		Ready: false,
		User:  user,
	}
}

func (p *Player) SetReady() {
	p.Ready = true
}

func (p *Player) NextTurn() {
	p.Ready = false
}

func (p *Player) ToJSON(actor_count int) playerJSON {
	return playerJSON{
		ID:         p.ID,
		ActorCount: actor_count,
		Ready:      p.Ready,
		User:       p.User,
	}
}
