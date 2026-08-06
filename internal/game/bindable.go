package game

import (
	"github.com/google/uuid"
)

type Bindable[P any] struct {
	ID      uuid.UUID `json:"ID"`
	Context Context   `json:"context"`
	Payload P         `json:"payload"`
}

func bind[P any](payload P, context Context) Bindable[P] {
	return Bindable[P]{
		ID:      uuid.New(),
		Context: context,
		Payload: payload,
	}
}

type Delta interface {
	Filter(*Game, Context) bool
	Delta(*Game, Context) []uuid.UUID
}
