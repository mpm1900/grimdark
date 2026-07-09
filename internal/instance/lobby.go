package instance

import (
	"maps"
	"slices"

	"github.com/google/uuid"
)

type Lobby struct {
	Client  *Client
	Clients map[uuid.UUID]*Client
}

type LobbyJSON struct {
	Client  *Client   `json:"client"`
	Clients []*Client `json:"clients"`
}

func NewLobby() Lobby {
	return Lobby{
		Client:  nil,
		Clients: make(map[uuid.UUID]*Client),
	}
}

func (l Lobby) ToJSON() LobbyJSON {
	return LobbyJSON{
		Client:  l.Client,
		Clients: slices.Collect(maps.Values(l.Clients)),
	}
}
