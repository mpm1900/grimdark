package game

import "fmt"

type TriggerOn string

const (
	OnActorEnter    TriggerOn = "on-actor-enter"
	OnActorLeave    TriggerOn = "on-actor-leave"
	OnDamageSend    TriggerOn = "on-damage-send"
	OnDamageRecieve TriggerOn = "on-damage-recieve"
	OnModifierAdd   TriggerOn = "on-modifier-add"
	OnTurnEnd       TriggerOn = "on-turn-end"
)

type Trigger struct {
	Action
	On       TriggerOn
	Validate func(g Game, t_context Context, m_context Context) bool
}

type TriggerCommand struct {
	Bindable[Trigger]
	Priority int
}

func (a Trigger) Bind(context Context) TriggerCommand {
	bindable := bind(a, context)
	command := TriggerCommand{
		Bindable: bindable,
		Priority: a.Config.Priority,
	}
	return command
}

func (c TriggerCommand) Resolve(g *Game) []Transaction {
	if c.Payload.Resolve == nil {
		return []Transaction{}
	}

	action_context := ActionContext{
		Action:       c.Payload.Action,
		Source:       g.GetSourceAction(c.Context),
		transactions: []Transaction{},
	}

	context := c.Context
	if c.Payload.MapContext != nil {
		context = c.Payload.MapContext(*g, context, action_context)
	}
	g.SetActiveContext(context)

	log := NewLog(fmt.Sprintf("$source$'s $action$ trigger (%s)", c.Payload.On), map[string]string{
		"$source$": action_context.Source.Name,
		"$action$": c.Payload.Config.Name,
	}).Bind(context)
	log.Payload.Type = "trigger"
	g.PushLog(log)

	return c.Payload.Resolve(g, context, action_context)
}
