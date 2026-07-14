package team

import (
	"context"
	"fmt"
	"grimdark/internal/db"
	"grimdark/internal/game"
	"time"

	"github.com/google/uuid"
)

func GetTeamsByUser(ctx context.Context, queries *db.Queries, user_id uuid.UUID) ([]db.Team, error) {
	teams, err := queries.GetTeamsByUser(ctx, user_id)
	return teams, err
}

func GetTeamByID(ctx context.Context, queries *db.Queries, id *uuid.UUID) (db.Team, error) {
	if id == nil {
		return db.Team{}, fmt.Errorf("id is nil")
	}

	team, err := queries.GetTeamByID(ctx, *id)
	return team, err
}

func CreateTeam(ctx context.Context, queries *db.Queries, user_id uuid.UUID, config game.TeamConfig) (db.Team, error) {
	if config.Name == "" {
		config.Name = fmt.Sprintf("New Team (%s)", time.Now())
	}
	return queries.CreateTeam(ctx, db.CreateTeamParams{
		UserID: user_id,
		Config: config,
	})
}

func UpdateTeam(ctx context.Context, queries *db.Queries, team_id uuid.UUID, config game.TeamConfig) (db.Team, error) {
	return queries.UpdateTeam(ctx, db.UpdateTeamParams{
		ID:     team_id,
		Config: config,
	})
}

func DeleteTeam(ctx context.Context, queries *db.Queries, team_id uuid.UUID) error {
	return queries.DeleteTeam(ctx, team_id)
}
