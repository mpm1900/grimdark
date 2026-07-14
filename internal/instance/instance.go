package instance

import (
	"context"
	"grimdark/internal/game"
	"time"

	"github.com/google/uuid"
)

type InstanceStatus string

const (
	InstanceStatusInit     InstanceStatus = "init"
	InstanceStatusRunning  InstanceStatus = "running"
	InstanceStatusComplete InstanceStatus = "complete"
)

type Instance struct {
	ID      uuid.UUID
	ctx     context.Context
	Status  InstanceStatus
	Lobby   Lobby
	Game    *game.Game
	Tick    time.Duration
	onEmpty func(uuid.UUID)

	Register    chan *Client
	Unregister  chan *Client
	ReadRequest chan Request
}

type InstanceJSON struct {
	ID     uuid.UUID      `json:"ID"`
	Status InstanceStatus `json:"status"`
	Lobby  LobbyJSON      `json:"lobby"`
}

func NewInstance(ctx context.Context, id uuid.UUID, onEmpty func(uuid.UUID)) *Instance {
	i := &Instance{
		ID:          id,
		ctx:         ctx,
		Status:      InstanceStatusInit,
		onEmpty:     onEmpty,
		Register:    make(chan *Client),
		Unregister:  make(chan *Client),
		ReadRequest: make(chan Request),

		Game: game.NewGame(id),
		Tick: time.Second / 2,
	}

	i.Lobby = NewLobby(i)
	return i
}

func (i Instance) ToJSON() InstanceJSON {
	return InstanceJSON{
		ID:     i.ID,
		Lobby:  i.Lobby.ToJSON(),
		Status: i.Status,
	}
}

func (i *Instance) RegisterClient(client *Client) {
	if existing, ok := i.Lobby.GetClient(client.ID); ok && existing != client {
		existing.cancel()
	}
	i.Lobby.AddClient(client)
}

func (i *Instance) UnregisterClient(client *Client) bool {
	return i.Lobby.RemoveClient(client.ID)
}

func (i *Instance) BroadcastGame() {
	json := i.Game.ToJSON()
	for _, client := range i.Lobby.Clients() {
		if !client.TryWriteResponse(NewGameMessage(client, json)) {
			// If we can't send, it's usually better to just log it for now
			// rather than immediately unregistering, unless the client is truly dead.
			// i.UnregisterClient(client)
		}
	}
}

func (i *Instance) OnConnectResponse(client *Client) {
	client.TryWriteResponse(OnConnectMessage(client, i.Lobby.ToJSON()))
}
func (i *Instance) PostConnectResponse(client_id uuid.UUID) {
	client, ok := i.Lobby.GetClient(client_id)
	if !ok {
		return
	}

	client.TryWriteResponse(PostConnectMessage(client, i.Game.ToJSON()))
}

func (i *Instance) TargetIDsResponse(client_id uuid.UUID, context game.Context) {
	client, ok := i.Lobby.GetClient(client_id)
	if !ok {
		return
	}

	client.TryWriteResponse(TargetIDsResponse(client, context))
}
func (i *Instance) ValidateContextResponse(client_id uuid.UUID, context game.Context, valid bool) {
	client, ok := i.Lobby.GetClient(client_id)
	if !ok {
		return
	}

	client.TryWriteResponse(ValidateContextMessage(client, context, valid))
}

func (i *Instance) BroadcastLobby() {
	for _, client := range i.Lobby.Clients() {
		client.TryWriteResponse(NewLobbyMessage(client, i.Lobby.ToJSON()))
	}
}
func (i *Instance) BroadcastGameStart() {
	for _, client := range i.Lobby.Clients() {
		client.TryWriteResponse(GameStartMessage(client))
	}
}

func (i *Instance) Run() {
	for {
		select {
		case client := <-i.Register:
			_, exists := i.Game.GetPlayer(client.ID)
			if !exists && len(i.Game.State().Players) >= 2 {
				continue
			}

			i.RegisterClient(client)
			i.BroadcastLobby()

			if !exists {
				player := game.NewPlayer(*client.User)
				i.Game.AddPlayers(player)
				i.BroadcastGame()
			}

			i.OnConnectResponse(client)
			if exists {
				i.PostConnectResponse(client.ID)
			}
		case client := <-i.Unregister:
			removed := i.UnregisterClient(client)
			if !removed {
				continue
			}

			if len(i.Lobby.Players) == 0 {
				if i.onEmpty != nil {
					i.onEmpty(i.ID)
				}
				continue
			}

			i.BroadcastLobby()
		case request := <-i.ReadRequest:
			Reducer(i, request)
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
		if i.Game.Status == game.GameStatusRunning {
			i.Game.Status = game.GameStatusIdle
		}
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
