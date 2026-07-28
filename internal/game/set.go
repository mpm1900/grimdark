package game

import (
	"maps"
	"slices"
)

type Set[T comparable] map[T]struct{}

func (s Set[T]) Push(v T) {
	s[v] = struct{}{}
}

func (s Set[T]) Delete(v T) {
	delete(s, v)
}

func (s Set[T]) Has(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) Size() int {
	return len(s)
}

func (s Set[T]) ToArray() []T {
	return slices.Collect(maps.Keys(s))
}
