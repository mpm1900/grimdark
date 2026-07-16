package game

import (
	"encoding/json"
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
