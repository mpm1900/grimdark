package instance

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

type RequestType = string

const (
	PostConnect RequestType = "post-connect"
	LobbyReady  RequestType = "ready"

	PushAction    RequestType = "push-action"
	CancelAction  RequestType = "cancel-action"
	TurnReady     RequestType = "turn-ready"
	ResolvePrompt RequestType = "resolve-prompt"

	GetTargets      RequestType = "get-targets"
	ValidateContext RequestType = "validate-context"
)

type Request struct {
	ClientID   uuid.UUID        `json:"client_ID"`
	Context    game.Context     `json:"context"`
	RequestID  uuid.UUID        `json:"request_ID"`
	TeamConfig *game.TeamConfig `json:"team_config"`
	Type       RequestType      `json:"type"`
}
