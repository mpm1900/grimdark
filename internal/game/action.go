package game

import (
	"github.com/google/uuid"
)

type ActionResolver func(g *Game, ctx Context, this ActionContext) []Transaction

type Action struct {
	ID     uuid.UUID
	Config ActionConfig

	Resolve          ActionResolver
	ValidateContext  GameFilter
	TargetsPredicate Filter[Actor]
	MapContext       func(g Game, ctx Context, this ActionContext) Context

	IsActive      bool
	IsDisabled    bool
	DisabledCheck func(g Game, source Actor) bool // TODO
}

type actionJSON struct {
	ID         uuid.UUID    `json:"ID"`
	Config     ActionConfig `json:"config"`
	IsDisabled bool         `json:"is_disabled"`
}

func (a Action) Disabled(g Game, source Actor) bool {
	if !a.IsDisabled && a.DisabledCheck != nil {
		return a.DisabledCheck(g, source)
	}

	return a.IsDisabled
}

func (a Action) CanResolve(g Game, context Context, this *ActionContext) bool {
	source, ok := g.GetSource(context)
	if !ok {
		return false
	}

	if this != nil {
		if source.IsStunned {
			this.Push(
				PushLog(NewLog("$source$ was stunned.", map[string]string{
					"$source$": this.Source.Name,
				})).Bind(context),
			)
		}
		if source.IsStaggered {
			this.Push(
				PushLog(NewLog("$source$ was staggered.", map[string]string{
					"$source$": this.Source.Name,
				})).Bind(context),
			)
		}
	}

	context_valid := a.ValidateContext == nil || a.ValidateContext(g, context)
	source_valid := source.IsAlive && source.IsActive() && source.CanAct()
	action_valid := !a.Disabled(g, source)
	return action_valid && context_valid && source_valid
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
	g.mutate(func(s *State) {
		s.ActiveContext = &c.Context
	})

	action_context := ActionContext{
		Action:       c.Payload,
		Source:       g.GetSourceAction(c.Context),
		transactions: []Transaction{},
	}

	context := c.Context
	if c.Payload.MapContext != nil {
		context = c.Payload.MapContext(*g, context, action_context)
	}

	action_context.Push(
		PushLog(NewLog("$source$ used $action$.", map[string]string{
			"$source$": action_context.Source.Name,
			"$action$": c.Payload.Config.Name,
		})).Bind(context),
	)

	if c.Payload.Resolve == nil || !c.Payload.CanResolve(*g, context, &action_context) {
		action_context.Push(
			PushLog(NewLog("$action$ failed.", map[string]string{
				"$action$": c.Payload.Config.Name,
			})).Bind(context),
		)

		return action_context.transactions
	}

	return c.Payload.Resolve(g, context, action_context)
}

func (a Action) ToJSON(g Game, source Actor) actionJSON {
	json := actionJSON{
		ID:         a.ID,
		Config:     a.Config,
		IsDisabled: a.Disabled(g, source),
	}

	if json.Config.Accuracy != nil {
		acc := *json.Config.Accuracy * source.Stats[Accuracy]
		json.Config.Accuracy = &acc
	}

	json.Config.CritChance = GetCriticalChance(int(json.Config.CritChance) + source.Stages[CriticalChance])
	json.Config.CritModifier = json.Config.CritModifier * source.Stats[CriticalDamage]

	return json
}
