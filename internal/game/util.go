package game

import (
	"math/rand/v2"
)

func P[T any](v T) *T {
	return &v
}

func Chance(chance float64) bool {
	roll := rand.Float64()
	return chance >= roll
}
