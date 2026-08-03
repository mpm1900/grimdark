package game

import "math/rand/v2"

type AttackConfig struct {
	BeforeAttack     ActionEffect
	OnSuccess        ActionEffect
	OnFailure        ActionEffect
	OnFinally        ActionEffect
	OnCriticalResult AttackEffectResult
	OnSuccessResult  AttackEffectResult
	OnFailureResult  AttackEffectResult
	OnFinallyResult  AttackEffectResult
	OnAttackSuccess  ActionEffect
	OnAttackFailure  ActionEffect
	HandleTarget     func(config AttackConfig, g *Game, context Context, this *ActionContext, hit int, target Actor) bool
}

func HandleAttackDamage(config AttackConfig, g *Game, context Context, this *ActionContext, hit int, target Actor) bool {
	result := this.Action.Config.GetDamageResult(
		DamageConfig{
			Source:          this.Source,
			Target:          target,
			Context:         context,
			RandomRoll:      rand.Float64(),
			PendingDamage:   this.PendingDamage(target.ID),
			UseBaseAccuracy: false,
		},
	)
	result.Print(this.Source)
	dmg_ctx := MakeContextFor(this.Source, target)

	this.Push(DamageTargets(result.Damage, true).Bind(dmg_ctx))
	this.RecordDamage(target.ID, result.Damage)

	MultiHitLogs(result, context, this, hit)
	PostDamageLogs(result, context, this)
	DamageSideEffects(g, context, result, this, config)
	g.MutateActor(this.Source.ID, func(a Actor) Actor {
		if this.Action.Weapon == nil {
			return a
		}

		slot := this.Action.Weapon.Slot
		a.Stacks[slot.String()] -= 1
		return a
	})

	return result.Success()
}

func HandleAttackCallbacks(config AttackConfig, g *Game, context Context, this *ActionContext, success bool, targets []Actor) {
	for _, target := range targets {
		trigger_ctx := context.CloneWithTarget(target)
		if this.PendingDamage(target.ID) > 0 {
			g.On(OnAttackSuccess, trigger_ctx)
			if config.OnAttackSuccess != nil {
				config.OnAttackSuccess(g, trigger_ctx, this)
			}
		} else {
			g.On(OnAttackFailure, trigger_ctx)
			if config.OnAttackFailure != nil {
				config.OnAttackFailure(g, trigger_ctx, this)
			}
		}
	}

	if success && config.OnSuccess != nil {
		config.OnSuccess(g, context, this)
	}
	if !success && config.OnFailure != nil {
		config.OnFailure(g, context, this)
	}
	if config.OnFinally != nil {
		config.OnFinally(g, context, this)
	}
}

func MakeAttack(config AttackConfig) ActionResolver {
	return func(g *Game, context Context, this ActionContext) []Transaction {
		targets := g.GetTargets(context)
		success := true
		hits := this.Action.Config.Repeats + 1

		if config.BeforeAttack != nil {
			config.BeforeAttack(g, context, &this)
		}

		for hit := range hits {
			if !success && this.Action.Config.StopOnMiss {
				break
			}

			for _, target := range targets {
				if config.HandleTarget == nil {
					config.HandleTarget = HandleAttackDamage
				}

				success = success && config.HandleTarget(
					config,
					g,
					context,
					&this,
					hit,
					target,
				)
			}
		}

		HandleAttackCallbacks(
			config,
			g,
			context,
			&this,
			success,
			targets,
		)

		return this.Done()
	}
}
