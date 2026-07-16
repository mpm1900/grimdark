package game

import (
	"cmp"
	"slices"

	"github.com/google/uuid"
)

func (g *Game) SetPosition(actor_id uuid.UUID, position_id uuid.UUID) {
	actor, ok := g.GetActor(actor_id)
	if !ok {
		return
	}

	var evicted_id uuid.UUID
	g.mutate(func(s *State) {
		updated := false
		evicted_id, updated = s.SetPosition(position_id, actor)

		if !updated {
			return
		}

		s.UpdateActor(actor_id, func(a Actor) Actor {
			a.SetPosition(position_id)
			return a
		})

		trigger_context := MakeContextFor(actor)

		//leave
		if position_id == uuid.Nil {
			s.Commands = slices.DeleteFunc(s.Commands, func(cmd Command) bool {
				return cmd.Context.ParentID == actor.ID
			})
			s.Modifiers = slices.DeleteFunc(s.Modifiers, func(mod Modifier) bool {
				return mod.Context.ParentID == actor.ID
			})

			if actor.IsAlive {
				log := NewLog("$source$ left the battle.", SourceTerms(actor))
				g.PushLogMeta(log.Bind(trigger_context))
			} else {
				log := NewLog("$source$ died.", SourceTerms(actor))
				g.PushLogMeta(log.Bind(trigger_context))
			}
			g.On(OnActorLeave, trigger_context)
		}

		//join
		if actor.PositionID == uuid.Nil {
			log := NewLog("$source$ joined the battle.", SourceTerms(actor))
			g.PushLogMeta(log.Bind(trigger_context))
			g.On(OnActorEnter, trigger_context)
		}

		//move
		if actor.PositionID != uuid.Nil && position_id != uuid.Nil {
			g.On(OnActorMove, trigger_context)
		}
	})

	if evicted_id != uuid.Nil {
		g.SetPosition(evicted_id, actor.PositionID)
	}
}
func (g *Game) PushForwards(actor_id uuid.UUID) {
	g.moveActor(actor_id, -1)
}
func (g *Game) PushToFront(actor_id uuid.UUID) {
	for g.moveActor(actor_id, -1) {
	}
}
func (g *Game) PushBackwards(actor_id uuid.UUID) {
	g.moveActor(actor_id, 1)
}
func (g *Game) PushToBack(actor_id uuid.UUID) {
	for g.moveActor(actor_id, 1) {
	}
}
func (g *Game) moveActor(actor_id uuid.UUID, direction int) bool {
	actor, ok := g.GetActor(actor_id)
	if !ok {
		return false
	}
	position, ok := g.state.GetPositionByActorID(actor_id)
	if !ok {
		return false
	}

	next, ok := g.NextAllyPositionByRank(actor.PlayerID, position.Rank, direction)
	if !ok {
		return false
	}

	g.SetPosition(actor.ID, next.ID)
	log_ctx := MakeContextFrom(actor)
	log_ctx.PositionIDs = []uuid.UUID{next.ID}
	if direction > 0 {
		log := NewLog("$source$ moved backwards.", SourceTerms(actor))
		g.PushLogMeta(log.Bind(log_ctx))
	}
	if direction < 0 {
		log := NewLog("$source$ moved forwards.", SourceTerms(actor))
		g.PushLogMeta(log.Bind(log_ctx))
	}

	return true
}

func (g *Game) condensePositions() bool {
	moved := false
	for {
		actorID, ok := g.nextActorBehindGap()
		if !ok {
			return moved
		}

		if !g.moveActor(actorID, -1) {
			return moved
		}
		moved = true
	}
}

func (g *Game) nextActorBehindGap() (uuid.UUID, bool) {
	state := g.State()
	for _, player := range state.Players {
		positions := state.GetPositionsByPlayerID(player.ID)
		slices.SortStableFunc(positions, func(a, b Position) int {
			return cmp.Compare(a.Rank, b.Rank)
		})

		foundGap := false
		for _, position := range positions {
			if position.ActorID == uuid.Nil {
				foundGap = true
				continue
			}

			if foundGap {
				return position.ActorID, true
			}
		}
	}

	return uuid.Nil, false
}
