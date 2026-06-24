package game

func BasicAttack() ActionResolver {
	return func(g Game, context Context, this ActionContext) []Transaction {
		targets := g.GetTargets(context)
		for _, target := range targets {
			result := this.Action.Config.GetDamageResult(this.Source, target, targets)
			dmg_ctx := MakeContextFor(this.Source, target)
			damage := result.Damage * result.Random

			this.Push(DamageTargets(damage, true).Bind(dmg_ctx))
			PostDamageEffects(result, context, &this)
		}

		return this.Done()
	}
}
func PostDamageEffects(result DamageResult, context Context, this *ActionContext) {
	if result.Success {
		if result.Target.IsProtected {
			this.Push(PushLog(NewLog(
				"$target$ was protected.",
				map[string]string{
					"$target$": result.Target.Name,
				},
			)).Bind(context))
		} else {
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
	}

	if !result.Success {
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
func AddSourceEffects(effects ...Effect) ActionResolver {
	return func(g Game, ctx Context, this ActionContext) []Transaction {
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
