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

func MapStage(stage int, mod, input float64) float64 {
	if stage > 6 {
		stage = 6
	}
	if stage < -6 {
		stage = -6
	}
	if stage > 0 {
		return input * (float64(stage) + mod) / mod
	}
	if stage < 0 {
		return input * mod / (float64(-stage) + mod)
	}
	return input
}

func Factorial(n int) int {
	var result int = 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func GetCriticalChance(stage int) float64 {
	if stage > 3 {
		stage = 3
	}
	if stage < -3 {
		return -3
	}

	return 1.0 / float64(Factorial(4-stage))
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
