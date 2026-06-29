package game

import (
	"slices"

	"github.com/google/uuid"
)

type State struct {
	ActiveContext *Context
	Actors        []Actor
	Commands      Queue[Command]
	Modifiers     []Modifier
	Players       []Player
	Prompts       Queue[PromptCommand]
	Transactions  Queue[Transaction]
	Triggers      Queue[TriggerCommand]
}

func (s *State) Clone() State {
	players := slices.Clone(s.Players)
	for i := range players {
		players[i].Positions = slices.Clone(players[i].Positions)
	}

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

	prompts := slices.Clone(s.Prompts)
	for i := range prompts {
		prompts[i].Context = prompts[i].Context.Clone()
	}

	var activeContext *Context
	if s.ActiveContext != nil {
		cloned := s.ActiveContext.Clone()
		activeContext = &cloned
	}

	return State{
		ActiveContext: activeContext,
		Actors:        actors,
		Commands:      commands,
		Modifiers:     modifiers,
		Players:       players,
		Prompts:       prompts,
		Transactions:  transactions,
		Triggers:      triggers,
	}
}

func (s State) FindPlayerByID(id uuid.UUID) (Player, bool) {
	if id == uuid.Nil {
		return Player{}, false
	}

	for _, p := range s.Players {
		if p.ID == id {
			return p, true
		}
	}

	return Player{}, false
}
func (s State) FindActorByID(id uuid.UUID) (Actor, bool) {
	if id == uuid.Nil {
		return Actor{}, false
	}

	for _, a := range s.Actors {
		if a.ID == id {
			return a, true
		}
	}

	return Actor{}, false
}
func (s State) FindActorsWhere(g Game, where Filter[Actor], context Context) []Actor {
	actors := []Actor{}
	for _, a := range s.Actors {
		if where(g, a, context) {
			actors = append(actors, a)
		}
	}

	return actors
}

func (s *State) UpdatePlayer(id uuid.UUID, updater func(Player) Player) {
	for i, p := range s.Players {
		if p.ID == id {
			s.Players[i] = updater(p)
		}
	}
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
