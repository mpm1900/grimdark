package instance

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

type ResponseType string

const (
	ResponseTypeGame            = "game"
	ResponseTypeClients         = "clients"
	ResponseTypeJoinSuccess     = "join-success"
	ResponseTypeValidateContext = "validate-context"
	ResponseTypeTargetIDs       = "target-IDs"

	ResponseTypeNewChat = "new-chat"
)

type Response struct {
	Type    ResponseType   `json:"type"`
	Game    *game.GameJSON `json:"state"`
	Clients []*Client      `json:"clients"`
	Valid   *bool          `json:"valid"`
	Context *game.Context  `json:"context"`
}

func NewGameMessage(client *Client, g *game.Game) Response {
	json := g.ToJSON()

	return Response{
		Type:    ResponseTypeGame,
		Game:    &json,
		Clients: nil,
	}
}

func NewClientsMessage(clients []*Client) Response {
	return Response{
		Type:    ResponseTypeClients,
		Game:    nil,
		Clients: clients,
	}
}

func PostRegisterMessage(client *Client, g *game.Game) Response {
	json := g.ToJSON()

	return Response{
		Type:    ResponseTypeJoinSuccess,
		Game:    &json,
		Clients: []*Client{client},
	}
}

func TargetIDsResponse(client *Client, context game.Context, targetIDs []uuid.UUID) Response {
	context.ActorIDs = targetIDs
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
