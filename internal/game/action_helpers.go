package game

import (
	"math/rand/v2"
	"strconv"

	"github.com/google/uuid"
)

type ActionEffect func(g Game, context Context, this *ActionContext)
type AttackEffectResult func(g Game, context Context, this *ActionContext, result DamageResult)
type StatusEffectResult func(g Game, context Context, this *ActionContext, result AccuracyResult)

type AttackConfig struct {
	OnSuccess       ActionEffect
	OnFailure       ActionEffect
	OnSuccessResult AttackEffectResult
	OnFailureResult AttackEffectResult
}

type StatusConfig struct {
	OnSuccess       ActionEffect
	OnFailure       ActionEffect
	OnSuccessResult StatusEffectResult
	OnFailureResult StatusEffectResult
}

func BasicAttack(config AttackConfig) ActionResolver {
	return func(g *Game, context Context, this ActionContext) []Transaction {
		targets := g.GetTargets(context)
		success := true
		hits := this.Action.Config.Hits
		for hit := range hits {
			// on multi-hit attack, break after the first miss
			if !success {
				break
			}

			for _, target := range targets {
				result := this.Action.Config.GetDamageResult(this.Source, target, targets)
				dmg_ctx := MakeContextFor(this.Source, target)

				this.Push(DamageTargets(result.Damage).Bind(dmg_ctx))
				MultiHitLogs(result, context, &this, hit)
				PostDamageLogs(result, context, &this)

				success = success && result.Success()
				DamageSideEffects(g, context, result, &this, config)
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

func MultiHitLogs(result DamageResult, context Context, this *ActionContext, hit int) {
	if this.Action.Config.Hits > 1 {
		this.Push(PushLog(NewLog(
			"Hits $hit$ time(s)!",
			map[string]string{
				"$hit$": strconv.Itoa(hit + 1),
			},
		)).Bind(context))
	}
}
func PostDamageLogs(result DamageResult, context Context, this *ActionContext) {
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
		if !result.AccuracyResult.Pass {
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
func DamageSideEffects(g *Game, context Context, result DamageResult, this *ActionContext, config AttackConfig) {
	if result.Success() {
		trigger_context := MakeContextFor(this.Source, result.Target)
		g.On(OnDamageSend, trigger_context)
		g.On(OnDamageRecieve, trigger_context)

		if this.Action.Config.Recoil > 0 {
			recoil_ctx := MakeContextFor(this.Source, this.Source)
			this.Push(DamageTargets(result.Damage * this.Action.Config.Recoil).Bind(recoil_ctx))
		}

		if this.Action.Config.Lifesteal > 0 {
			lifesteal_ctx := MakeContextFor(this.Source, this.Source)
			this.Push(DamageTargets(result.Damage * this.Action.Config.Lifesteal * -1.0).Bind(lifesteal_ctx))
		}

		if config.OnSuccessResult != nil {
			config.OnSuccessResult(*g, context, this, result)
		}
	}
	if !result.Success() && config.OnFailureResult != nil {
		config.OnFailureResult(*g, context, this, result)
	}
}

// resolvers
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

func AddTargetsEffects(config StatusConfig, effects ...Effect) ActionResolver {
	return func(g *Game, ctx Context, this ActionContext) []Transaction {
		targets := g.GetTargets(ctx)
		success := false
		for _, target := range targets {
			result := this.Action.Config.GetAccuracyResult(this.Source, target)

			success = success || result.Success()
			if result.Success() {
				modifiers := make([]Modifier, len(effects))
				for i, effect := range effects {
					modifiers[i] = effect.Bind(ctx)
				}
				this.Push(AddModifiers(modifiers...).Bind(NewContext()))
				if config.OnSuccessResult != nil {
					config.OnSuccessResult(*g, ctx, &this, result)
				}
			}
			if !result.Success() && config.OnFailureResult != nil {
				config.OnFailureResult(*g, ctx, &this, result)
			}
		}

		if success && config.OnSuccess != nil {
			config.OnSuccess(*g, ctx, &this)
		}

		if !success && config.OnFailure != nil {
			config.OnFailure(*g, ctx, &this)
		}

		return this.Done()
	}
}

// context mappers
func CtxToAllActiveTargets() func(g Game, ctx Context, this ActionContext) Context {
	return func(g Game, ctx Context, this ActionContext) Context {
		c := ctx.CloneWithTargets(g.FindActors(ActiveActors, ctx))
		return c
	}
}
func CtxToAllOtherActiveTargets() func(g Game, ctx Context, this ActionContext) Context {
	return func(g Game, ctx Context, this ActionContext) Context {
		c := ctx.CloneWithTargets(g.FindActors(ActiveActors, ctx))
		c.RemoveTarget(this.Source)
		return c
	}
}

// switches
func SwitchWithSource() Action {
	return Action{
		ID: uuid.MustParse("019f0f7c-50e5-7153-bede-2e8a3ef3dd60"),
		Config: ActionConfig{
			Name: "Switch",
			Description: "User switches out for an ally.",
		},
		Resolve: func(g *Game, ctx Context, this ActionContext) []Transaction {
			if len(ctx.ActorIDs) != 1 {
				return this.Done()
			}

			this.Push(SetPositionSource(uuid.Nil).Bind(ctx))

			target_ctx := NewContext()
			target_ctx.SourceID = ctx.ActorIDs[0]
			this.Push(SetPositionSource(this.Source.PositionID).Bind(target_ctx))

			return this.Done()
		},
		TargetsPredicate: CombineFilters(Allies, InactiveActors, AliveActors),
		ValidateContext:  ContextTargetLength(1),
	}
}
func SwitchIn(n int) Action {
	return Action{
		ID: uuid.MustParse("019f0f7c-9bf6-7bbe-8e88-e5fea98d0930"),
		Config: ActionConfig{
			Name: "Switch",
		},
		Resolve: func(g *Game, ctx Context, this ActionContext) []Transaction {
			if len(ctx.ActorIDs) != len(ctx.PositionIDs) {
				return this.Done()
			}

			for i, position_id := range ctx.PositionIDs {
				actor_id := ctx.ActorIDs[i]
				switch_ctx := NewContext()
				switch_ctx.SourceID = actor_id
				this.Push(SetPositionSource(position_id).Bind(switch_ctx))
			}

			return this.Done()
		},
		TargetsPredicate: CombineFilters(Allies, InactiveActors, AliveActors),
		ValidateContext:  ContextTargetLength(n),
	}
}
