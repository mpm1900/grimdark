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
	Actions        Stat = "actions"
	CriticalChance Stat = "critical-chance"
	CriticalDamage Stat = "critical-damage"
	DamageReflect  Stat = "damage-reflect"
	EffectChance   Stat = "effect-chance"
)

type Stack string

const (
	Wounds Stack = "wounds"
)

// mapped stats are piped through the funciton that factors in level and other factors
var mappedStats map[Stat]struct{} = map[Stat]struct{}{
	Health:         {},
	Speed:          {},
	Melee:          {},
	Ranged:         {},
	Special:        {},
	MartialDefense: {},
	SpecialDefense: {},
}

// percent stats are x100 before sent to clients
var percentStats map[Stat]struct{} = map[Stat]struct{}{
	Accuracy:       {},
	Evasion:        {},
	CriticalChance: {},
	CriticalDamage: {},
	DamageReflect:  {},
	EffectChance:   {},
}

// gets the "opposite" stat
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

func (s Stat) GetRatio(source, target Actor, is_unmodified bool) float64 {
	source_value := source.Stats[s]
	target_value := target.Stats[s.GetDefense()]

	if is_unmodified {
		base := target.UnmodifiedStats[s.GetDefense()]
		if target_value > base {
			target_value = base
		}
	}

	return source_value / target_value
}
