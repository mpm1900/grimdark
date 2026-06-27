package game

type StageBuilder[T comparable] struct {
	Mod    float64
	Stages map[T]int
}

func NewStageBuilder[T comparable](initial map[T]int) *StageBuilder[T] {
	sb := StageBuilder[T]{
		Mod:    2,
		Stages: initial,
	}
	return &sb
}

func (sb *StageBuilder[T]) Map(key T, updater func(key T, value int) int) {
	stage, ok := sb.Stages[key]
	if !ok {
		return
	}

	sb.Stages[key] = updater(key, stage)
}

func (sb *StageBuilder[T]) MapWhere(filter func(key T, value int) bool, updater func(key T, value int) int) {
	for k, v := range sb.Stages {
		if filter(k, v) {
			sb.Stages[k] = updater(k, v)
		}
	}
}

func (sb *StageBuilder[T]) Resolve(key T, input float64) float64 {
	stage, ok := sb.Stages[key]
	if !ok {
		stage = 0.0
	}

	return MapStage(stage, sb.Mod, input)
}

func (sb *StageBuilder[T]) ResolveAll(input map[T]float64) map[T]float64 {
	output := make(map[T]float64, len(input))
	for key, i := range input {
		output[key] = sb.Resolve(key, i)
	}

	return output
}
