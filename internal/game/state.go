package game

import (
	"cmp"
	"slices"

	"github.com/google/uuid"
)

type State struct {
	ActiveContext *Context
	Actors        []Actor
	Commands      Queue[Command]
	Modifiers     []Modifier
	Positions     []Position
	Players       []Player
	Prompts       Queue[PromptCommand]
	Transactions  Queue[Transaction]
	Triggers      Queue[TriggerCommand]
}

func (s *State) Clone() State {
	players := slices.Clone(s.Players)
	positions := slices.Clone(s.Positions)

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
		triggers[i].ParentContext = triggers[i].ParentContext.Clone()
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
		Positions:     positions,
		Players:       players,
		Prompts:       prompts,
		Transactions:  transactions,
		Triggers:      triggers,
	}
}

func (s State) GetPlayer(id uuid.UUID) (Player, bool) {
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
func (s State) GetPosition(id uuid.UUID) (Position, bool) {
	if id == uuid.Nil {
		return Position{}, false
	}

	for _, p := range s.Positions {
		if p.ID == id {
			return p, true
		}
	}

	return Position{}, false
}
func (s State) GetPositionByActorID(actor_id uuid.UUID) (Position, bool) {
	for _, pos := range s.Positions {
		if pos.ActorID == actor_id {
			return pos, true
		}
	}

	return Position{}, false
}
func (s State) GetPositionsByPlayerID(player_id uuid.UUID) []Position {
	positions := []Position{}
	for _, pos := range s.Positions {
		if pos.PlayerID == player_id {
			positions = append(positions, pos)
		}
	}
	slices.SortStableFunc(positions, func(a, b Position) int {
		return cmp.Compare(a.Rank, b.Rank)
	})

	return positions
}
func (s State) GetOpenPositions(player_ID uuid.UUID) []Position {
	positions := []Position{}
	for _, pos := range s.Positions {
		if pos.PlayerID == player_ID && pos.ActorID == uuid.Nil {
			positions = append(positions, pos)
		}
	}

	return positions
}
func (s State) GetOpenPositionIDs(player_ID uuid.UUID) []uuid.UUID {
	positions := []uuid.UUID{}
	for _, pos := range s.GetOpenPositions(player_ID) {
		positions = append(positions, pos.ID)
	}

	return positions
}
func (s State) GetActor(id uuid.UUID) (Actor, bool) {
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
func (s State) FindActors(g *Game, where Filter[Actor], context Context) []Actor {
	actors := []Actor{}
	for _, a := range s.Actors {
		if where(g, a, context) {
			actors = append(actors, a)
		}
	}

	return actors
}
func (s State) FindCommands(g *Game, where Filter[Command], context Context) []Command {
	commands := []Command{}
	for _, c := range s.Commands {
		if where(g, c, context) {
			commands = append(commands, c)
		}
	}

	return commands
}

func (s *State) UpdatePlayer(id uuid.UUID, updater func(Player) Player) {
	for i, p := range s.Players {
		if p.ID == id {
			s.Players[i] = updater(p)
		}
	}
}
func (s *State) UpdatePosition(id uuid.UUID, updater func(Position) Position) {
	for i, p := range s.Positions {
		if p.ID == id {
			s.Positions[i] = updater(p)
		}
	}
}
func (s *State) UpdatePositionActor(position_id uuid.UUID, actor_id uuid.UUID) {
	s.UpdatePosition(position_id, func(p Position) Position {
		p.ActorID = actor_id
		return p
	})
}
func (s *State) UpdatePositionWhere(where func(Position) bool, updater func(Position) Position) {
	for i, p := range s.Positions {
		if where(p) {
			s.Positions[i] = updater(p)
		}
	}
}

func (s *State) SetPosition(position_id uuid.UUID, actor Actor) (uuid.UUID, bool) {
	var evicted_id uuid.UUID
	updated := false

	var current Position
	if position_id != uuid.Nil {
		var ok bool
		current, ok = s.GetPosition(position_id)
		if !ok || current.PlayerID != actor.PlayerID {
			return evicted_id, false
		}
	}

	for _, pos := range s.Positions {
		if pos.ActorID == actor.ID && pos.ID != position_id {
			s.UpdatePositionActor(pos.ID, uuid.Nil)
			updated = true
		}
	}

	if position_id == uuid.Nil {
		return evicted_id, updated || actor.IsActive()
	}

	if current.ActorID != uuid.Nil && current.ActorID != actor.ID {
		evicted_id = current.ActorID
	}

	s.UpdatePositionActor(position_id, actor.ID)
	return evicted_id, true
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
