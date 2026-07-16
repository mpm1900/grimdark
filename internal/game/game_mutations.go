package game

import (
	"cmp"
	"fmt"
	"math/rand/v2"
	"slices"

	"github.com/google/uuid"
)

func (g *Game) ActorNextTurnEffects() {
	state := g.State()
	for _, actor := range state.Actors {
		g.MutateActor(actor.ID, func(a Actor) Actor {
			a.NextTurn()
			return a
		})
	}
	for _, player := range state.Players {
		g.MutatePlayer(player.ID, func(p Player) Player {
			p.NextTurn()
			return p
		})
	}
}

func (g *Game) AddActors(actors ...Actor) {
	g.mutate(func(s *State) {
		s.Actors = append(s.Actors, actors...)
	})
}

func (g *Game) AddModifiers(modifiers ...Modifier) {
	g.mutate(func(s *State) {
		for _, mod := range modifiers {
			success := true
			if mod.Payload.Check != nil {
				success = mod.Payload.Check(g, mod.Context)
			}

			if success {
				s.Modifiers = append(s.Modifiers, mod)

				trigger_context := mod.Context.Clone()
				trigger_context.ModifierID = mod.ID
				g.On(OnModifierAdd, trigger_context)
				if mod.Payload.CheckSuccess != nil {
					mod.Payload.CheckSuccess(g, mod.Payload, mod.Context)
				}
			}

			if !success {
				if mod.Payload.CheckFailure != nil {
					mod.Payload.CheckFailure(g, mod.Payload, mod.Context)
				}
				continue
			}
		}
	})
}

func (g *Game) AddPlayers(players ...Player) {
	g.mutate(func(s *State) {
		s.Players = append(s.Players, players...)
		for _, p := range players {
			s.Positions = append(s.Positions,
				Position{
					ID:       uuid.New(),
					ActorID:  uuid.Nil,
					PlayerID: p.ID,
					Rank:     0,
				},
				Position{
					ID:       uuid.New(),
					ActorID:  uuid.Nil,
					PlayerID: p.ID,
					Rank:     1,
				},
				Position{
					ID:       uuid.New(),
					ActorID:  uuid.Nil,
					PlayerID: p.ID,
					Rank:     2,
				})
		}
	})
}

func (g *Game) DamageTargets(context Context, damage float64) {
	for _, target := range g.GetTargets(context) {
		g.MutateActor(target.ID, func(a Actor) Actor {
			resolved, ok := g.GetActor(target.ID)
			if !ok || !resolved.IsAlive {
				return a
			}

			target_damage := damage
			if target_damage > resolved.GetRemainingHealth() {
				target_damage = resolved.GetRemainingHealth()
			}
			if target_damage < -a.Stacks[Wounds] {
				target_damage = -a.Stacks[Wounds]
			}

			a.ApplyDamage(target_damage, resolved)
			log_ctx := MakeContextFor(a, a)

			if target_damage > 0 {
				ratio := target_damage / target.Stats[Health]

				g.PushLogMeta(NewLog(
					fmt.Sprintf("$target$ lost %d%% HP.", int(ratio*100)),
					TargetTerms(a),
				).Bind(log_ctx))
			}

			if target_damage < 0 {
				ratio := -target_damage / target.Stats[Health]
				g.PushLogMeta(NewLog(
					fmt.Sprintf("$target$ healed %d%% HP.", int(ratio*100)),
					TargetTerms(a),
				).Bind(log_ctx))
			}

			return a
		})
	}
}

func (g *Game) DecrementModifiers() {
	g.mutate(func(s *State) {
		for i, modifier := range s.Modifiers {
			if modifier.Payload.Delay != nil {
				*s.Modifiers[i].Payload.Delay--
			}
			if modifier.Payload.Duration != nil {
				*s.Modifiers[i].Payload.Duration--
			}
		}

		s.Modifiers = slices.DeleteFunc(s.Modifiers, func(modifier Modifier) bool {
			return modifier.Payload.Duration != nil && *modifier.Payload.Duration <= 0
		})
	})
}

func (g *Game) DeleteCommandWhere(where func(Command) bool) {
	g.mutate(func(s *State) {
		s.Commands = slices.DeleteFunc(s.Commands, where)
	})
}

