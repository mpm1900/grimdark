package game

import (
	"github.com/google/uuid"
)

type PlayerPosition struct {
	ID       uuid.UUID `json:"ID"`
	ActorID  uuid.UUID `json:"actor_ID"`
	PlayerID uuid.UUID `json:"player_ID"`
	Rank     int       `json:"rank"`
}

type Player struct {
	ID        uuid.UUID        `json:"ID"`
	User      User             `json:"user"`
	Positions []PlayerPosition `json:"positions"`
}

func NewPlayer() Player {
	player_ID := uuid.New()
	return Player{
		ID: player_ID,
		Positions: []PlayerPosition{
			{
				ID:       uuid.New(),
				ActorID:  uuid.Nil,
				PlayerID: player_ID,
				Rank:     0,
			},
			{
				ID:       uuid.New(),
				ActorID:  uuid.Nil,
				PlayerID: player_ID,
				Rank:     1,
			},
			{
				ID:       uuid.New(),
				ActorID:  uuid.Nil,
				PlayerID: player_ID,
				Rank:     2,
			},
		},
	}
}

func (p PlayerPosition) GetDistanceFrom(other PlayerPosition) int {
	a := p.Rank
	b := other.Rank
	if p.PlayerID != other.PlayerID {
		return a + b + 1
	}
	return Abs(a - b)
}

func (p Player) GetOpenPositions() []uuid.UUID {
	positions := []uuid.UUID{}
	for _, pos := range p.Positions {
		if pos.ActorID == uuid.Nil {
			positions = append(positions, pos.ID)
		}
	}

	return positions
}
func (p Player) GetPosition(position_id uuid.UUID) (PlayerPosition, bool) {
	for _, pos := range p.Positions {
		if pos.ID == position_id {
			return pos, true
		}
	}

	return PlayerPosition{}, false
}
func (p Player) GetPositionByActorID(actor_id uuid.UUID) (PlayerPosition, bool) {
	for _, pos := range p.Positions {
		if pos.ActorID == actor_id {
			return pos, true
		}
	}

	return PlayerPosition{}, false
}

func (p *Player) UpdatePositions(where func(PlayerPosition) bool, updater func(PlayerPosition) PlayerPosition) {
	for i, pos := range p.Positions {
		if where(pos) {
			p.Positions[i] = updater(pos)
		}
	}
}
func (p *Player) UpdatePosition(position_id uuid.UUID, updater func(PlayerPosition) PlayerPosition) {
	p.UpdatePositions(func(pp PlayerPosition) bool {
		return pp.ID == position_id
	}, updater)
}
func (p *Player) UpdatePositionActor(position_id uuid.UUID, actor_id uuid.UUID) {
	p.UpdatePosition(position_id, func(pp PlayerPosition) PlayerPosition {
		pp.ActorID = actor_id
		return pp
	})
}
func (p *Player) SetPosition(position_id uuid.UUID, actor Actor) (uuid.UUID, bool) {
	var evicted_id uuid.UUID
	updated := false

	for _, pos := range p.Positions {
		if pos.ActorID == actor.ID && pos.ID != position_id {
			p.UpdatePositionActor(pos.ID, uuid.Nil)
			updated = true
		}
	}

	if position_id == uuid.Nil {
		return evicted_id, updated || actor.IsActive()
	}

	current, ok := p.GetPosition(position_id)
	if !ok {
		return evicted_id, false
	}

	if current.ActorID != uuid.Nil && current.ActorID != actor.ID {
		evicted_id = current.ActorID
	}

	p.UpdatePositionActor(position_id, actor.ID)
	return evicted_id, true
}
