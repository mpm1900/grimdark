package server

import (
	"context"
	"encoding/json"
	"grimdark/internal/auth"
	"grimdark/internal/db"
	"grimdark/internal/game"
	"grimdark/internal/team"
	"io"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
)

type TeamsHandler struct {
	mux     *http.ServeMux
	queries *db.Queries
}

func NewTeamsHandler(ctx context.Context, queries *db.Queries) *TeamsHandler {
	handler := &TeamsHandler{
		mux:     http.NewServeMux(),
		queries: queries,
	}

	return handler
}

func (th *TeamsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	th.mux.ServeHTTP(w, r)
}

func (th *TeamsHandler) HandleGetTeams(w http.ResponseWriter, r *http.Request) {
	logger := slog.Default()
	user, ok := auth.AuthenticatedUserFromContext(r.Context())
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	results, err := team.GetTeamsByUser(r.Context(), th.queries, user.ID)
	if err != nil {
		logger.Error("GetTeams: failed to fetch teams", "user_id", user.ID, "err", err)
		http.Error(w, "failed to fetch teams", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(results); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type upsertTeamBody struct {
	ID     *uuid.UUID      `json:"id"`
	Config game.TeamConfig `json:"config"`
}

func readUpsertTeamsBody(r *http.Request) (*upsertTeamBody, error) {
	req, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var body upsertTeamBody
	if err := json.Unmarshal(req, &body); err != nil {
		return nil, err
	}

	return &body, nil
}

func (th *TeamsHandler) HandleUpsertTeam(w http.ResponseWriter, r *http.Request) {
	logger := slog.Default()
	user, ok := auth.AuthenticatedUserFromContext(r.Context())
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	body, err := readUpsertTeamsBody(r)
	if err != nil {
		logger.Error("Signup: failed to read request body", "err", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	result, err := team.GetTeamByID(r.Context(), th.queries, body.ID)
	if err == nil {
		if result.UserID != user.ID {
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}

		// update
		result, err = team.UpdateTeam(r.Context(), th.queries, result.ID, body.Config)
	} else {
		// create
		result, err = team.CreateTeam(r.Context(), th.queries, user.ID, body.Config)
	}
	if err != nil {
		logger.Error("UpsertTeam: failed to upsert team", "user_id", user.ID, "err", err)
		http.Error(w, "failed to save team", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (th *TeamsHandler) HandleDeleteTeam(w http.ResponseWriter, r *http.Request) {
	user, ok := auth.AuthenticatedUserFromContext(r.Context())
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	team_id_str := r.PathValue("team_id")
	team_id, err := uuid.Parse(team_id_str)
	if err != nil {
		http.Error(w, "invalid team_id", http.StatusBadRequest)
		return
	}

	result, err := team.GetTeamByID(r.Context(), th.queries, &team_id)
	if err != nil {
		http.Error(w, "team not found", http.StatusNotFound)
		return
	}
	if result.UserID != user.ID {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}

	err = team.DeleteTeam(r.Context(), th.queries, team_id)
	if err != nil {
		http.Error(w, "error deleting team", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
