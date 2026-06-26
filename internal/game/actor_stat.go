package game

type Stat string

const (
	Health         Stat = "health"
	Speed          Stat = "speed"
	Melee          Stat = "melee"
	Ranged         Stat = "ranged"
	Special        Stat = "special"
	MartialDefense Stat = "martial-defense"
	SpecialDefense Stat = "special-defense"
	Accuracy       Stat = "accuracy"
	Evasion        Stat = "evasion"
)

func (s Stat) GetDefense() Stat {
	switch s {
	case Health:
		return Health
	case Speed:
		return Speed
	case Melee:
		return MartialDefense
	case Ranged:
		return MartialDefense
	case MartialDefense:
		return MartialDefense
	case Special:
		return SpecialDefense
	case SpecialDefense:
		return SpecialDefense
	case Accuracy:
		return Evasion
	case Evasion:
		return Accuracy
	default:
		return Health
	}
}

func (s Stat) GetRatio(source, target Actor, useBaseStats bool) float64 {
	source_value := source.Stats[s]
	target_value := target.Stats[s.GetDefense()]

	if useBaseStats {
		base_target_value := target.UnmodifiedStats[s.GetDefense()]
		if target_value > base_target_value {
			target_value = base_target_value
		}
	}

	return source_value / target_value
}
