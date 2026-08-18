package game

import (
	"github.com/google/uuid"
)

type ActionResolver func(g *Game, ctx Context, this ActionContext) []Transaction

type ActionTag string

const (
	ATSystem ActionTag = "system"
	ATActor  ActionTag = "actor"

	ATStruggle    ActionTag = "struggle"
	ATItem        ActionTag = "item"
	ATWeapon      ActionTag = "weapon"
	ATConditional ActionTag = "conditional"

	ATRetreat  ActionTag = "retreat"
	ATMovement ActionTag = "movement"
	ATSwap     ActionTag = "swap"
	ATBack     ActionTag = "back"
	ATForward  ActionTag = "forward"
	ATFront    ActionTag = "front"
)

type Action struct {
	ID               uuid.UUID
	Entity           Entity
	ActiveCheck      func(source Actor) bool
	Config           ActionConfig
	DisabledCheck    func(g *Game, source Actor) bool
	LogTemplate      *string
	MapContext       func(g *Game, ctx Context, this ActionContext) Context
	Resolve          ActionResolver
	Tags             []ActionTag
	TargetsPredicate Filter[Actor]
	Uncancelable     bool
	ValidateRuntime  GameFilter
	ValidateContext  GameFilter
	Weapon           *Weapon
}

type Command struct {
	Bindable[Action]
	Priority int
}

func (a Action) Disabled(g *Game, source Actor) bool {
	state, ok := source.ActionStates[a.ID]
	if ok && state.Cooldown > 0 {
		return true
	}
	if ok && a.Config.Uses != nil && state.Uses >= *a.Config.Uses {
		return true
	}

	if !state.IsDisabled && a.DisabledCheck != nil {
		return a.DisabledCheck(g, source)
	}

	return state.IsDisabled
}

func (a Action) CanResolve(g *Game, context Context, this *ActionContext) bool {
	source, ok := g.GetSource(context)
	if !ok {
		return false
	}
	runtime_valid := a.ValidateRuntime == nil || a.ValidateRuntime(g, context)
	source_valid := source.IsAlive && source.Active() && source.CanAct()
	action_valid := !a.Disabled(g, source)
	valid := action_valid && runtime_valid && source_valid

	if this != nil {
		if source.IsDisabled {
			this.Push(PushLog(NewLog("$source$ was stunned.", SourceTerms(source))).Bind(context))
		}
		if !action_valid {
			this.Push(PushLog(NewLog("$action$ was disabled.", ActionTerms(a))).Bind(context))
		} else if !valid {
			this.Push(PushLog(NewLog("$action$ failed.", ActionTerms(a))).Bind(context))
		}
	}

	return valid
}

func (a Action) Bind(context Context) Command {
	bindable := bind(a, context)
	command := Command{
		Bindable: bindable,
		Priority: a.Config.Priority,
	}
	return command
}

func (a Action) ToPrompt() Prompt {
	return Prompt{
		Action: a,
	}
}

func (c Command) Resolve(g *Game) []Transaction {
	action_context := ActionContext{
		Action:         c.Payload,
		Source:         g.GetSourceAction(c.Context),
		transactions:   []Transaction{},
		pending_damage: map[uuid.UUID]float64{},
	}

	state, ok := action_context.Source.ActionStates[c.Payload.ID]
	if ok {
		if state.BypassAccuracy {
			action_context.Action.Config.Accuracy = nil
		}
	}

	context := c.Context
	if c.Payload.MapContext != nil {
		context = c.Payload.MapContext(g, context, action_context)
	}

	g.SetActiveContext(context)
	g.ResetLogDepth()

	if c.Payload.LogTemplate == nil {
		log := NewLog("$source$ used $action$.", CommandTerms(action_context.Source, c))
		action_context.Push(PushLogDepth(log, 0).Bind(context))
	} else {
		log := NewLog(*c.Payload.LogTemplate, CommandTerms(action_context.Source, c))
		action_context.Push(PushLogDepth(log, 0).Bind(context))
	}

	if c.Payload.Resolve == nil || !c.Payload.CanResolve(g, context, &action_context) {
		return action_context.transactions
	}

	g.IncLogDepth()
	return c.Payload.Resolve(g, context, action_context)
}
