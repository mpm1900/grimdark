package instance

import (
	"maps"

	"github.com/google/uuid"
)

type Lobby struct {
	Players    map[uuid.UUID]*Client
	Spectators map[uuid.UUID]*Client
	ready      map[uuid.UUID]bool
}

type LobbyJSON struct {
	Client     *Client            `json:"client"`
	Players    []*Client          `json:"players"`
	Spectators []*Client          `json:"spectators"`
	Ready      map[uuid.UUID]bool `json:"ready"`
}

func NewLobby() Lobby {
	return Lobby{
		Players:    make(map[uuid.UUID]*Client),
		Spectators: make(map[uuid.UUID]*Client),
		ready:      make(map[uuid.UUID]bool),
	}
}

func (l Lobby) ToJSON() LobbyJSON {
	players := make([]*Client, 0, len(l.Players))
	for _, player := range l.Players {
		players = append(players, player)
	}

	spectators := make([]*Client, 0, len(l.Spectators))
	for _, spectator := range l.Spectators {
		spectators = append(spectators, spectator)
	}

	return LobbyJSON{
		Client:     nil,
		Players:    players,
		Spectators: spectators,
		Ready:      l.ready,
	}
}

func (l *Lobby) Clients() map[uuid.UUID]*Client {
	clients := make(map[uuid.UUID]*Client, len(l.Players)+len(l.Spectators))
	maps.Copy(clients, l.Players)
	maps.Copy(clients, l.Spectators)
	return clients
}
func (l *Lobby) GetClient(client_id uuid.UUID) (*Client, bool) {
	p, ok := l.Players[client_id]
	if ok {
		return p, ok
	}

	s, ok := l.Spectators[client_id]
	return s, ok
}
func (l *Lobby) Ready() bool {
	for _, ready := range l.ready {
		if !ready {
			return false
		}
	}
	return true
}

func (l *Lobby) AddClient(client *Client) {
	if len(l.Players) < 2 {
		client.Role = ClientRolePlayer
		l.ready[client.ID] = false
		l.Players[client.ID] = client
	} else {
		client.Role = ClientRoleSpectator
		l.Spectators[client.ID] = client
	}
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
func (l *Lobby) SetReady(client_id uuid.UUID) {
	l.ready[client_id] = true
}
