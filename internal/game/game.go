package game

import (
	"fmt"
	"slices"
	"time"

	"github.com/google/uuid"
)

type gamestate uint8

const (
	unresolved gamestate = 0
	resolving  gamestate = 1
	resolved   gamestate = 2
)

type ActionContext struct {
	Action       Action
	Source       Actor
	transactions []Transaction
}

func (ac *ActionContext) Push(transactions ...Transaction) {
	ac.transactions = append(ac.transactions, transactions...)
}
func (ac *ActionContext) Concat(transactions []Transaction) {
	ac.transactions = append(ac.transactions, transactions...)
}
func (ac *ActionContext) Done() []Transaction {
	return ac.transactions
}

type gamemeta struct {
	appliedEffects map[uuid.UUID]map[uuid.UUID]int
}

func (gm *gamemeta) Apply(actorID uuid.UUID, modifierID uuid.UUID) {
	_, ok := gm.appliedEffects[actorID]
	if !ok {
		gm.appliedEffects[actorID] = map[uuid.UUID]int{}
	}

	old, ok := gm.appliedEffects[actorID][modifierID]
	if ok {
		gm.appliedEffects[actorID][modifierID] = old + 1
	} else {
		gm.appliedEffects[actorID][modifierID] = 1
	}
}

type Game struct {
	state    State
	resolved State

	gamestate gamestate
	meta      gamemeta

	Logs []Bindable[Log]
}

func NewGame() Game {
	var state = State{
		Actors:       []Actor{},
		Players:      []Player{},
		Transactions: Queue[Transaction]{},
		Modifiers:    []Modifier{},
		Commands:     []Command{},
		Triggers:     []Command{},
	}
	var system_modifiers = []Modifier{
		// map stat stages
		EffectActorsAll(EffectPriorityMapStages, func(a Actor, ctx Context) Actor {
			builder := NewStageBuilder(a.Stages)
			stats := builder.ResolveAll(a.Stats)

			builder.Mod = 3
			accuracy := builder.Resolve(Accuracy, a.Stats[Accuracy])
			evasion := builder.Resolve(Evasion, a.Stats[Evasion])

			a.Stats = stats
			a.Stats[Accuracy] = accuracy
			a.Stats[Evasion] = evasion

			return a
		}).Bind(NewContext()),
		// map base stats
		EffectActorsAll(EffectPriorityMapBaseStats, func(a Actor, ctx Context) Actor {
			a.mapBaseStats()
			return a
		}).Bind(NewContext()),
	}

	state.Modifiers = append(state.Modifiers, system_modifiers...)

	return Game{
		state:     state,
		resolved:  state,
		gamestate: unresolved,
		meta: gamemeta{
			appliedEffects: map[uuid.UUID]map[uuid.UUID]int{},
		},
	}
}

// setters
func (g *Game) mutate(updater func(*State)) {
	if g.gamestate == resolving {
		fmt.Println("!!! Tried to mutate state inside of resolve()")
		return
	}

	updater(&g.state)
	g.gamestate = unresolved
}
func (g *Game) modify(updater func(*State)) {
	if g.gamestate != resolving {
		fmt.Println("!!! Tried to modify state outside of resolve()")
		return
	}

	updater(&g.resolved)
}
func (g *Game) PushLog(log Log, context Context) {
	fmt.Println(log.Resolve())
	g.Logs = append(g.Logs, log.Bind(context))
}

