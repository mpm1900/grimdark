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
	CriticalChance Stat = "critical-chance"
	CriticalDamage Stat = "critical-damage"
	Range          Stat = "range"
)

var mappedStats map[Stat]struct{} = map[Stat]struct{}{
	Health:         {},
	Speed:          {},
	Melee:          {},
	Ranged:         {},
	Special:        {},
	MartialDefense: {},
	SpecialDefense: {},
}

var percentStats map[Stat]struct{} = map[Stat]struct{}{
	Accuracy:       {},
	Evasion:        {},
	CriticalChance: {},
}

var statDefenses map[Stat]Stat = map[Stat]Stat{
	Melee:    MartialDefense,
	Ranged:   MartialDefense,
	Special:  SpecialDefense,
	Evasion:  Accuracy,
	Accuracy: Evasion,
}

func (s Stat) GetDefense() Stat {
	defense, ok := statDefenses[s]
	if !ok {
		return s
	}

	return defense
}

func (s Stat) GetRatio(source, target Actor, useBaseStats bool) float64 {
	source_value := source.Stats[s]
	target_value := target.Stats[s.GetDefense()]

	if useBaseStats {
		base := target.UnmodifiedStats[s.GetDefense()]
		if target_value > base {
			target_value = base
		}
	}

	return source_value / target_value
}
