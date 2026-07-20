package game

import "github.com/google/uuid"

func (a *Actor) ApplyDamage(damage float64, resolved Actor) {
	a.Stacks[Wounds] = a.Stacks[Wounds] + damage
	if a.Stacks[Wounds] < 0 {
		a.Stacks[Wounds] = 0
	}

	a.IsAlive = resolved.Stats[Health] > a.Stacks[Wounds]
}

func (a *Actor) NextTurn() {
	if a.Active() {
		a.Meta.ActiveTurns++
	} else {
		a.Meta.InactiveTurns++
	}

	for aid, state := range a.ActionStates {
		if state.Cooldown > 0 {
			a.UpdateActionState(aid, func(s ActionState) ActionState {
				s.Cooldown--
				return s
			})
		}
	}
}

func (a *Actor) SetActionCooldown(action_id uuid.UUID, cooldown int) {
	// +1 is for semantics, since cooldowns are decremented at turn end,
	// a one-turn cooldown needs a value of 2.
	a.UpdateActionState(action_id, func(s ActionState) ActionState {
		s.Cooldown = cooldown + 1
		s.Uses += 1
		return s
	})
}

func (a *Actor) SetPosition(position_id uuid.UUID) {
	if position_id == uuid.Nil {
		a.Meta.InactiveTurns = 0
		a.Meta.LastUsedActionID = uuid.Nil
	} else {
		a.Meta.Seen = true
	}
	if a.PositionID == uuid.Nil {
		a.Meta.ActiveTurns = 0
		a.Meta.ActiveHits = 0
	}
	a.PositionID = position_id
}

func (a *Actor) UpdateActionState(action_id uuid.UUID, updater func(ActionState) ActionState) {
	state, ok := a.ActionStates[action_id]
	if !ok {
		a.ActionStates[action_id] = updater(ActionState{})
	} else {
		a.ActionStates[action_id] = updater(state)
	}
}

func (g *Game) UpdatePromptCommand(context Context) {
	g.mutate(func(s *State) {
		for i, cmd := range s.Prompts {
			is_player := cmd.Context.PlayerID == context.PlayerID
			is_source := cmd.Context.SourceID == context.SourceID
			is_action := cmd.Context.ActionID == context.ActionID
			if is_player && is_source && is_action {
				s.Prompts[i].Context = context
				s.Prompts[i].Ready = cmd.Payload.ValidateContext(g, context)
			}
		}
	})
}
