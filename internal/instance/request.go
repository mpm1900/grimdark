package instance

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

type RequestType = string

const (
	SetTeam        RequestType = "set-team"
	Reset          RequestType = "reset"
	PushAction     RequestType = "push-action"
	CancelAction   RequestType = "cancel-action"
	RunGameActions RequestType = "run-game-actions" // TEMP
	ResolvePrompt  RequestType = "resolve-prompt"

	GetTargets      RequestType = "get-targets"
	ValidateContext RequestType = "validate-context"
)

type Request struct {
	Type     RequestType  `json:"type"`
	ClientID uuid.UUID    `json:"client_ID"`
	Context  game.Context `json:"context"`
}
