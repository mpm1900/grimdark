package game

import (
	"encoding/json"
	"slices"

	"github.com/google/uuid"
)

type Context struct {
	ActionID uuid.UUID
	EffectID uuid.UUID
	PlayerID uuid.UUID

	ParentID    uuid.UUID
	SourceID    uuid.UUID
	ActorIDs    []uuid.UUID
	PositionIDs []uuid.UUID
}

type contextJSON struct {
	ActionID    *uuid.UUID  `json:"action_ID"`
	EffectID    *uuid.UUID  `json:"effect_ID"`
	PlayerID    *uuid.UUID  `json:"player_ID"`
	ParentID    *uuid.UUID  `json:"parent_ID"`
	SourceID    *uuid.UUID  `json:"source_ID"`
	ActorIDs    []uuid.UUID `json:"actor_IDs"`
	PositionIDs []uuid.UUID `json:"position_IDs"`
}

func (c Context) MarshalJSON() ([]byte, error) {
	return json.Marshal(contextJSON{
		ActionID:    NilifyUUID(c.ActionID),
		EffectID:    NilifyUUID(c.EffectID),
		PlayerID:    NilifyUUID(c.PlayerID),
		ParentID:    NilifyUUID(c.ParentID),
		SourceID:    NilifyUUID(c.SourceID),
		ActorIDs:    (c.ActorIDs),
		PositionIDs: (c.PositionIDs),
	})
}

func (c *Context) UnmarshalJSON(data []byte) error {
	var payload contextJSON
	if err := json.Unmarshal(data, &payload); err != nil {
		return err
	}

	*c = NewContext()
	if payload.ActionID != nil {
		c.ActionID = *payload.ActionID
	}
	if payload.EffectID != nil {
		c.EffectID = *payload.EffectID
	}
	if payload.PlayerID != nil {
		c.PlayerID = *payload.PlayerID
	}
	if payload.ParentID != nil {
		c.ParentID = *payload.ParentID
	}
	if payload.SourceID != nil {
		c.SourceID = *payload.SourceID
	}
	if payload.ActorIDs != nil {
		c.ActorIDs = payload.ActorIDs
	}
	if payload.PositionIDs != nil {
		c.PositionIDs = payload.PositionIDs
	}

	return nil
}

func NewContext() Context {
	return Context{
		ActionID:    uuid.Nil,
		EffectID:    uuid.Nil,
		PlayerID:    uuid.Nil,
		ParentID:    uuid.Nil,
		SourceID:    uuid.Nil,
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

func MakeContextPlayer(player_id uuid.UUID) Context {
	ctx := NewContext()
	ctx.PlayerID = player_id
	return ctx
}
func MakeContextFrom(actor Actor) Context {
	ctx := NewContext()
	ctx.SourceID = actor.ID
	ctx.ParentID = actor.ID
	ctx.PlayerID = actor.PlayerID
	return ctx
}
func MakeContextFor(source Actor, targets ...Actor) Context {
	ctx := MakeContextFrom(source)
	for _, target := range targets {
		ctx.AddTarget(target)
	}
	return ctx
}
func MakeModifierContext(source Actor, target Actor) Context {
	ctx := MakeContextFrom(source)
	ctx.ParentID = target.ID
	ctx.AddModifierTarget(target)
	return ctx
}

func (c *Context) AddTarget(target Actor) {
	if target.IsActive() {
		c.PositionIDs = append(c.PositionIDs, target.PositionID)
	} else {
		c.ActorIDs = append(c.ActorIDs, target.ID)
	}
}
func (c *Context) AddModifierTarget(target Actor) {
	c.ActorIDs = append(c.ActorIDs, target.ID)
}
func (c *Context) RemoveTarget(target Actor) {
	actor_ids := []uuid.UUID{}
	pos_ids := []uuid.UUID{}

	for _, aid := range c.ActorIDs {
		if aid != target.ID {
			actor_ids = append(actor_ids, aid)
		}
	}
	c.ActorIDs = actor_ids

	if target.IsActive() {
		for _, pid := range c.PositionIDs {
			if pid != target.PositionID {
				pos_ids = append(pos_ids, pid)
			}
		}
		c.PositionIDs = pos_ids
	}
}
func (c *Context) ClearTargets() {
	c.ActorIDs = []uuid.UUID{}
	c.PositionIDs = []uuid.UUID{}
}
func (c Context) CloneWithTarget(target Actor) Context {
	clone := c.Clone()
	clone.ActorIDs = []uuid.UUID{}
	clone.PositionIDs = []uuid.UUID{}
	clone.AddTarget(target)

	return clone
}

func (c Context) CloneWithTargets(targets []Actor) Context {
	clone := c.Clone()
	clone.ActorIDs = []uuid.UUID{}
	clone.PositionIDs = []uuid.UUID{}
	for _, target := range targets {
		clone.AddTarget(target)
	}

	return clone
}
func (c Context) HasTarget(target Actor) bool {
	if target.IsActive() && slices.Contains(c.PositionIDs, target.PositionID) {
		return true
	}

	return slices.Contains(c.ActorIDs, target.ID)
}
func (c Context) GetTargetCount() int {
	return len(c.ActorIDs) + len(c.PositionIDs)
}

func (g *Game) GetSource(context Context) (Actor, bool) {
	return g.GetActor(context.SourceID)
}
func (g *Game) GetSourceAction(context Context) Actor {
	source, _ := g.GetSource(context)
	return source
}
func (g *Game) GetParent(context Context) (Actor, bool) {
	return g.GetActor(context.ParentID)
}
func (g *Game) GetTargets(context Context) []Actor {
	actors := []Actor{}
	for _, a := range g.State().Actors {
		if slices.Contains(context.ActorIDs, a.ID) {
			actors = append(actors, a)
		}
		if a.IsActive() && slices.Contains(context.PositionIDs, a.PositionID) {
			actors = append(actors, a)
		}
	}

	return actors
}
func (g *Game) FindAction(context Context) (Action, bool) {
	actor, actor_ok := g.GetSource(context)
	if actor_ok {
		action, action_ok := actor.GetActionByID(context.ActionID)
		if action_ok {
			return action, true
		}
	}

	for _, global_action := range GLOBAL_ACTIONS {
		if global_action.ID == context.ActionID {
			return global_action, true
		}
	}

	return Action{}, false
}
func (g *Game) HydrateToCommand(context Context) (Command, bool) {
	action, ok := g.FindAction(context)
	if !ok {
		return Command{}, false
	}

	return action.Bind(context), true
}
