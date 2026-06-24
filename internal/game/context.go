package game

import (
	"slices"

	"github.com/google/uuid"
)

type Context struct {
	ActionID *uuid.UUID
	EffectID *uuid.UUID
	PlayerID *uuid.UUID

	ParentID    *uuid.UUID
	SourceID    *uuid.UUID
	ActorIDs    []uuid.UUID
	PositionIDs []uuid.UUID
}

func NewContext() Context {
	return Context{
		ActionID: nil,
		EffectID: nil,
		PlayerID: nil,

		ParentID:    nil,
		SourceID:    nil,
		ActorIDs:    []uuid.UUID{},
		PositionIDs: []uuid.UUID{},
	}
}

func (c Context) Clone() Context {
	return Context{
		ActionID:    c.ActionID,
		EffectID:    c.EffectID,
		PlayerID:    c.PlayerID,
		ParentID:    c.ParentID,
		SourceID:    c.SourceID,
		ActorIDs:    slices.Clone(c.ActorIDs),
		PositionIDs: slices.Clone(c.PositionIDs),
	}
}

func MakeContextFrom(actor Actor) Context {
	ctx := NewContext()
	ctx.SourceID = &actor.ID
	ctx.ParentID = &actor.ID
	ctx.PlayerID = &actor.PlayerID
	return ctx
}
func MakeContextFor(source, target Actor) Context {
	ctx := MakeContextFrom(source)
	ctx.AddTarget(target)
	return ctx
}

func (c *Context) AddTarget(target Actor) {
	if target.PositionID != nil {
		c.PositionIDs = append(c.PositionIDs, *target.PositionID)
	} else {
		c.ActorIDs = append(c.ActorIDs, target.ID)
	}
}
func (c Context) CloneWithTarget(target Actor) Context {
	c = c.Clone()
	c.ActorIDs = []uuid.UUID{}
	c.PositionIDs = []uuid.UUID{}
	if target.PositionID != nil {
		c.PositionIDs = append(c.PositionIDs, *target.PositionID)
	} else {
		c.ActorIDs = append(c.ActorIDs, target.ID)
	}

	return c
}

func (g *Game) GetSource(context Context) (Actor, bool) {
	if context.SourceID == nil {
		return Actor{}, false
	}

	return g.GetActor(*context.SourceID)
}
func (g *Game) GetSourceAction(context Context) Actor {
	source, _ := g.GetSource(context)
	return source
}
func (g *Game) GetParent(context Context) (Actor, bool) {
	if context.ParentID == nil {
		return Actor{}, false
	}

	return g.GetActor(*context.ParentID)
}
func (g *Game) GetTargets(context Context) []Actor {
	actors := []Actor{}
	for _, a := range g.State().Actors {
		if slices.Contains(context.ActorIDs, a.ID) {
			actors = append(actors, a)
		}
		if a.PositionID != nil && slices.Contains(context.PositionIDs, *a.PositionID) {
			actors = append(actors, a)
		}
	}

	return actors
}
