package game

import "fmt"

func (g *Game) NextPhase() {
	g.mutate(func(s *State) {
		s.ActiveContext = nil
	})

	switch g.Phase {
	case PhaseStart:
		g.Phase = PhaseMain
	case PhaseInit, PhaseMain:
		g.Phase = PhaseEnd
	case PhaseEnd:
		g.Phase = PhaseCleanup
	case PhaseCleanup:
		// Keep cleanup stable so callers can run end-of-turn bookkeeping once
		// without immediately wrapping back to main in the same loop tick.
	}
}
func (g *Game) NextTurn() {
	g.ActorNextTurnEffects()
	g.Turn++
	g.Phase = PhaseMain
	log := NewLog(fmt.Sprintf("Turn %d", g.Turn), map[string]string{})
	log.Type = "turn"
	g.PushLogMeta(log.Bind(NewContext()))
}
func (g *Game) EndPhase() {
	g.DecrementModifiers()
	if g.Turn > 0 {
		g.On(OnTurnEnd, NewContext())
	}
}
func (g *Game) NextTransaction() {
	tx, err := g.state.Transactions.Dequeue()
	if err != nil {
		return
	}

	tx.Resolve(g)
}
func (g *Game) NextTrigger() {
	trig, err := g.state.Triggers.Dequeue()
	if err != nil {
		return
	}

	g.PushTransactions(trig.Resolve(g))
}
func (g *Game) NextCommand() {
	g.SortCommands()
	cmd, err := g.state.Commands.Dequeue()
	if err != nil {
		return
	}

	g.PushTransactions(cmd.Resolve(g))
	g.SetLastUsedAction(cmd.Context.SourceID, cmd.Payload.ID)

	source, ok := g.GetSource(cmd.Context)
	if !ok {
		g.SetCooldown(cmd, ActionState{})
		return
	}

	state, ok := source.ActionStates[cmd.Payload.ID]
	if !ok {
		g.SetCooldown(cmd, ActionState{})
		return
	}

	g.SetCooldown(cmd, state)
}
func (g *Game) NextPrompt() {
	cmd, err := g.state.Prompts.Dequeue()
	if err != nil {
		return
	}

	g.PushTransactions(cmd.Resolve(g))
}

func (g *Game) Next() bool {
	if len(g.state.Transactions) > 0 {
		g.NextTransaction()
		return true
	}

	if len(g.state.Prompts) > 0 {
		if g.PromptsReady() {
			g.NextPrompt()
			return true
		}
		return false
	}

	if !g.Validate() {
		return false
	}

	if len(g.state.Triggers) > 0 {
		g.NextTrigger()
		return true
	}

	if len(g.state.Commands) > 0 {
		g.NextCommand()
		return true
	}

	return false
}
