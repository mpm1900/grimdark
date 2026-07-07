package game

import (
	"fmt"
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

func AddResultEffects(chance float64, effects ...Effect) AttackEffectResult {
	return func(g Game, context Context, this *ActionContext, result DamageResult) {
		if !Chance(chance * this.Source.Stats[EffectChance]) {
			return
		}

		// insulated actors are immune to the secondary effects of actions
		if result.Target.IsInsulated {
			this.Push(PushLog(NewLog(
				"$target$ was insulated from secondary effects.",
				TargetTerms(result.Target),
			)).Bind(context))
			return
		}

		// actors are immune to the secondary effects of actions they are immune too as well
		_, immune := result.Target.AffinityImmunities[this.Action.Config.Affinity]
		if immune {
			this.Push(PushLog(NewLog(
				"$target$ was immune to $aff$.",
				CombineTerms(
					ActionTerms(this.Action),
					TargetTerms(result.Target),
				),
			)).Bind(context))
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
	if result.Success() && this.Action.Config.Hits > 1 && this.Action.Config.StopOnMiss {
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
				"Super Effective!", ETerms(),
			)).Bind(context))
		}
		if result.BaseAffinityStage <= -2 {
			this.Push(PushLog(NewLog(
				"Not Very Effective...", ETerms(),
			)).Bind(context))
		}
		if result.Critical {
			this.Push(PushLog(NewLog(
				"Critical hit!", ETerms(),
			)).Bind(context))
		}
	}

	if !result.Success() {
		_, immune := result.Target.AffinityImmunities[this.Action.Config.Affinity]
		if immune {
			this.Push(PushLog(NewLog(
				"$target$ was immune to $aff$.",
				CombineTerms(
					ActionTerms(this.Action),
					TargetTerms(result.Target),
				),
			)).Bind(context))
		}
		if result.Target.IsProtected {
			this.Push(PushLog(NewLog(
				"$target$ was protected.",
				TargetTerms(result.Target),
			)).Bind(context))
		}
		if !result.AccuracyResult.Pass {
			this.Push(PushLog(NewLog(
				"$action$ missed $target$.",
				CombineTerms(
					ActionTerms(this.Action),
					TargetTerms(result.Target),
				),
			)).Bind(context))
		}
	}
}
func DamageSideEffects(g *Game, context Context, result DamageResult, this *ActionContext, config AttackConfig) {
	trigger_context := MakeContextFor(this.Source, result.Target)
	if result.Success() {
		g.On(OnDamageSend, trigger_context)
		g.On(OnDamageRecieve, trigger_context)
		if result.Critical {
			g.On(OnCriticalHit, trigger_context)
		}

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
		if !result.AccuracyResult.Pass {
			g.On(OnMiss, trigger_context)
		}
		config.OnFailureResult(*g, context, this, result)
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
func CtxTargetPreCollateral() func(g Game, ctx Context, this ActionContext) Context {
	return func(g Game, ctx Context, this ActionContext) Context {
		if len(ctx.PositionIDs) == 0 {
			return ctx
		}

		target_position, ok := g.GetPosition(ctx.PositionIDs[0])
		if !ok {
			return ctx
		}

		c := ctx.Clone()
		c.ClearTargets()
		positions := g.GetEnemyPositionsByRank(ctx.PlayerID, target_position.Rank, -1)
		for _, pos := range positions {
			c.PositionIDs = append(c.PositionIDs, pos.ID)
		}

		return c
	}
}

// switches
func Retreat() Action {
	return Action{
		ID: uuid.MustParse("019f0f7c-50e5-7153-bede-2e8a3ef3dd60"),
		Config: ActionConfig{
			Name:        "Retreat",
			Description: "User switches out for an ally.",
			TargetCount: 1,
		},
		Type:    ATSystem,
		Subtype: ASRetreat,
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

var si_ids map[int]uuid.UUID = map[int]uuid.UUID{
	1: uuid.MustParse("019f0f7c-9bf6-7bbe-8e88-e5fea98d0930"),
	2: uuid.MustParse("019f1016-2713-7bff-8435-e958d2f216f0"),
	3: uuid.MustParse("019f1809-423a-7f2f-9a94-96532636ab4b"),
}

func SwitchIn(n int) Action {
	noun := "actor"
	if n > 1 {
		noun = "actors"
	}
	return Action{
		ID: si_ids[n],
		Config: ActionConfig{
			Name:        "Switch In",
			Description: fmt.Sprintf("Switch %s into battle.", noun),
			TargetCount: n,
		},
		Type:    ATSystem,
		Subtype: ASSwap,
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
		ActiveCheck: func(source Actor) bool {
			return false
		},
	}
}

func Swap() Action {
	return Action{
		ID:      uuid.MustParse("019f20f2-0860-7149-a11e-fbc3df357824"),
		Type:    ATSystem,
		Subtype: ASSwap,
		Config: ActionConfig{
			Name:        "Swap",
			Description: "User switches places with target ally.",
			TargetCount: 1,
		},
		Resolve: func(g *Game, ctx Context, this ActionContext) []Transaction {
			targets := g.GetTargets(ctx)
			for _, target := range targets {
				this.Push(SwapPositions(this.Source, target).Bind(ctx))
			}

			return this.Done()
		},
		TargetsPredicate: CombineFilters(OtherAllies, ActiveActors, AliveActors),
		ValidateContext:  ContextTargetLength(1),
	}
}
func MoveForwards() Action {
	return Action{
		ID:      uuid.MustParse("019f3ab7-eca5-7a8f-a5a9-214159374007"),
		Type:    ATSystem,
		Subtype: ASForward,
		Config: ActionConfig{
			Name:        "Move Forwards",
			Description: "User move forwards, displacing other actors back.",
			TargetCount: 0,
		},
		Resolve: func(g *Game, ctx Context, this ActionContext) []Transaction {
			this.Push(PushSourceForwards().Bind(ctx))

			return this.Done()
		},
		DisabledCheck: func(g Game, source Actor) bool {
			pos, ok := g.GetPosition(source.PositionID)
			if !ok {
				return true
			}

			return pos.Rank == 0
		},
		TargetsPredicate: NoneActors,
		ValidateContext:  ContextTargetLength(0),
	}
}
func MoveFront() Action {
	return Action{
		ID:      uuid.MustParse("019f3ab7-eca5-70fc-9483-f147b1dbdebb"),
		Type:    ATSystem,
		Subtype: ASFront,
		Config: ActionConfig{
			Name:        "Move to Front",
			Description: "User move the front of battle, displacing other actors back.",
			TargetCount: 0,
		},
		Resolve: func(g *Game, ctx Context, this ActionContext) []Transaction {
			this.Push(PushSourceToFront().Bind(ctx))

			return this.Done()
		},
		DisabledCheck: func(g Game, source Actor) bool {
			pos, ok := g.GetPosition(source.PositionID)
			if !ok {
				return true
			}

			return pos.Rank == 0
		},
		TargetsPredicate: NoneActors,
		ValidateContext:  ContextTargetLength(0),
	}
}
func MoveBackwards() Action {
	return Action{
		ID:      uuid.MustParse("019f3ab7-eca5-7ad4-a884-4ad97433a1e1"),
		Type:    ATSystem,
		Subtype: ASBack,
		Config: ActionConfig{
			Name:        "Move Backwards",
			Description: "User move backwards, displacing other actors forward.",
			TargetCount: 0,
		},
		Resolve: func(g *Game, ctx Context, this ActionContext) []Transaction {
			this.Push(PushSourceBackwards().Bind(ctx))

			return this.Done()
		},
		DisabledCheck: func(g Game, source Actor) bool {
			pos, ok := g.GetPosition(source.PositionID)
			if !ok {
				return true
			}

			return pos.Rank == 2
		},
		TargetsPredicate: NoneActors,
		ValidateContext:  ContextTargetLength(0),
	}
}

// global actions
var GLOBAL_ACTIONS = []Action{
	MoveForwards(),
	MoveFront(),
	MoveBackwards(),
	Retreat(),
	SwitchIn(1),
	SwitchIn(2),
	SwitchIn(3),
}
