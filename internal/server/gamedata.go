package server

import (
	"context"
	"encoding/json"
	"grimdark/internal/game/actors"
	"maps"
	"net/http"
	"slices"
)

type GamedataHandler struct {
	mux *http.ServeMux
}

func NewGamedataHandler(ctx context.Context) *GamedataHandler {
	handler := &GamedataHandler{
		mux: http.NewServeMux(),
	}

	return handler
}

func (dh *GamedataHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dh.mux.ServeHTTP(w, r)
}

func (dh *GamedataHandler) HandleGetActors(w http.ResponseWriter, r *http.Request) {
	actors := slices.Collect(maps.Values(actors.All))
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(actors); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
