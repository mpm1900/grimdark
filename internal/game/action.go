package game

import (
	"github.com/google/uuid"
)

type ActionResolver func(g *Game, ctx Context, this ActionContext) []Transaction

type Action struct {
	ID     uuid.UUID
	Config ActionConfig

	Resolve          ActionResolver
	Validate         GameFilter
	ValidateContext  GameFilter
	TargetsPredicate Filter[Actor]
	MapContext       func(g Game, ctx Context, this ActionContext) Context

	IsActive      bool
	DisabledCheck func(g Game, source Actor) bool
}

type actionJSON struct {
	ID         uuid.UUID    `json:"ID"`
	Config     ActionConfig `json:"config"`
	Cooldown   int          `json:"cooldown"`
	IsDisabled bool         `json:"is_disabled"`
}

func (a Action) Disabled(g Game, source Actor) bool {
	state, ok := source.ActionsState[a.ID]
	if ok && state.Cooldown > 0 {
		return true
	}

	if !state.IsDisabled && a.DisabledCheck != nil {
		return a.DisabledCheck(g, source)
	}

	return false
}

func (a Action) CanResolve(g Game, context Context, this *ActionContext) bool {
	source, ok := g.GetSource(context)
	if !ok {
		return false
	}
	runtime_valid := a.Validate == nil || a.Validate(g, context)
	source_valid := source.IsAlive && source.IsActive() && source.CanAct()
	action_valid := !a.Disabled(g, source)
	valid := action_valid && runtime_valid && source_valid

	if this != nil {
		if source.IsStunned {
			this.Push(PushLog(NewLog("$source$ was stunned.", SourceTerms(source))).Bind(context))
		}
		if !action_valid {
			this.Push(PushLog(NewLog("$action$ as disabled.", ActionTerms(a))).Bind(context))
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

type Command struct {
	Bindable[Action]
	Priority int
}

func (c Command) Resolve(g *Game) []Transaction {
	action_context := ActionContext{
		Action:       c.Payload,
		Source:       g.GetSourceAction(c.Context),
		transactions: []Transaction{},
	}

	context := c.Context
	if c.Payload.MapContext != nil {
		context = c.Payload.MapContext(*g, context, action_context)
	}

	g.SetActiveContext(context)
	g.ResetLogDepth()

	log := NewLog("$source$ used $action$.", CommandTerms(action_context.Source, c))
	action_context.Push(PushLogDepth(log, 0).Bind(context))

	if c.Payload.Resolve == nil || !c.Payload.CanResolve(*g, context, &action_context) {
		return action_context.transactions
	}

	g.IncLogDepth()
	return c.Payload.Resolve(g, context, action_context)
}

func (a Action) ToJSON(g Game, source Actor) actionJSON {
	state := source.ActionsState[a.ID]
	json := actionJSON{
		ID:         a.ID,
		Config:     a.Config,
		Cooldown:   state.Cooldown,
		IsDisabled: a.Disabled(g, source),
	}

	if json.Config.Accuracy != nil {
		acc := *json.Config.Accuracy * source.Stats[Accuracy]
		json.Config.Accuracy = &acc
	}

	json.Config.CritChance = GetCriticalChance(json.Config.CritStage + source.Stages[CriticalChance])
	json.Config.CritModifier = json.Config.CritModifier * source.Stats[CriticalDamage]

	return json
}
