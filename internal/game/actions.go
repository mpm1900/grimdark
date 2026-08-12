package game

import (
	"github.com/google/uuid"
)

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

		if !success {
			this.Push(PushLog(NewLog(
				"$action$ failed.",
				CombineTerms(
					ActionTerms(this.Action),
				),
			)).Bind(ctx))

			if config.OnFailure != nil {
				config.OnFailure(g, ctx, &this)
			}
		}

		return this.Done()
	}
}

func Struggle() Action {
	id := uuid.MustParse("019f5dab-239d-717a-9cc8-8a06f6461596")
	return Action{
		ID:   id,
		Tags: []ActionTag{ATActor},
		Entity: MakeEntity(
			id,
			"Struggle",
			"User takes 1/4th of their max HP as recoil damage.",
		),
		Config: ActionConfig{
			Stat:        Melee,
			Accuracy:    P(1.0),
			Power:       50,
			TargetCount: 1,
		},
		LogTemplate: P("$source$ flails in a struggle."),
		Resolve: MakeAttack(AttackConfig{
			OnSuccessResult: func(g *Game, context Context, this *ActionContext, result DamageResult) {
				hp := this.Source.Stats[Health]
				recoil := hp * 0.25
				recoil_ctx := MakeContextFor(this.Source, this.Source)
				this.Push(DamageTargets(recoil, false, false).Bind(recoil_ctx))
			},
		}),
		ValidateContext:  ContextTargetLength(1),
		TargetsPredicate: CombineFilters(ActiveActors, NotSourceActor),
	}
}