// getters
func (g *Game) Base() State {
	return g.state
}
func (g *Game) State() State {
	if g.gamestate == unresolved {
		g.resolve()
	}

	return g.resolved
}
func (g *Game) AppliedEffects(actor_id uuid.UUID) map[uuid.UUID]int {
	effect_ids, ok := g.meta.appliedEffects[actor_id]
	if !ok {
		return map[uuid.UUID]int{}
	}

	return effect_ids
}
func (g *Game) GetModifiers() []Modifier {
	modifiers := []Modifier{}
	for _, m := range g.state.Modifiers {
		if m.Payload.Ready() {
			modifiers = append(modifiers, m)
		}
	}

	slices.SortStableFunc(modifiers, func(a, b Modifier) int {
		return int(a.Priority - b.Priority)
	})

	return modifiers
}
func (g *Game) GetTriggers() []Trigger {
	triggers := []Trigger{}
	for _, mod := range g.GetModifiers() {
		triggers = append(triggers, mod.Payload.Triggers...)
	}

	return triggers
}
func (g *Game) GetActor(id uuid.UUID) (Actor, bool) {
	return g.State().FindActorByID(id)
}
func (g *Game) FindActors(where Filter[Actor], context Context) []Actor {
	return g.State().FindActorsWhere(where, context)
}

func (g *Game) resolve() {
	g.gamestate = resolving
	g.resolved = g.state.Clone()
	g.meta.appliedEffects = map[uuid.UUID]map[uuid.UUID]int{}

	modifiers := g.GetModifiers()
	for _, mod := range modifiers {
		mod.Resolve(g)
	}

	g.gamestate = resolved
}

// mutations
func (g *Game) AddActors(actors ...Actor) {
	g.mutate(func(s *State) {
		s.Actors = append(s.Actors, actors...)
	})
}
func (g *Game) AddModifiers(modifiers ...Modifier) {
	g.mutate(func(s *State) {
		s.Modifiers = append(s.Modifiers, modifiers...)
	})

	for _, mod := range modifiers {
		if mod.Payload.GetLog != nil {
			log, ok := mod.Payload.GetLog(*g, mod.Payload, mod.Context)
			if ok {
				g.PushLog(log, mod.Context)
			}
		}
	}

}
func (g *Game) PushCommand(command Command) {
	g.mutate(func(s *State) {
		s.Commands = append(s.Commands, command)
	})
}
func (g *Game) On(on TriggerOn, context Context) {
	triggers := []Command{}
	for _, modifier := range g.GetModifiers() {
		for _, trigger := range modifier.Payload.Triggers {
			if trigger.Validate(*g, context, modifier.Context) {
				triggers = append(triggers, trigger.Bind(context))
			}
		}
	}

	fmt.Println("TRIGGER:", on, len(triggers))

	g.mutate(func(s *State) {
		s.Triggers = append(s.Triggers, triggers...)
	})
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
func (g *Game) DamageTargets(context Context, damage float64, trigger bool) {
	for _, target := range g.GetTargets(context) {
		g.MutateActor(target.ID, func(a Actor) Actor {
			found, ok := g.GetActor(target.ID)
			if !ok {
				return a
			}

			if damage > found.GetRemainingHealth() {
				damage = found.GetRemainingHealth()
			}

			a.Damage = a.Damage + damage
			if a.Damage < 0 {
				a.Damage = 0
			}

			a.IsAlive = found.Stats[Health] > a.Damage

			if damage > 0 {
				g.PushLog(
					NewLog(
						fmt.Sprintf("$target$ lost %d HP.", int(damage)),
						map[string]string{
							"$target$": a.Name,
						},
					),
					MakeContextFor(a, a),
				)

				if trigger {
					trigger_context := context.CloneWithTarget(target)
					g.On(OnDamageRecieve, trigger_context)
				}
			}

			if !a.IsAlive && found.IsAlive {
				g.PushLog(
					NewLog(
						"$target$ died.",
						map[string]string{
							"$target$": a.Name,
						},
					),
					MakeContextFor(a, a),
				)
			}

			return a
		})
	}
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

// control
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

	g.PushTransactions(trig.ResolveTrigger(*g))
}
func (g *Game) NextCommand() {
	cmd, err := g.state.Commands.Dequeue()
	if err != nil {
		return
	}

	g.PushTransactions(cmd.Resolve(*g))
}

func (g *Game) Next() bool {
	if len(g.state.Transactions) > 0 {
		g.NextTransaction()
		return true
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

// temp functions
func (g *Game) Flush() {
	for g.Next() {
		time.Sleep(time.Second / 5)
	}
}
