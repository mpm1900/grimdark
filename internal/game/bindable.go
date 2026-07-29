package game

import (
	"github.com/google/uuid"
)

type Bindable[P any] struct {
	ID      uuid.UUID `json:"ID"`
	Context Context   `json:"context"`
	Payload P         `json:"payload"`
}

type Resolveable interface {
	Filter(*Game, Context) bool
	Delta(*Game, Context) []uuid.UUID
}

func bind[P any](payload P, context Context) Bindable[P] {
	return Bindable[P]{
		ID:      uuid.New(),
		Context: context,
		Payload: payload,
	}
}

func resolve(g *Game, context Context, res Resolveable) []uuid.UUID {
	if !res.Filter(g, context) {
		return []uuid.UUID{}
	}

	return res.Delta(g, context)
}
