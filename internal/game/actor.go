package game

import (
	"maps"

	"github.com/google/uuid"
)

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

type Status string

const (
	StatusNone   Status = "none"
	StatusBurned Status = "burned"
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
		base_target_value := target.Unmodified[s.GetDefense()]
		if target_value > base_target_value {
			target_value = base_target_value
		}
	}

	return source_value / target_value
}

type ActorDef struct {
	ID         uuid.UUID
	Name       string
	Affinities map[Affinity]struct{}
	Stats      map[Stat]float64
}

type Actor struct {
	ActorDef
	Level      int
	PlayerID   uuid.UUID
	PositionID uuid.UUID

	AffinityDamage     map[Affinity]int
	AffinityResistance map[Affinity]int
	Unmodified         map[Stat]float64
	Stages             map[Stat]int
	Aux                map[Stat]float64
	Stats              map[Stat]float64

	Damage      float64
	Status      Status
	IsAlive     bool
	IsProtected bool
}

func NewActorDef() ActorDef {
	return ActorDef{
		ID:         uuid.New(),
		Affinities: map[Affinity]struct{}{},
		Stats: map[Stat]float64{
			Health:         100,
			Speed:          100,
			Melee:          100,
			Ranged:         100,
			Special:        100,
			MartialDefense: 100,
			SpecialDefense: 100,
			Accuracy:       1,
			Evasion:        1,
		},
	}
}

func NewActor(playerID uuid.UUID, def ActorDef) Actor {
	return Actor{
		ActorDef:           def,
		PlayerID:           playerID,
		PositionID:         uuid.Nil,
		Level:              100,
		AffinityDamage:     map[Affinity]int{},
		AffinityResistance: map[Affinity]int{},
		Stages:             map[Stat]int{},
		Aux:                map[Stat]float64{},
		Stats:              maps.Clone(def.Stats),

		Damage:      0,
		Status:      StatusNone,
		IsAlive:     true,
		IsProtected: false,
	}
}

func (a Actor) Clone() Actor {
	return Actor{
		ActorDef: ActorDef{
			ID:         a.ID,
			Name:       a.Name,
			Affinities: maps.Clone(a.Affinities),
			Stats:      maps.Clone(a.ActorDef.Stats),
		},

		AffinityDamage:     maps.Clone(a.AffinityDamage),
		AffinityResistance: maps.Clone(a.AffinityResistance),
		Unmodified:         maps.Clone(a.Unmodified),
		Stages:             maps.Clone(a.Stages),
		Aux:                maps.Clone(a.Aux),
		Stats:              maps.Clone(a.Stats),
		Level:              a.Level,
		PlayerID:           a.PlayerID,
		PositionID:         a.PositionID,
		Damage:             a.Damage,
		Status:             a.Status,
		IsAlive:            a.IsAlive,
		IsProtected:        a.IsProtected,
	}
}

func mapBaseStat(actor Actor, stat Stat, stats map[Stat]float64, aux float64) float64 {
	base := stats[stat]*2 + aux
	ratio := float64(actor.Level) / 100
	result := base*ratio + 5
	if stat == Health {
		result += float64(actor.Level)
	}
	return result
}

func (a *Actor) mapBaseStats() {
	a.Unmodified = maps.Clone(a.Stats)

	for stat, _ := range a.Stats {
		if stat == Accuracy || stat == Evasion {
			continue
		}

		aux, ok := a.Aux[stat]
		if !ok {
			aux = 0
		}

		a.Stats[stat] = mapBaseStat(*a, stat, a.Stats, aux)
		a.Unmodified[stat] = mapBaseStat(*a, stat, a.Unmodified, 0)
	}
}

func (a Actor) GetAffinityDamage(affinity Affinity) int {
	base := maps.Clone(a.AffinityDamage)
	for affinity := range a.Affinities {
		b, ok := base[affinity]
		if !ok {
			b = 0
		}

		base[affinity] = b + 1
	}

	value, ok := base[affinity]
	if ok {
		return value
	}

	return 0
}

func (a Actor) GetAffinityResistance(affinity Affinity) int {
	value, ok := a.AffinityResistance[affinity]
	if ok {
		return value
	}

	return 0
}

func (a Actor) GetRemainingHealth() float64 {
	health := a.Stats[Health]
	return health - a.Damage
}
