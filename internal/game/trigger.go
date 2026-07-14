package game

import (
	"fmt"

	"github.com/google/uuid"
)

type TriggerOn string

const (
	OnActorEnter    TriggerOn = "on-actor-enter"
	OnActorLeave    TriggerOn = "on-actor-leave"
	OnActorMove     TriggerOn = "on-actor-move"
	OnDamageSend    TriggerOn = "on-damage-send"
	OnDamageRecieve TriggerOn = "on-damage-recieve"
	OnAttackSuccess TriggerOn = "on-attack-success"
	OnAttackFailure TriggerOn = "on-attack-failure"
	OnCriticalHit   TriggerOn = "on-critical-hit"
	OnMiss          TriggerOn = "on-miss"
	OnModifierAdd   TriggerOn = "on-modifier-add"
	OnItemConsume   TriggerOn = "on-item-consume"
	OnTurnEnd       TriggerOn = "on-turn-end"
)

type Trigger struct {
	Action
	On       TriggerOn
	Validate func(g *Game, t_context Context, m_context Context) bool
}

type TriggerCommand struct {
	Bindable[Trigger]
	ParentContext Context
	Priority      int
}

func (a Trigger) Bind(context Context) TriggerCommand {
	return a.BindWithParent(context, context)
}

func (a Trigger) BindWithParent(context Context, parent_context Context) TriggerCommand {
	bindable := bind(a, context)
	command := TriggerCommand{
		Bindable:      bindable,
		ParentContext: parent_context,
		Priority:      a.Config.Priority,
	}
	return command
}

func (c TriggerCommand) Resolve(g *Game) []Transaction {
	if c.Payload.Resolve == nil {
		return []Transaction{}
	}

	action_context := ActionContext{
		Action:       c.Payload.Action,
		Source:       c.GetParent(g),
		transactions: []Transaction{},
	}

	context := c.Context.Clone()
	if context.EffectID == uuid.Nil {
		context.EffectID = c.ParentContext.EffectID
	}
	if context.ModifierID == uuid.Nil {
		context.ModifierID = c.ParentContext.ModifierID
	}
	context.ParentID = c.ParentContext.ParentID
	if c.Payload.MapContext != nil {
		context = c.Payload.MapContext(g, context, action_context)
	}

	g.SetActiveContext(context)
	g.ResetLogDepth()

	log := NewLog(fmt.Sprintf("$source$'s $action$ trigger (%s)", c.Payload.On), map[string]string{
		"$source$": action_context.Source.Name,
		"$action$": c.getName(g),
	})
	log_context := context.Clone()
	log_context.SourceID = action_context.Source.ID
	log_context.PlayerID = action_context.Source.PlayerID
	tx := log.Bind(log_context)
	tx.Payload.Type = "trigger"
	g.PushLogDepth(tx, 0)

	g.IncLogDepth()
	return c.Payload.Resolve(g, context, action_context)
}

func (c TriggerCommand) GetParent(g *Game) Actor {
	parent, ok := g.GetParent(c.ParentContext)
	if ok {
		return parent
	}

	return g.GetSourceAction(c.ParentContext)
}

func (c TriggerCommand) getName(g *Game) string {
	for _, modifier := range g.GetModifiers() {
		same_effect := modifier.Context.EffectID == c.ParentContext.EffectID
		same_parent := modifier.Context.ParentID == c.ParentContext.ParentID
		same_source := modifier.Context.SourceID == c.ParentContext.SourceID
		if same_effect && same_parent && same_source && modifier.Payload.Name != "" {
			return modifier.Payload.Name
		}
	}

	return c.Payload.Config.Name
}
