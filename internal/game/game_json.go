package game

import (
	"bytes"
	"encoding/json"
	"reflect"
	"slices"

	"github.com/google/uuid"
)

type GameJSON struct {
	ActiveContext *Context               `json:"active_context"`
	Actors        []actorJSON            `json:"actors"`
	Commands      []Bindable[actionJSON] `json:"commands"`
	InstanceID    uuid.UUID              `json:"instance_ID"`
	Logs          []Bindable[Log]        `json:"logs"`
	Modifiers     []Modifier             `json:"modifiers"`
	Phase         GamePhase              `json:"phase"`
	Positions     []Position             `json:"positions"`
	PlayerID      uuid.UUID              `json:"player_ID"`
	Players       []playerJSON           `json:"players"`
	Prompts       []Bindable[actionJSON] `json:"prompts"`
	Status        GameStatus             `json:"status"`
	Turn          int                    `json:"turn"`
}

func (g *Game) ToJSON() GameJSON {
	state := g.State()
	players := make([]playerJSON, len(state.Players))
	for i, player := range state.Players {
		actors := g.GetActorsByPlayer(player.ID)
		players[i] = player.ToJSON(len(actors))
	}
	actors := make([]actorJSON, len(state.Actors))
	for i, actor := range state.Actors {
		actors[i] = actor.ToJSON(g)
	}

	prompts := []Bindable[actionJSON]{}
	for _, prompt := range state.Prompts {
		if prompt.Ready {
			continue
		}

		prompts = append(prompts, bind(prompt.Payload.ToJSON(g, Actor{}), prompt.Context))
	}
	commands := []Bindable[actionJSON]{}
	for _, command := range state.Commands {
		commands = append(commands, bind(command.Payload.ToJSON(g, Actor{}), command.Context))
	}

	return GameJSON{
		ActiveContext: state.ActiveContext,
		Actors:        actors,
		Commands:      commands,
		InstanceID:    g.InstanceID,
		Logs:          g.Logs,
		Modifiers:     g.meta.modifiers,
		Phase:         g.Phase,
		Positions:     slices.Clone(g.State().Positions),
		PlayerID:      uuid.Nil,
		Players:       players,
		Prompts:       prompts,
		Status:        g.Status,
		Turn:          g.Turn,
	}
}

func CloneGameJSON(value GameJSON) (GameJSON, error) {
	data, err := json.Marshal(value)
	if err != nil {
		return GameJSON{}, err
	}

	var clone GameJSON
	if err := json.Unmarshal(data, &clone); err != nil {
		return GameJSON{}, err
	}

	return clone, nil
}

func DiffGameJSON(previous, next GameJSON) (GameJSON, bool) {
	patch := next

	patch.Actors = diffActors(previous.Actors, next.Actors)

	logs, ok := diffLogs(previous.Logs, next.Logs)
	if !ok {
		return GameJSON{}, false
	}
	patch.Logs = logs

	if actionBindingsEqual(previous.Commands, next.Commands) {
		patch.Commands = nil
	} else {
		patch.Commands = nonNilSlice(next.Commands)
	}
	if jsonEqual(previous.Modifiers, next.Modifiers) {
		patch.Modifiers = nil
	} else {
		patch.Modifiers = nonNilSlice(next.Modifiers)
	}
	if jsonEqual(previous.Positions, next.Positions) {
		patch.Positions = nil
	} else {
		patch.Positions = nonNilSlice(next.Positions)
	}
	if jsonEqual(previous.Players, next.Players) {
		patch.Players = nil
	} else {
		patch.Players = nonNilSlice(next.Players)
	}
	if actionBindingsEqual(previous.Prompts, next.Prompts) {
		patch.Prompts = nil
	} else {
		patch.Prompts = nonNilSlice(next.Prompts)
	}

	return patch, true
}

func diffActors(previous, next []actorJSON) []actorJSON {
	previousByID := make(map[uuid.UUID]actorJSON, len(previous))
	for _, actor := range previous {
		previousByID[actor.ID] = actor
	}

	var changed []actorJSON
	for _, actor := range next {
		if old, ok := previousByID[actor.ID]; !ok || !jsonEqual(old, actor) {
			changed = append(changed, actor)
		}
	}

	return changed
}

func diffLogs(previous, next []Bindable[Log]) ([]Bindable[Log], bool) {
	if len(next) < len(previous) {
		return nil, false
	}
	for i, log := range previous {
		if !reflect.DeepEqual(log, next[i]) {
			return nil, false
		}
	}
	if len(next) == len(previous) {
		return nil, true
	}

	return next[len(previous):], true
}

func nonNilSlice[T any](values []T) []T {
	if values == nil {
		return []T{}
	}

	return values
}

func actionBindingsEqual(previous, next []Bindable[actionJSON]) bool {
	if len(previous) != len(next) {
		return false
	}
	for i := range previous {
		if !reflect.DeepEqual(previous[i].Context, next[i].Context) {
			return false
		}
		if !jsonEqual(previous[i].Payload, next[i].Payload) {
			return false
		}
	}

	return true
}

func jsonEqual(a, b any) bool {
	encodedA, err := json.Marshal(a)
	if err != nil {
		return false
	}
	encodedB, err := json.Marshal(b)
	if err != nil {
		return false
	}

	return bytes.Equal(encodedA, encodedB)
}

func (json *GameJSON) ForPlayer(player_ID uuid.UUID) {
	json.PlayerID = player_ID
	prompts := slices.Clone(json.Prompts)
	prompts = slices.DeleteFunc(prompts, func(p Bindable[actionJSON]) bool {
		return p.Context.PlayerID != player_ID
	})
	commands := slices.Clone(json.Commands)
	commands = slices.DeleteFunc(commands, func(p Bindable[actionJSON]) bool {
		return p.Context.PlayerID != player_ID
	})
	actors := []actorJSON{}
	for _, a := range json.Actors {
		a.IsPlayer = a.PlayerID == player_ID
		if a.IsPlayer || a.Seen {
			actors = append(actors, a)
		}
	}
	json.Prompts = prompts
	json.Commands = commands
	json.Actors = actors
}

func (g *Game) MarshalJSON() ([]byte, error) {
	return json.Marshal(g.ToJSON())
}
