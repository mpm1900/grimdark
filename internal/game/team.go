package game

import "github.com/google/uuid"

type ActorConfig struct {
	Class   uuid.UUID
	Items   []uuid.UUID
	Name    string
	WeaponL uuid.UUID
	WeaponR uuid.UUID
}

type TeamConfig struct {
	User   User
	Actors []ActorConfig
}
