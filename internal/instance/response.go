package instance

import (
	"grimdark/internal/game"
)

type ResponseType string

const (
	// system messages
	ResponseTypeOnConnect   = "on-connect"
	ResponseTypePostConnect = "post-connect"
	ResponseTypeGameStart   = "game-start"

	// updates
	ResponseTypeGame    = "game"
	ResponseTypeClients = "clients"

	// request responses
	ResponseTypeValidateContext = "validate-context"
	ResponseTypeTargetIDs       = "target-IDs"
)

type Response struct {
	Type  ResponseType   `json:"type"`
	Game  *game.GameJSON `json:"game"`
	Lobby *LobbyJSON     `json:"lobby"`

	Valid   *bool         `json:"valid"`
	Context *game.Context `json:"context"`
}

func NewGameMessage(client *Client, g game.GameJSON) Response {
	g.ForPlayer(client.ID)
	return Response{
		Type: ResponseTypeGame,
		Game: &g,
	}
}

func NewLobbyMessage(client *Client, lobby LobbyJSON) Response {
	lobby.Client = client
	return Response{
		Type:  ResponseTypeClients,
		Game:  nil,
		Lobby: &lobby,
	}
}

func OnConnectMessage(client *Client, lobby LobbyJSON) Response {
	lobby.Client = client
	return Response{
		Type:  ResponseTypeOnConnect,
		Game:  nil,
		Lobby: &lobby,
	}
}
func PostConnectMessage(client *Client, g game.GameJSON) Response {
	g.ForPlayer(client.ID)
	lobby := LobbyJSON{
		Client:  client,
		Players: []*Client{},
	}
	return Response{
		Type:  ResponseTypePostConnect,
		Game:  &g,
		Lobby: &lobby,
	}
}

func TargetIDsResponse(client *Client, context game.Context) Response {
	return Response{
		Type:    ResponseTypeTargetIDs,
		Context: &context,
	}
}

func ValidateContextMessage(client *Client, context game.Context, valid bool) Response {
	return Response{
		Type:    ResponseTypeValidateContext,
		Context: &context,
		Valid:   &valid,
	}
}
