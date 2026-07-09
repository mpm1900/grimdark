package instance

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

type RequestType = string

const (
	PushAction    RequestType = "push-action"
	CancelAction  RequestType = "cancel-action"
	ResolvePrompt RequestType = "resolve-prompt"
	LoadTeam      RequestType = "load-team"

	RunGameActions RequestType = "run-game-actions" // TEMP

	GetTargets      RequestType = "get-targets"
	ValidateContext RequestType = "validate-context"
)

type Request struct {
	Type       RequestType      `json:"type"`
	ClientID   uuid.UUID        `json:"client_ID"`
	Context    game.Context     `json:"context"`
	TeamConfig *game.TeamConfig `json:"team_config"`
}
