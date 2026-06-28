package game

type Augment string

type augmentvalue struct {
	increase *Stat
	decrease *Stat
}

const (
	AugmentDefault Augment = "default"
)

var AUGMENT_MATRIX map[Augment]augmentvalue = map[Augment]augmentvalue{
	AugmentDefault: {
		increase: nil,
		decrease: nil,
	},
}

func (a Augment) GetMultiplier(stat Stat) float64 {
	value, ok := AUGMENT_MATRIX[a]
	if !ok {
		return 1.0
	}

	if value.decrease != nil && *value.decrease == stat {
		return 0.91
	}
	if value.increase != nil && *value.increase == stat {
		return 1.1
	}

	return 1.0
}
