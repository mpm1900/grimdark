package game

import "github.com/google/uuid"

type Prompt struct {
	Action
}

type PromptCommand struct {
	Bindable[Prompt]
	Priority int
	Ready    bool
}

func (a Prompt) Bind(context Context) PromptCommand {
	bindable := bind(a, context)
	command := PromptCommand{
		Bindable: bindable,
		Priority: a.Config.Priority,
	}
	return command
}

func (c PromptCommand) Resolve(g *Game) []Transaction {
	if c.Payload.Resolve == nil {
		return []Transaction{}
	}

	action_context := ActionContext{
		Action:         c.Payload.Action,
		Source:         g.GetSourceAction(c.Context),
		transactions:   []Transaction{},
		pending_damage: map[uuid.UUID]float64{},
	}

	context := c.Context
	if c.Payload.MapContext != nil {
		context = c.Payload.MapContext(g, context, action_context)
	}
	g.SetActiveContext(context)

	return c.Payload.Resolve(g, context, action_context)
}
