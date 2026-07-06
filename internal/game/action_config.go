package game

import (
	"math/rand/v2"
)

type ActionConfig struct {
	Accuracy     *float64 `json:"accuracy"`
	Affinity     Affinity `json:"affinity"`
	Cooldown     int      `json:"cooldown"`
	CritStage    int      `json:"crit_stage"`
	CritChance   float64  `json:"crit_chance"`
	CritModifier float64  `json:"crit_modifier"`
	Description  string   `json:"description"`
	Hits         int      `json:"hits"`
	Lifesteal    float64  `json:"lifesteal"`
	Name         string   `json:"name"`
	Power        float64  `json:"power"`
	Priority     int      `json:"priority"`
	Range        *int     `json:"range"`
	Recoil       float64  `json:"recoil"`
	Stat         Stat     `json:"stat"`
	TargetCount  int      `json:"target_count"`
	// TOOD
	// - cost
}

/**
 * notes about AccuracyResults
 *  - actor.IsProtected is checked for every action that goes through accuracy checks
 *    - source actions do not check accuracy and as a result do not check protected status
 */
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
	BaseDamage        float64
	Multipliers       float64
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
	critical_stage := ac.CritStage + source.Stages[CriticalChance]
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

const MULTI_TARGET_MODIFIER = 0.75
const DAMAGE_RAND_MIN = 0.8
const DAMAGE_RAND_MAX = 1.05

func (ac ActionConfig) GetDamageResult(source, target Actor, context Context, random_roll float64, pending_damage float64) DamageResult {
	multipliers := 1.0
	accuracy := ac.GetAccuracyResult(source, target)
	affinity, total_stage, base_stage := ac.Affinity.GetAffinityModifier(source, target)
	base := ac.GetBaseDamage(source, target, accuracy.Critical)
	raw := base * affinity
	multipliers *= affinity

	if accuracy.Critical {
		raw = raw * ac.CritModifier * source.Stats[CriticalDamage]
		multipliers *= ac.CritModifier * source.Stats[CriticalDamage]
	}

	if context.GetTargetCount() > 1 {
		raw = raw * MULTI_TARGET_MODIFIER
		multipliers *= MULTI_TARGET_MODIFIER
	}

	if !accuracy.Success() {
		raw = 0.0
	}

	random := random_roll*(DAMAGE_RAND_MAX-DAMAGE_RAND_MIN) + DAMAGE_RAND_MIN
	damage := raw * random
	health := target.GetRemainingHealth() - pending_damage
	if health < 0 {
		health = 0
	}
	if damage > health {
		damage = health
	}

	return DamageResult{
		AccuracyResult:    accuracy,
		Affinity:          affinity,
		AffinityStage:     total_stage,
		BaseAffinityStage: base_stage,
		BaseDamage:        base,
		Multipliers:       multipliers,
		Raw:               raw,
		Random:            random,
		Damage:            damage,
	}
}

func (dr DamageResult) Success() bool {
	return dr.AccuracyResult.Success() && dr.Damage > 0
}
