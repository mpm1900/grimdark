package game

import (
	"slices"

	"github.com/google/uuid"
)

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

func (g *Game) GetPosition(position_id uuid.UUID) (Position, bool) {
	return g.State().GetPosition(position_id)
}
func (g *Game) GetDistance(a uuid.UUID, b uuid.UUID) (int, bool) {
	position_a, aok := g.GetPosition(a)
	position_b, bok := g.GetPosition(b)
	if !aok || !bok {
		return 0, false
	}

	return position_a.GetDistanceFrom(position_b), true
}

func (g *Game) NextAllyPositionByRank(player_ID uuid.UUID, rank int, direction int) (Position, bool) {
	var next Position
	found := false

	for _, position := range g.State().Positions {
		if position.PlayerID != player_ID {
			continue
		}

		if direction < 0 && position.Rank < rank && (!found || position.Rank > next.Rank) {
			next = position
			found = true
		}
		if direction > 0 && position.Rank > rank && (!found || position.Rank < next.Rank) {
			next = position
			found = true
		}
	}

	return next, found
}
func (g *Game) GetEnemyPositionsByRank(player_ID uuid.UUID, rank int, direction int) []Position {
	positions := []Position{}

	// sort by rank so that if an early actor needs to stop the collateral, the front-most ones stop first
	state_positions := g.State().Positions
	slices.SortStableFunc(state_positions, func(a, b Position) int {
		return a.Rank - b.Rank
	})
	for _, position := range state_positions {
		if position.PlayerID == player_ID {
			continue
		}

		if direction < 0 && position.Rank <= rank {
			positions = append(positions, position)
		}
		if direction > 0 && position.Rank >= rank {
			positions = append(positions, position)
		}

		target, ok := g.GetActor(position.ActorID)
		if ok {
			if target.IsBulwark {
				break
			}
		}
	}

	return positions
}
