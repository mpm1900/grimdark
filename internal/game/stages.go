package game

func mapStage(stage int, mod, input float64) float64 {
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

	return mapStage(stage, sb.Mod, input)
}

func (sb *StageBuilder[T]) ResolveAll(input map[T]float64) map[T]float64 {
	output := make(map[T]float64, len(input))
	for key, i := range input {
		output[key] = sb.Resolve(key, i)
	}

	return output
}

type IOStageBuilder[T comparable] struct {
	I *StageBuilder[T]
	O *StageBuilder[T]
}

func NewIOStageBuilder[T comparable](i, o map[T]int) IOStageBuilder[T] {
	return IOStageBuilder[T]{
		I: NewStageBuilder(i),
		O: NewStageBuilder(o),
	}
}

func (iosb *IOStageBuilder[T]) Resolve(key T, input float64) float64 {
	o, ok := iosb.O.Stages[key]
	if !ok {
		o = 0
	}
	i, ok := iosb.I.Stages[key]
	if !ok {
		i = 0
	}

	return mapStage(o-i, iosb.O.Mod, input)
}

func (iosb *IOStageBuilder[T]) ResolveAll(input map[T]float64) map[T]float64 {
	output := make(map[T]float64, len(input))
	for key, i := range input {
		output[key] = iosb.Resolve(key, i)
	}

	return output
}

type ABStageBuilder struct {
	IOStageBuilder[struct{}]
}

func NewABStageBuilder(a, b int) ABStageBuilder {
	return ABStageBuilder{
		IOStageBuilder: IOStageBuilder[struct{}]{
			I: NewStageBuilder(map[struct{}]int{
				{}: a,
			}),
			O: NewStageBuilder(map[struct{}]int{
				{}: b,
			}),
		}}
}

func (absb *ABStageBuilder) Resolve(input float64) float64 {
	return absb.IOStageBuilder.Resolve(struct{}{}, input)
}
