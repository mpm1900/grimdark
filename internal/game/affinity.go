package game

type Affinity string

const (
	Arcane    Affinity = "arcane"
	Cryo      Affinity = "cryo"
	Fire      Affinity = "fire"
	Kinetic   Affinity = "kinetic"
	Lightning Affinity = "lightning"
	Poison    Affinity = "poison"
	Psychic   Affinity = "psychic"
)

// AFFINITY_MATRIX maps defender affinity -> incoming affinity -> stage modifier.
var AFFINITY_MATRIX = map[Affinity]map[Affinity]int{
	Arcane: {
		Arcane:    2,
		Fire:      2,
		Lightning: 2,
		Poison:    -2,
		Psychic:   -2,
	},
	Cryo: {
		Cryo:      -2,
		Fire:      -2,
		Kinetic:   -2,
		Lightning: 2,
		Poison:    2,
	},
	Fire: {
		Arcane:  -2,
		Cryo:    2,
		Fire:    -2,
		Poison:  2,
		Psychic: -2,
	},
	Kinetic: {
		Cryo:      2,
		Lightning: -2,
		Poison:    2,
		Psychic:   -2,
	},
	Lightning: {
		Arcane:    -2,
		Cryo:      -2,
		Kinetic:   2,
		Lightning: -2,
		Psychic:   2,
	},
	Poison: {
		Arcane:  2,
		Cryo:    -2,
		Fire:    2,
		Kinetic: -2,
		Poison:  -2,
	},
	Psychic: {
		Arcane:    2,
		Fire:      2,
		Kinetic:   2,
		Lightning: -2,
		Poison:    -2,
		Psychic:   -2,
	},
}

func (a Affinity) GetBaseModifier(target Actor) int {
	result := 0

	for target_affinity := range target.Class.Affinities {
		stage, ok := AFFINITY_MATRIX[a][target_affinity]
		if !ok {
			continue
		}

		result += stage
	}

	return result
}

func (a Affinity) GetAffinityModifier(source, target Actor) (float64, int, int) {
	immunity, has_immunity := target.AffinityImmunities[a]
	if has_immunity {
		return immunity, 0, 0
	}

	base := a.GetBaseModifier(target)
	source_damage := source.GetAffinityDamage(a)
	target_resistance := target.GetAffinityResistance(a)
	total := base + source_damage - target_resistance

	return MapStage(total, 2, 1), total, base
}
