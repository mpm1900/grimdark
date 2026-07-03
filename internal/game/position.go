package game

import "github.com/google/uuid"

type Position struct {
	ID       uuid.UUID `json:"ID"`
	ActorID  uuid.UUID `json:"actor_ID"`
	PlayerID uuid.UUID `json:"player_ID"`
	Rank     int       `json:"rank"`
}

func (p Position) GetDistanceFrom(other Position) int {
	a := p.Rank
	b := other.Rank
	if p.PlayerID != other.PlayerID {
		return a + b + 1
	}
	return Abs(a - b)
}
