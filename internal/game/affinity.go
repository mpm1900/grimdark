package game

type Affinity string

const (
	Arcane    Affinity = "arcane"
	Blood     Affinity = "blood"
	Holy      Affinity = "holy"
	Fire      Affinity = "fire"
	Physical  Affinity = "physical"
	Lightning Affinity = "lightning"
	Poison    Affinity = "poison"
	Psychic   Affinity = "psychic"
)

var AFFINITY_MATRIX = map[Affinity]map[Affinity]int{
	Arcane: {
		Arcane:    0,
		Blood:     0,
		Holy:      2,
		Fire:      -2,
		Physical:  0,
		Lightning: 0,
		Poison:    0,
		Psychic:   0,
	},
	Blood: {
		Arcane:    0,
		Blood:     0,
		Holy:      0,
		Fire:      2,
		Physical:  0,
		Lightning: -2,
		Poison:    0,
		Psychic:   0,
	},
	Holy: {
		Arcane:    -2,
		Blood:     0,
		Holy:      0,
		Fire:      0,
		Physical:  0,
		Lightning: 0,
		Poison:    2,
		Psychic:   0,
	},
	Fire: {
		Arcane:    2,
		Blood:     -2,
		Holy:      0,
		Fire:      -2,
		Physical:  0,
		Lightning: 0,
		Poison:    -2,
		Psychic:   2,
	},
	Physical: {
		Arcane:    0,
		Blood:     0,
		Holy:      0,
		Fire:      0,
		Physical:  0,
		Lightning: 0,
		Poison:    0,
		Psychic:   -2,
	},
	Lightning: {
		Arcane:    0,
		Blood:     2,
		Holy:      0,
		Fire:      0,
		Physical:  -2,
		Lightning: 0,
		Poison:    0,
		Psychic:   0,
	},
	Poison: {
		Arcane:    0,
		Blood:     0,
		Holy:      -2,
		Fire:      2,
		Physical:  0,
		Lightning: 0,
		Poison:    -2,
		Psychic:   0,
	},
	Psychic: {
		Arcane:    0,
		Blood:     0,
		Holy:      0,
		Fire:      0,
		Physical:  2,
		Lightning: 0,
		Poison:    0,
		Psychic:   0,
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
