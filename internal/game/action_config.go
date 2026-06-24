package game

import (
	"math/rand/v2"
)

type ActionConfig struct {
	Name string

	Affinity       Affinity
	Stat           Stat
	Power          float64
	BypassAccuracy bool
	Accuracy       float64
	CritChance     float64
	CritModifier   float64
	Priority       int
}

type AccuracyResult struct {
	Success      bool
	Accuracy     float64
	AccuracyRoll float64
	Critical     bool
	CriticalRoll float64
	Target       Actor
}

type DamageResult struct {
	AccuracyResult
	Damage            float64
	Random            float64
	Affinity          float64
	AffinityStage     int
	BaseAffinityStage int
}

func (ac ActionConfig) GetBaseDamage(source, target Actor, useBaseStats bool) float64 {
	adp_ratio := ac.Stat.GetRatio(source, target, useBaseStats) * ac.Power
	level_mod := float64(source.Level*2)/5 + 2
	base := (adp_ratio*level_mod)/50 + 2
	return base
}

func (ac ActionConfig) GetAccuracy(source, target Actor, useBaseStats bool) float64 {
	ratio := Accuracy.GetRatio(source, target, useBaseStats)
	return ratio * ac.Accuracy
}

func (ac ActionConfig) GetAccuracyResult(source, target Actor) AccuracyResult {
	accuracy_roll := rand.Float64() * 100
	critical_roll := rand.Float64() * 100
	critical := ac.CritChance > critical_roll

	accuracy := ac.GetAccuracy(source, target, critical)
	success := ac.BypassAccuracy || accuracy > accuracy_roll

	if !success {
		critical = false
	}

	return AccuracyResult{
		Success:      success,
		Accuracy:     accuracy,
		AccuracyRoll: accuracy_roll,
		Critical:     critical,
		CriticalRoll: critical_roll,
		Target:       target,
	}
}

func (ac ActionConfig) GetDamageResult(source, target Actor, targets []Actor) DamageResult {
	accuracy := ac.GetAccuracyResult(source, target)
	affinity, total_stage, base_stage := ac.Affinity.GetAffinityModifier(source, target)
	base := ac.GetBaseDamage(source, target, accuracy.Critical)
	damage := base * affinity

	if accuracy.Critical {
		damage = damage * ac.CritModifier
	}

	if len(targets) > 1 {
		damage = damage * 0.75
	}

	if !accuracy.Success || target.IsProtected {
		damage = 0.0
	}

	return DamageResult{
		AccuracyResult:    accuracy,
		Affinity:          affinity,
		AffinityStage:     total_stage,
		BaseAffinityStage: base_stage,
		Damage:            damage,
		Random:            rand.Float64()*(1.05-0.8) + 0.8,
	}
}
