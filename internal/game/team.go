package game

import "github.com/google/uuid"

type ActorConfig struct {
	Class   uuid.UUID   `json:"class"`
	Items   []uuid.UUID `json:"items"`
	Name    string      `json:"name"`
	WeaponL uuid.UUID   `json:"weapon_l"`
	WeaponR uuid.UUID   `json:"weapon_r"`
}

type TeamConfig struct {
	Actors []ActorConfig `json:"actors"`
	Name   string        `json:"name"`
}
