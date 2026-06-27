package game

import "github.com/google/uuid"

type ActionResolver func(g *Game, ctx Context, this ActionContext) []Transaction
type ActionContextMapper func(g Game, ctx Context, this ActionContext) Context

type Action struct {
	ID       uuid.UUID
	Config   ActionConfig
	Disabled bool

	Resolve          ActionResolver
	ValidateContext  GameFilter
	TargetsPredicate Filter[Actor]
	MapContext       ActionContextMapper
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

	can_act := !source.IsStaggered && !source.IsStunned
	context_valid := a.ValidateContext == nil || a.ValidateContext(g, context)
	source_valid := source.IsActive() && source.IsAlive && can_act
	action_valid := !a.Disabled
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

func (c Command) ResolveTrigger(g *Game) []Transaction {
	if !c.Payload.CanResolve(*g, c.Context, nil) || c.Payload.Resolve == nil {
		return []Transaction{}
	}

	action_context := ActionContext{
		Action:       c.Payload,
		Source:       g.GetSourceAction(c.Context),
		transactions: []Transaction{},
	}

	context := c.Context
	if c.Payload.MapContext != nil {
		context = c.Payload.MapContext(*g, context, action_context)
	}
	return c.Payload.Resolve(g, context, action_context)
}
