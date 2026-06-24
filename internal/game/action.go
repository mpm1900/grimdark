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
	Damage float64
	Random float64
}

func (ac ActionConfig) GetDamage(source, target Actor, useBaseStats bool) float64 {
	adp_ratio := ac.Stat.GetRatio(source, target, useBaseStats) * ac.Power
	level_mod := float64(source.Level*2)/5 + 2
	base := (adp_ratio*level_mod)/50 + 2
	affinity_mod := ac.Affinity.GetAffinityModifier(source, target)
	return (base * affinity_mod)
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
	damage := ac.GetDamage(source, target, accuracy.Critical)

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
		AccuracyResult: accuracy,
		Damage:         damage,
		Random:         rand.Float64()*(1.05-0.8) + 0.8,
	}
}

type Action struct {
	Config           ActionConfig
	Resolve          func(g Game, ctx Context, this ActionContext) []Transaction
	Validate         Filter[Game]
	TargetsPredicate Filter[Actor]
	MapContext       func(g Game, ctx Context, this ActionContext) Context
}

func (a Action) CanResolve(g Game, context Context) bool {
	source, ok := g.GetSource(context)
	if !ok {
		return false
	}

	return source.IsAlive && source.PositionID != nil
}

func (a Action) Bind(context Context) Command {
	bindable := bind(a, context)
	command := Command{
		Bindable: bindable,
		Priority: 0,
	}
	return command
}

type Command struct {
	Bindable[Action]
	Priority int
}

func (c Command) Resolve(g Game) []Transaction {
	action_context := ActionContext{
		Action:       c.Payload,
		Source:       g.GetSourceAction(c.Context),
		transactions: []Transaction{},
	}

	context := c.Context
	if c.Payload.MapContext != nil {
		context = c.Payload.MapContext(g, context, action_context)
	}

	action_context.Push(
		PushLog(NewLog("$source$ used $action$.", map[string]string{
			"$source$": action_context.Source.Name,
			"$action$": c.Payload.Config.Name,
		})).Bind(context),
	)

	if c.Payload.Resolve == nil || !c.Payload.CanResolve(g, context) {
		action_context.Push(
			PushLog(NewLog("$action$ failed.", map[string]string{
				"$action$": c.Payload.Config.Name,
			})).Bind(context),
		)

		return action_context.transactions
	}

	return c.Payload.Resolve(g, context, action_context)
}

func (c Command) ResolveTrigger(g Game) []Transaction {
	if !c.Payload.CanResolve(g, c.Context) || c.Payload.Resolve == nil {
		return []Transaction{}
	}

	action_context := ActionContext{
		Action:       c.Payload,
		Source:       g.GetSourceAction(c.Context),
		transactions: []Transaction{},
	}

	context := c.Context
	if c.Payload.MapContext != nil {
		context = c.Payload.MapContext(g, context, action_context)
	}
	return c.Payload.Resolve(g, context, action_context)
}

// helpers
func BasicAttack(g Game, context Context, this ActionContext) []Transaction {
	targets := g.GetTargets(context)
	for _, target := range targets {
		result := this.Action.Config.GetDamageResult(this.Source, target, targets)
		dmg_ctx := MakeContextFor(this.Source, target)
		damage := result.Damage * result.Random

		this.Push(DamageTargets(damage, true).Bind(dmg_ctx))
		CreatePostDamageEffects(result, context, &this)
	}

	return this.Done()
}

func CreatePostDamageEffects(result DamageResult, context Context, this *ActionContext) {
	if result.Critical {
		this.Push(PushLog(NewLog(
			"Critical hit!",
			map[string]string{},
		)).Bind(context))
	}

	if !result.Success {
		this.Push(PushLog(NewLog(
			"$source$'s $action$ missed $target$.",
			map[string]string{
				"$action$": this.Action.Config.Name,
				"$source$": this.Source.Name,
				"$target$": result.Target.Name,
			},
		)).Bind(context))
	}

	if result.Success && result.Target.IsProtected {
		this.Push(PushLog(NewLog(
			"$target$ was protected.",
			map[string]string{
				"$target$": result.Target.Name,
			},
		)).Bind(context))
	}
}
