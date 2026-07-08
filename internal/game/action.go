package game

import (
	"github.com/google/uuid"
)

type ActionResolver func(g *Game, ctx Context, this ActionContext) []Transaction

type ActionTag string

const (
	ATSystem ActionTag = "system"
	ATActor  ActionTag = "actor"
	ATItem   ActionTag = "item"
	ATWeapon ActionTag = "weapon"

	ATRetreat ActionTag = "retreat"
	ATSwap    ActionTag = "swap"
	ATBack    ActionTag = "back"
	ATForward ActionTag = "forward"
	ATFront   ActionTag = "front"
)

type Action struct {
	ID          uuid.UUID
	Config      ActionConfig
	LogTemplate *string
	Tags        []ActionTag

	Resolve          ActionResolver
	Validate         GameFilter
	ValidateContext  GameFilter
	TargetsPredicate Filter[Actor]
	MapContext       func(g Game, ctx Context, this ActionContext) Context
	ActiveCheck      func(source Actor) bool
	DisabledCheck    func(g Game, source Actor) bool
}

type actionJSON struct {
	ID         uuid.UUID    `json:"ID"`
	Config     ActionConfig `json:"config"`
	Cooldown   int          `json:"cooldown"`
	IsDisabled bool         `json:"is_disabled"`
	Tags       []ActionTag  `json:"tags"`
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

	if c.Payload.LogTemplate == nil {
		log := NewLog("$source$ used $action$.", CommandTerms(action_context.Source, c))
		action_context.Push(PushLogDepth(log, 0).Bind(context))
	} else {
		log := NewLog(*c.Payload.LogTemplate, CommandTerms(action_context.Source, c))
		action_context.Push(PushLogDepth(log, 0).Bind(context))
	}

	if c.Payload.Resolve == nil || !c.Payload.CanResolve(*g, context, &action_context) {
		return action_context.transactions
	}

	g.IncLogDepth()
	return c.Payload.Resolve(g, context, action_context)
}

func (a Action) ToJSON(g Game, source Actor) actionJSON {
	state := source.ActionsState[a.ID]
	config := a.Config
	config.Cooldown = config.Cooldown + state.CooldownBonus
	config.Priority = config.Priority + state.PriorityBonus
	if config.Range != nil {
		config.Range = P(*config.Range + state.RangeBonus)
	}
	json := actionJSON{
		ID:         a.ID,
		Config:     config,
		Cooldown:   state.Cooldown,
		IsDisabled: a.Disabled(g, source),
		Tags:       a.Tags,
	}

	if json.Config.Accuracy != nil {
		acc := *json.Config.Accuracy * source.Stats[Accuracy]
		json.Config.Accuracy = &acc
	}

	json.Config.CritChance = GetCriticalChance(json.Config.CritStage + source.Stages[CriticalChance])
	json.Config.CritModifier = json.Config.CritModifier * source.Stats[CriticalDamage]

	return json
}
func (a Action) ToJSONStatic() actionJSON {
	return actionJSON{
		ID:         a.ID,
		Config:     a.Config,
		Cooldown:   0,
		IsDisabled: false,
		Tags:       a.Tags,
	}
}
