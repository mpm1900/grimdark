package game

import (
	"math/rand/v2"

	"github.com/google/uuid"
)

func MakeAttack(config AttackConfig) ActionResolver {
	return func(g *Game, context Context, this ActionContext) []Transaction {
		targets := g.GetTargets(context)
		success := true
		hits := this.Action.Config.Hits

		if config.BeforeAttack != nil {
			config.BeforeAttack(g, context, &this)
		}

		for hit := range hits {
			if !success && this.Action.Config.StopOnMiss {
				break
			}

			for _, target := range targets {
				result := this.Action.Config.GetDamageResult(
					DamageConfig{
						Source: this.Source,
						Target: target,
						Context: context,
						RandomRoll: rand.Float64(),
						PendingDamage: this.PendingDamage(target.ID),
						UseBaseAccuracy: false,
					},
				)
				result.Print(this.Source)
				success = success && result.Success()
				dmg_ctx := MakeContextFor(this.Source, target)

				this.Push(DamageTargets(result.Damage, true).Bind(dmg_ctx))
				this.RecordDamage(target.ID, result.Damage)

				MultiHitLogs(result, context, &this, hit)
				PostDamageLogs(result, context, &this)
				DamageSideEffects(g, context, result, &this, config)
				g.MutateActor(this.Source.ID, func(a Actor) Actor {
					if this.Action.Weapon == nil {
						return a
					}

					slot := this.Action.Weapon.Slot
					a.Stacks[slot.String()] -= 1
					return a
				})
			}
		}

		for _, target := range targets {
			trigger_ctx := context.CloneWithTarget(target)
			if this.PendingDamage(target.ID) > 0 {
				g.On(OnAttackSuccess, trigger_ctx)
				if config.OnAttackSuccess != nil {
					config.OnAttackSuccess(g, trigger_ctx, &this)
				}
			} else {
				g.On(OnAttackFailure, trigger_ctx)
				if config.OnAttackFailure != nil {
					config.OnAttackFailure(g, trigger_ctx, &this)
				}
			}
		}

		if success && config.OnSuccess != nil {
			config.OnSuccess(g, context, &this)
		}
		if !success && config.OnFailure != nil {
			config.OnFailure(g, context, &this)
		}
		if config.OnFinally != nil {
			config.OnFinally(g, context, &this)
		}

		return this.Done()
	}
}

func AddGlobalEffects(config StatusConfig, chance float64, effects ...Effect) ActionResolver {
	return func(g *Game, ctx Context, this ActionContext) []Transaction {
		ctx.ParentID = uuid.Nil
		chance = chance * this.Source.Stats[EffectChance]
		if !Chance(chance) {
			if config.OnFailureResult != nil {
				config.OnFailureResult(g, ctx, &this, AccuracyResult{})
			}
			if config.OnFailure != nil {
				config.OnFailure(g, ctx, &this)
			}

			return this.Done()
		}

		modifiers := make([]Modifier, len(effects))
		for i, effect := range effects {
			modifiers[i] = effect.Bind(ctx)
		}

		this.Push(AddModifiers(modifiers...).Bind(NewContext()))
		if config.OnSuccessResult != nil {
			config.OnSuccessResult(g, ctx, &this, AccuracyResult{})
		}
		if config.OnSuccess != nil {
			config.OnSuccess(g, ctx, &this)
		}
		return this.Done()
	}
}

func AddSourceEffects(config StatusConfig, chance float64, effects ...Effect) ActionResolver {
	return func(g *Game, ctx Context, this ActionContext) []Transaction {
		chance = chance * this.Source.Stats[EffectChance]
		if !Chance(chance) {
			if config.OnFailureResult != nil {
				config.OnFailureResult(g, ctx, &this, AccuracyResult{})
			}
			if config.OnFailure != nil {
				config.OnFailure(g, ctx, &this)
			}

			return this.Done()
		}

		_, immune := this.Source.AffinityImmunities[this.Action.Config.Affinity]
		if immune {
			this.Push(PushLog(NewLog(
				"$target$ was immune to $aff$.",
				CombineTerms(
					ActionTerms(this.Action),
					TargetTerms(this.Source),
				),
			)).Bind(ctx))

			if config.OnFailureResult != nil {
				config.OnFailureResult(g, ctx, &this, AccuracyResult{})
			}
			if config.OnFailure != nil {
				config.OnFailure(g, ctx, &this)
			}

			return this.Done()
		}
		modifiers := make([]Modifier, len(effects))
		for i, effect := range effects {
			modifiers[i] = effect.Bind(ctx)
		}

		this.Push(AddModifiers(modifiers...).Bind(NewContext()))
		if config.OnSuccessResult != nil {
			config.OnSuccessResult(g, ctx, &this, AccuracyResult{})
		}
		if config.OnSuccess != nil {
			config.OnSuccess(g, ctx, &this)
		}
		return this.Done()
	}
}

func AddTargetsEffects(config StatusConfig, modifier_context Context, effects ...Effect) ActionResolver {
	return func(g *Game, ctx Context, this ActionContext) []Transaction {
		targets := g.GetTargets(ctx)
		success := false
		for _, target := range targets {
			result := this.Action.Config.GetAccuracyResult(this.Source, target, false)
			_, immune_affinity := target.AffinityImmunities[this.Action.Config.Affinity]
			if immune_affinity {
				this.Push(PushLog(NewLog(
					"$target$ was immune to $aff$.",
					CombineTerms(
						ActionTerms(this.Action),
						TargetTerms(result.Target),
					),
				)).Bind(ctx))
			}

			result_success := result.Success() && !immune_affinity
			success = success || result_success
			if result_success {
				modifiers := make([]Modifier, len(effects))
				target_ctx := MakeModifierContext(this.Source, target)
				CopyContext(&modifier_context, &target_ctx)
				for i, effect := range effects {
					modifiers[i] = effect.Bind(target_ctx)
				}
				this.Push(AddModifiers(modifiers...).Bind(NewContext()))
				if config.OnSuccessResult != nil {
					config.OnSuccessResult(g, ctx, &this, result)
				}
			}
			if !result_success && config.OnFailureResult != nil {
				config.OnFailureResult(g, ctx, &this, result)
			}
		}

		if success && config.OnSuccess != nil {
			config.OnSuccess(g, ctx, &this)
		}

		if !success && config.OnFailure != nil {
			config.OnFailure(g, ctx, &this)
		}

		return this.Done()
	}
}

func Struggle() Action {
	return Action{
		ID:   uuid.MustParse("019f5dab-239d-717a-9cc8-8a06f6461596"),
		Tags: []ActionTag{ATActor},
		Config: ActionConfig{
			Name:        "Struggle",
			Description: "User takes 1/4th of their max HP as recoil damage.",
			Stat:        Melee,
			Accuracy:    P(1.0),
			Power:       50,
			Hits:        1,
			TargetCount: 1,
		},
		LogTemplate: P("$source$ flails in a struggle."),
		Resolve: MakeAttack(AttackConfig{
			OnSuccessResult: func(g *Game, context Context, this *ActionContext, result DamageResult) {
				hp := this.Source.Stats[Health]
				recoil := hp * 0.25
				recoil_ctx := MakeContextFor(this.Source, this.Source)
				this.Push(DamageTargets(recoil, false).Bind(recoil_ctx))
			},
		}),
		ValidateContext:  ContextTargetLength(1),
		TargetsPredicate: CombineFilters(ActiveActors, NotSourceActor),
	}
}
