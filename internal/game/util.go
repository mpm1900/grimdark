package game

import (
	"math/rand/v2"

	"github.com/google/uuid"
)

func P[T any](v T) *T {
	return &v
}

func Chance(chance float64) bool {
	roll := rand.Float64()
	return chance >= roll
}

func NilifyUUID(uid uuid.UUID) *uuid.UUID {
	if uid == uuid.Nil {
		return nil
	}

	return &uid
}
