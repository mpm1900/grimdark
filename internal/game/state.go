package game

import (
	"slices"

	"github.com/google/uuid"
)

type State struct {
	// turn ...
	Players       []Player
	Actors        []Actor
	Transactions  Queue[Transaction]
	Modifiers     []Modifier
	Commands      Queue[Command]
	Triggers      Queue[Command]
	ActiveContext *Context
}

func (s *State) Clone() State {
	actors := make([]Actor, len(s.Actors))
	for i, actor := range s.Actors {
		actors[i] = actor.Clone()
	}

	transactions := slices.Clone(s.Transactions)
	for i := range transactions {
		transactions[i].Context = transactions[i].Context.Clone()
	}

	modifiers := slices.Clone(s.Modifiers)
	for i := range modifiers {
		modifiers[i].Context = modifiers[i].Context.Clone()
	}

	commands := slices.Clone(s.Commands)
	for i := range commands {
		commands[i].Context = commands[i].Context.Clone()
	}

	triggers := slices.Clone(s.Triggers)
	for i := range triggers {
		triggers[i].Context = triggers[i].Context.Clone()
	}

	var activeContext *Context
	if s.ActiveContext != nil {
		cloned := s.ActiveContext.Clone()
		activeContext = &cloned
	}

	return State{
		Players:       slices.Clone(s.Players),
		Actors:        actors,
		Transactions:  transactions,
		Modifiers:     modifiers,
		Commands:      commands,
		Triggers:      triggers,
		ActiveContext: activeContext,
	}
}

func (s State) FindActorByID(id uuid.UUID) (Actor, bool) {
	for _, a := range s.Actors {
		if a.ID == id {
			return a, true
		}
	}

	return Actor{}, false
}
func (s State) FindActorsWhere(where Filter[Actor], context Context) []Actor {
	actors := []Actor{}
	for _, a := range s.Actors {
		if where(a, context) {
			actors = append(actors, a)
		}
	}

	return actors
}

func (s *State) UpdateActor(id uuid.UUID, updater func(Actor) Actor) {
	for i, a := range s.Actors {
		if a.ID == id {
			s.Actors[i] = updater(a)
		}
	}
}
func (s *State) UpdateActorWhere(where func(Actor) bool, updater func(Actor) Actor) {
	for i, a := range s.Actors {
		if where(a) {
			s.Actors[i] = updater(a)
		}
	}
}
