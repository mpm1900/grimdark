package game

import (
	"math/rand/v2"
	"strconv"
)

type AttackEffect func(g Game, context Context, this *ActionContext)
type AttackEffectResult func(g Game, context Context, this *ActionContext, result DamageResult)

type AttackConfig struct {
	OnSuccess       AttackEffect
	OnFailure       AttackEffect
	OnSuccessResult AttackEffectResult
	OnFailureResult AttackEffectResult
}

func BasicAttack(config AttackConfig) ActionResolver {
	return func(g *Game, context Context, this ActionContext) []Transaction {
		targets := g.GetTargets(context)
		success := false
		hits := this.Action.Config.Hits
		for hit := range hits {
			for _, target := range targets {
				result := this.Action.Config.GetDamageResult(this.Source, target, targets)
				dmg_ctx := MakeContextFor(this.Source, target)

				this.Push(DamageTargets(result.Damage).Bind(dmg_ctx))
				MultiHitEffects(result, context, &this, hit)
				PostDamageEffects(result, context, &this)

				success = success || result.Success()
				if result.Success() {
					trigger_context := MakeContextFor(this.Source, target)
					g.On(OnDamageRecieve, trigger_context)

					if config.OnSuccessResult != nil {
						config.OnSuccessResult(*g, context, &this, result)
					}
				}
				if !result.Success() && config.OnFailureResult != nil {
					config.OnFailureResult(*g, context, &this, result)
				}
			}
		}

		if success && config.OnSuccess != nil {
			config.OnSuccess(*g, context, &this)
		}

		if !success && config.OnFailure != nil {
			config.OnFailure(*g, context, &this)
		}

		return this.Done()
	}
}

func AddResultEffects(chance float64, effects ...Effect) AttackEffectResult {
	return func(g Game, context Context, this *ActionContext, result DamageResult) {
		if !Chance(chance) {
			return
		}

		modifiers := []Modifier{}
		for _, effect := range effects {
			modifiers = append(modifiers, effect.Bind(MakeModifierContext(this.Source, result.Target)))
		}

		mutation := AddModifiers(modifiers...)
		this.Push(mutation.Bind(NewContext()))
	}
}

func MultiHitEffects(result DamageResult, context Context, this *ActionContext, hit int) {
	if this.Action.Config.Hits > 1 {
		this.Push(PushLog(NewLog(
			"Hits $hit$ time(s)!",
			map[string]string{
				"$hit$": strconv.Itoa(hit + 1),
			},
		)).Bind(context))
	}
}
func PostDamageEffects(result DamageResult, context Context, this *ActionContext) {
	if result.Success() {
		if result.BaseAffinityStage >= 2 {
			this.Push(PushLog(NewLog(
				"Super Effective!",
				map[string]string{},
			)).Bind(context))
		}
		if result.Critical {
			this.Push(PushLog(NewLog(
				"Critical hit!",
				map[string]string{},
			)).Bind(context))
		}
	}

	if !result.Success() {
		if result.Target.IsProtected {
			this.Push(PushLog(NewLog(
				"$target$ was protected.",
				map[string]string{
					"$target$": result.Target.Name,
				},
			)).Bind(context))
		}
		if !result.AccuracyResult.Success {
			this.Push(PushLog(NewLog(
				"$action$ missed $target$.",
				map[string]string{
					"$action$": this.Action.Config.Name,
					"$source$": this.Source.Name,
					"$target$": result.Target.Name,
				},
			)).Bind(context))
		}
	}
}

func AddSourceEffects(chance float64, effects ...Effect) ActionResolver {
	return func(g *Game, ctx Context, this ActionContext) []Transaction {
		roll := rand.Float64()
		if chance <= roll {
			return this.Done()
		}

		modifiers := make([]Modifier, len(effects))
		for i, effect := range effects {
			modifiers[i] = effect.Bind(ctx)
		}
		this.Push(AddModifiers(modifiers...).Bind(NewContext()))
		return this.Done()
	}
}

// context mappers
func CtxToAllActiveTargets() ActionContextMapper {
	return func(g Game, ctx Context, this ActionContext) Context {
		c := ctx.CloneWithTargets(g.FindActors(ActiveActors, ctx))
		return c
	}
}
func CtxToAllOtherActiveTargets() ActionContextMapper {
	return func(g Game, ctx Context, this ActionContext) Context {
		c := ctx.CloneWithTargets(g.FindActors(ActiveActors, ctx))
		c.RemoveTarget(this.Source)
		return c
	}
}
