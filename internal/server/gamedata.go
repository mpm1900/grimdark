package server

import (
	"context"
	"encoding/json"
	"grimdark/internal/game/actors"
	"grimdark/internal/game/items"
	"net/http"
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
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(actors.All); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (dh *GamedataHandler) HandleGetGlobalItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(items.Global); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
