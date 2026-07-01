package game

import (
	"math/rand/v2"
)

type ActionConfig struct {
	Accuracy     *float64 `json:"accuracy"`
	Affinity     Affinity `json:"affinity"`
	CritChance   float64  `json:"crit_chance"`
	CritModifier float64  `json:"crit_modifier"`
	Description  string   `json:"description"`
	Hits         int      `json:"hits"`
	Lifesteal    float64  `json:"lifesteal"`
	Name         string   `json:"name"`
	Power        float64  `json:"power"`
	Priority     int      `json:"priority"`
	Recoil       float64  `json:"recoil"`
	Stat         Stat     `json:"stat"`
	TargetCount  int      `json:"target_count"`
	// TOOD
	// - cost
}

type AccuracyResult struct {
	Accuracy     float64
	AccuracyRoll float64
	Critical     bool
	CriticalRoll float64
	Pass         bool
	Target       Actor
}

type DamageResult struct {
	AccuracyResult
	Affinity          float64
	AffinityStage     int
	BaseAffinityStage int
	Damage            float64
	Random            float64
	Raw               float64
}

func (ac ActionConfig) GetBaseDamage(source, target Actor, useBaseStats bool) float64 {
	adp_ratio := ac.Stat.GetRatio(source, target, useBaseStats) * ac.Power
	level_mod := float64(source.Level*2)/5 + 2
	base := (adp_ratio*level_mod)/50 + 2
	return base
}

func (ac ActionConfig) GetAccuracy(source, target Actor, useBaseStats bool) float64 {
	if ac.Accuracy == nil {
		return 1.0
	}

	ratio := Accuracy.GetRatio(source, target, useBaseStats)
	return ratio * *ac.Accuracy
}

func (ac ActionConfig) GetAccuracyResult(source, target Actor) AccuracyResult {
	accuracy_roll := rand.Float64()
	critical_roll := rand.Float64()
	critical_stage := int(ac.CritChance) + source.Stages[CriticalChance]
	critical_chance := GetCriticalChance(critical_stage)
	critical := critical_chance > critical_roll

	accuracy := ac.GetAccuracy(source, target, critical)
	success := ac.Accuracy == nil || accuracy > accuracy_roll

	if !success {
		critical = false
	}

	return AccuracyResult{
		Pass:         success,
		Accuracy:     accuracy,
		AccuracyRoll: accuracy_roll,
		Critical:     critical,
		CriticalRoll: critical_roll,
		Target:       target,
	}
}

func (ar AccuracyResult) Success() bool {
	return ar.Pass && !ar.Target.IsProtected
}

func (ac ActionConfig) GetDamageResult(source, target Actor, targets []Actor) DamageResult {
	accuracy := ac.GetAccuracyResult(source, target)
	affinity, total_stage, base_stage := ac.Affinity.GetAffinityModifier(source, target)
	base := ac.GetBaseDamage(source, target, accuracy.Critical)
	raw := base * affinity

	if accuracy.Critical {
		raw = raw * ac.CritModifier
	}

	if len(targets) > 1 {
		raw = raw * 0.75
	}

	if !accuracy.Success() {
		raw = 0.0
	}

	random := rand.Float64()*(1.05-0.8) + 0.8
	damage := raw * random
	if damage > target.GetRemainingHealth() {
		damage = target.GetRemainingHealth()
	}

	return DamageResult{
		AccuracyResult:    accuracy,
		Affinity:          affinity,
		AffinityStage:     total_stage,
		BaseAffinityStage: base_stage,
		Raw:               raw,
		Random:            random,
		Damage:            damage,
	}
}

func (dr DamageResult) Success() bool {
	return dr.AccuracyResult.Success() && dr.Damage > 0
}
