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

var affinityMatrix = map[Affinity]map[Affinity]int{
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
		Lightning: 2,
		Poison:    -2,
		Psychic:   -2,
	},
}

func (a Affinity) GetBaseModifier(target Actor) int {
	result := 0

	for affinity := range target.Affinities {
		stage, ok := affinityMatrix[a][affinity]
		if !ok {
			stage = 0
		}

		result += stage
	}

	return result
}

func (a Affinity) GetAffinityModifier(source, target Actor) float64 {
	base := a.GetBaseModifier(target)
	source_damage := source.GetAffinityDamage(a)
	target_resistance := target.GetAffinityResistance(a)
	base += source_damage - target_resistance

	return mapStage(base, 2, 1)
}
