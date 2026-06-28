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

func (p Player) GetOpenPositions() []uuid.UUID {
	positions := []uuid.UUID{}
	for pos, actor_id := range p.Positions {
		if actor_id == uuid.Nil {
			positions = append(positions, pos)
		}
	}

	return positions
}

func (p *Player) SetPosition(position_id uuid.UUID, actor Actor) (uuid.UUID, bool) {
	var evicted_id uuid.UUID
	updated := false

	for pos, actor_id := range p.Positions {
		if actor_id == actor.ID && pos != position_id {
			p.Positions[pos] = uuid.Nil
			updated = true
		}
	}

	if position_id == uuid.Nil {
		return evicted_id, updated || actor.IsActive()
	}

	current_id, ok := p.Positions[position_id]
	if !ok {
		return evicted_id, false
	}

	if current_id != uuid.Nil && current_id != actor.ID {
		evicted_id = current_id
	}

	p.Positions[position_id] = actor.ID
	return evicted_id, true
}
