package game

import "math/rand/v2"

func P[T any](v T) *T {
	return &v
}

func Chance(check float64) bool {
	chance := 1.0
	roll := rand.Float64()
	return chance >= roll
}