func (g *Game) MutatePlayer(id uuid.UUID, updater func(Player) Player) {
	g.mutate(func(s *State) {
		s.UpdatePlayer(id, updater)
	})
}
func (g *Game) MutateActor(id uuid.UUID, updater func(Actor) Actor) {
	g.mutate(func(s *State) {
		s.UpdateActor(id, updater)
	})
}
func (g *Game) MutateActorWhere(where func(Actor) bool, updater func(Actor) Actor) {
	g.mutate(func(s *State) {
		s.UpdateActorWhere(where, updater)
	})
}

func (g *Game) PushCommand(source Actor, command Command) {
	found := g.FindCommands(func(g *Game, c Command, ctx Context) bool {
		return c.Context.SourceID == command.Context.SourceID
	}, command.Context)

	if len(found) >= int(source.Stats[Actions]) {
		fmt.Println("Actor already actions queued.")
		return
	}

	g.mutate(func(s *State) {
		s.Commands = append(s.Commands, command)
	})
}

func (g *Game) PushPromptCommand(command PromptCommand) {
	g.mutate(func(s *State) {
		s.Prompts = append(s.Prompts, command)
	})
	g.Status = GameStatusWaiting
}

func (g *Game) PushTransaction(transaction Transaction) {
	g.mutate(func(s *State) {
		s.Transactions = append(s.Transactions, transaction)
	})
}
func (g *Game) PushTransactions(transactions []Transaction) {
	g.mutate(func(s *State) {
		s.Transactions = append(s.Transactions, transactions...)
	})
}

func (g *Game) RemoveModifiers(where Filter[Modifier], ctx Context) {
	g.mutate(func(s *State) {
		s.Modifiers = slices.DeleteFunc(s.Modifiers, func(m Modifier) bool {
			return where(g, m, ctx)
		})
	})
}

func (g *Game) SetActiveContext(context Context) {
	g.mutate(func(s *State) {
		cloned := context.Clone()
		s.ActiveContext = &cloned
	})
}

func (g *Game) SetLastUsedAction(actor_id uuid.UUID, action_id uuid.UUID) {
	g.mutate(func(s *State) {
		s.UpdateActor(actor_id, func(a Actor) Actor {
			a.Meta.LastUsedActionID = action_id
			return a
		})
	})
}

func (g *Game) SetCooldown(cmd Command, state ActionState) {
	g.mutate(func(s *State) {
		s.UpdateActor(cmd.Context.SourceID, func(a Actor) Actor {
			cooldown := cmd.Payload.Config.Cooldown + state.CooldownBonus
			a.SetActionCooldown(cmd.Payload.ID, cooldown)
			return a
		})
	})
}

func (g *Game) SortCommands() {
	g.mutate(func(s *State) {
		rand.Shuffle(len(s.Commands), func(i, j int) {
			s.Commands[i], s.Commands[j] = s.Commands[j], s.Commands[i]
		})

		slices.SortStableFunc(s.Commands, func(a, b Command) int {
			a_source, a_ok := g.GetSource(a.Context)
			b_source, b_ok := g.GetSource(b.Context)
			if !a_ok && !b_ok {
				return 0
			}
			if !a_ok {
				return 1
			}
			if !b_ok {
				return -1
			}

			a_state := a_source.ActionStates[a.Payload.ID]
			b_state := b_source.ActionStates[b.Payload.ID]
			a_priority := a.Priority + a_state.PriorityBonus
			b_priority := b.Priority + b_state.PriorityBonus
			by_priority := cmp.Compare(b_priority, a_priority)
			if by_priority != 0 {
				return by_priority
			}

			return cmp.Compare(b_source.Stats[Speed], a_source.Stats[Speed])
		})
	})
}

// modifiers
func (g *Game) ModifyActor(id uuid.UUID, updater func(Actor) Actor) {
	g.modify(func(s *State) {
		s.UpdateActor(id, updater)
	})
}
func (g *Game) ModifyActorWhere(where func(Actor) bool, updater func(Actor) Actor) {
	g.modify(func(s *State) {
		s.UpdateActorWhere(where, updater)
	})
}
