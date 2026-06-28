package instance

import (
	"context"
	"grimdark/internal/game"
	"maps"
	"slices"
	"time"

	"github.com/google/uuid"
)

type Instance struct {
	ID      uuid.UUID             `json:"ID"`
	ctx     context.Context       `json:"-"`
	Clients map[uuid.UUID]*Client `json:"clients,omitempty"`
	Game    game.Game             `json:"game"`
	Tick    time.Duration         `json:"-"`
	onEmpty func(uuid.UUID)       `json:"-"`

	Register    chan *Client `json:"-"`
	Unregister  chan *Client `json:"-"`
	ReadRequest chan Request `json:"-"`
}

type InstanceJSON struct {
	ID      uuid.UUID             `json:"ID"`
	Clients map[uuid.UUID]*Client `json:"clients,omitempty"`
}

func NewInstance(ctx context.Context, id uuid.UUID, onEmpty func(uuid.UUID)) *Instance {
	return &Instance{
		ID:          id,
		ctx:         ctx,
		Clients:     make(map[uuid.UUID]*Client),
		onEmpty:     onEmpty,
		Register:    make(chan *Client),
		Unregister:  make(chan *Client),
		ReadRequest: make(chan Request),

		Game: game.NewGame(),
		Tick: time.Second / 2,
	}
}

func (i Instance) ToJSON() InstanceJSON {
	return InstanceJSON{
		ID:      i.ID,
		Clients: maps.Clone(i.Clients),
	}
}

func (i *Instance) RegisterClient(client *Client) {
	if existing, ok := i.Clients[client.ID]; ok && existing != client {
		existing.cancel()
	}
	i.Clients[client.ID] = client
}

func (i *Instance) UnregisterClient(client *Client) bool {
	existing, ok := i.Clients[client.ID]
	if !ok || existing != client {
		return false
	}

	delete(i.Clients, client.ID)
	return true
}

func (i *Instance) BroadcastGame() {
	json := i.Game.ToJSON()
	for _, client := range i.Clients {
		if !client.TryWriteResponse(NewGameMessage(client, json)) {
			// If we can't send, it's usually better to just log it for now
			// rather than immediately unregistering, unless the client is truly dead.
			// i.UnregisterClient(client)
		}
	}
}

func (i *Instance) PostRegister(client *Client) {
	json := i.Game.ToJSON()
	client.TryWriteResponse(PostRegisterMessage(client, json))
}

func (i *Instance) TargetIDsResponse(clientID uuid.UUID, context game.Context) {
	client, ok := i.Clients[clientID]
	if !ok {
		return
	}

	client.TryWriteResponse(TargetIDsResponse(client, context))
}
func (i *Instance) ValidateContextResponse(clientID uuid.UUID, context game.Context, valid bool) {
	client, ok := i.Clients[clientID]
	if !ok {
		return
	}

	client.TryWriteResponse(ValidateContextMessage(client, context, valid))
}

func (i *Instance) BroadcastClients() {
	clients := slices.Collect(maps.Values(i.Clients))
	for _, client := range i.Clients {
		client.TryWriteResponse(NewClientsMessage(clients))
	}
}

const (
	state = iota
	clients
	none
)

func (i *Instance) Run() {
	for {
		select {
		case client := <-i.Register:
			_, exists := i.Game.GetPlayer(client.ID)
			if !exists && len(i.Game.State().Players) >= 2 {
				continue
			}

			i.RegisterClient(client)
			i.BroadcastClients()

			if !exists {
				player := game.NewPlayer()
				player.ID = client.ID
				player.User = *client.User

				i.Game.AddPlayers(player)
				i.BroadcastGame()
			}

			i.PostRegister(client)
		case client := <-i.Unregister:
			removed := i.UnregisterClient(client)
			if !removed {
				continue
			}

			if len(i.Clients) == 0 {
				if i.onEmpty != nil {
					i.onEmpty(i.ID)
				}
				return
			}

			i.BroadcastClients()
		case request := <-i.ReadRequest:
			switch Reducer(i, request) {
			case state:
				i.BroadcastGame()
			case clients:
				i.BroadcastClients()
			case none:
			}
		}
	}
}

func (i *Instance) FlushGame() bool {
	for i.Game.Next() {
		i.BroadcastGame()
		time.Sleep(i.Tick)
	}

	return len(i.Game.State().Prompts) > 0
}

func (i *Instance) RunGameActions() {
	if i.Game.Status == game.GameStatusRunning {
		return
	}

	i.Game.Status = game.GameStatusRunning
	i.BroadcastGame()

	defer func() {
		i.Game.Status = game.GameStatusIdle
		i.BroadcastGame()
	}()

resolveStep:
	for {
		if i.FlushGame() {
			break resolveStep
		}

		switch i.Game.Phase {
		case game.PhaseMain:
			i.Game.NextPhase()
			i.BroadcastGame()
			continue

		case game.PhaseEnd:
			i.Game.EndPhase()

			if i.FlushGame() {
				break resolveStep
			}

			i.Game.NextPhase()
			i.BroadcastGame()
			continue

		case game.PhaseCleanup:
			time.Sleep(i.Tick)
			i.Game.NextTurn()
			if i.Game.IsReadyToRun() {
				i.RunGameActions()
			} else {
				i.BroadcastGame()
				break resolveStep
			}

		default:
			i.Game.NextPhase()
			i.BroadcastGame()
		}
	}
}
