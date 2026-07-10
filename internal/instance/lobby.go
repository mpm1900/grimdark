package instance

import (
	"maps"
	"slices"

	"github.com/google/uuid"
)

type Lobby struct {
	Players    map[uuid.UUID]*Client
	Spectators map[uuid.UUID]*Client
}

type LobbyJSON struct {
	Client     *Client   `json:"client"`
	Players    []*Client `json:"players"`
	Spectators []*Client `json:"spectators"`
}

func NewLobby() Lobby {
	return Lobby{
		Players:    make(map[uuid.UUID]*Client),
		Spectators: make(map[uuid.UUID]*Client),
	}
}

func (l Lobby) ToJSON() LobbyJSON {
	return LobbyJSON{
		Client:     nil,
		Players:    slices.Collect(maps.Values(l.Players)),
		Spectators: slices.Collect(maps.Values(l.Spectators)),
	}
}

func (l *Lobby) Clients() map[uuid.UUID]*Client {
	clients := make(map[uuid.UUID]*Client, len(l.Players)+len(l.Spectators))
	maps.Copy(clients, l.Players)
	maps.Copy(clients, l.Spectators)
	return clients
}
func (l *Lobby) AddClient(client *Client) {
	if len(l.Players) < 2 {
		l.Players[client.ID] = client
	} else {
		l.Spectators[client.ID] = client
	}
}
func (l *Lobby) GetClient(client_id uuid.UUID) (*Client, bool) {
	p, ok := l.Players[client_id]
	if ok {
		return p, ok
	}

	s, ok := l.Spectators[client_id]
	return s, ok
}
func (l *Lobby) RemoveClient(client_id uuid.UUID) bool {
	for id, _ := range l.Players {
		if id == client_id {
			delete(l.Players, client_id)
			return true
		}
	}
	for id, _ := range l.Spectators {
		if id == client_id {
			delete(l.Spectators, client_id)
			return true
		}
	}

	return false
}
