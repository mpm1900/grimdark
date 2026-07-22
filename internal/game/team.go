package game

import "github.com/google/uuid"

type ActorConfig struct {
	Class   *uuid.UUID              `json:"class"`
	Items   []uuid.UUID             `json:"items"`
	Name    string                  `json:"name"`
	Weapons map[uuid.UUID]uuid.UUID `json:"weapons"`
}

type TeamConfig struct {
	Actors []ActorConfig `json:"actors"`
	Name   string        `json:"name"`
}
