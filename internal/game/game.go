package game

import (
	"cmp"
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"slices"

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
	applied_modifiers   map[uuid.UUID]map[uuid.UUID]int
	modifier_immunities map[uuid.UUID]struct{}
	modifiers           []Modifier
}

func (gm *gamemeta) apply(modifierID uuid.UUID, actorID uuid.UUID) {
	_, ok := gm.applied_modifiers[modifierID]
	if !ok {
		gm.applied_modifiers[modifierID] = map[uuid.UUID]int{}
	}

	old, ok := gm.applied_modifiers[modifierID][actorID]
	if ok {
		gm.applied_modifiers[modifierID][actorID] = old + 1
	} else {
		gm.applied_modifiers[modifierID][actorID] = 1
	}
}

type GameStatus string
type GamePhase string

const (
	GameStatusRunning GameStatus = "running"
	GameStatusIdle    GameStatus = "idle"
	GameStatusWaiting GameStatus = "waiting"
)

const (
	PhaseInit    GamePhase = "init"
	PhaseStart   GamePhase = "start"
	PhaseMain    GamePhase = "main"
	PhaseEnd     GamePhase = "end"
	PhaseCleanup GamePhase = "cleanup"
)

type Game struct {
	state    State
	resolved State

	gamestate gamestate
	meta      gamemeta

	Phase  GamePhase
	Status GameStatus
	Turn   int
	Logs   []Bindable[Log]
}

func NewGame() Game {
	var state = State{
		Actors:       []Actor{},
		Players:      []Player{},
		Transactions: Queue[Transaction]{},
		Modifiers:    []Modifier{},
		Commands:     []Command{},
		Triggers:     []Command{},
		Prompts:      []Command{},
	}
	var system_modifiers = []Modifier{
		// map stat stages
		EffectActorsAll(EffectPriorityMapStages, func(g Game, a Actor, ctx Context) Actor {
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
		EffectActorsAll(EffectPriorityMapBaseStats, func(g Game, a Actor, ctx Context) Actor {
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
			applied_modifiers: map[uuid.UUID]map[uuid.UUID]int{},
		},
		Logs:   []Bindable[Log]{},
		Phase:  PhaseInit,
		Status: GameStatusIdle,
		Turn:   0,
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
func (g *Game) PushLog(log Bindable[Log]) {
	fmt.Println(log.Payload.Resolve())
	g.Logs = append(g.Logs, log)
}
func (g *Game) AddModifierImmunityTag(tag uuid.UUID) {
	g.meta.modifier_immunities[tag] = struct{}{}
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
func (g *Game) AppliedModifiers(actor_id uuid.UUID) map[uuid.UUID]int {
	effect_ids := map[uuid.UUID]int{}
	for modifier_id, actors := range g.meta.applied_modifiers {
		count, ok := actors[actor_id]
		if ok {
			effect_ids[modifier_id] = count
		}
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

	for _, a := range g.state.Actors {
		if a.IsActive() && a.IsAlive {
			modifiers = append(modifiers, a.GetModifiers()...)
		}
	}

	slices.SortStableFunc(modifiers, func(a, b Modifier) int {
		return a.Priority - b.Priority
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
func (g *Game) GetPlayer(id uuid.UUID) (Player, bool) {
	return g.State().FindPlayerByID(id)
}
func (g *Game) GetActor(id uuid.UUID) (Actor, bool) {
	return g.State().FindActorByID(id)
}
func (g *Game) FindActors(where Filter[Actor], context Context) []Actor {
	return g.State().FindActorsWhere(*g, where, context)
}
func (g *Game) GetActionableActors() []Actor {
	return g.FindActors(CombineFilters(
		ActiveActors,
		AliveActors,
	), NewContext())
}
func (g *Game) IsReadyToRun() bool {
	return len(g.State().Commands) == len(g.GetActionableActors())
}

func (g *Game) resolve() {
	g.gamestate = resolving
	g.resolved = g.state.Clone()
	g.meta.applied_modifiers = map[uuid.UUID]map[uuid.UUID]int{}
	g.meta.modifier_immunities = map[uuid.UUID]struct{}{}

	modifiers := g.GetModifiers()
	g.meta.modifiers = modifiers
	for _, mod := range modifiers {
		mod.Resolve(g)
	}

	g.gamestate = resolved
}

// mutations
func (g *Game) AddPlayers(players ...Player) {
	g.mutate(func(s *State) {
		s.Players = append(s.Players, players...)
	})
}
func (g *Game) AddActor(actor Actor) {
	g.mutate(func(s *State) {
		s.Actors = append(s.Actors, actor)
	})
}
func (g *Game) AddModifiers(modifiers ...Modifier) {
	g.mutate(func(s *State) {
		for _, mod := range modifiers {
			success := true
			if mod.Payload.Check != nil {
				success = mod.Payload.Check(*g, mod.Context)
			}

			if success {
				s.Modifiers = append(s.Modifiers, mod)
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
func (g *Game) PushCommand(command Command) {
	g.mutate(func(s *State) {
		s.Commands = append(s.Commands, command)
	})
}
func (g *Game) DeleteCommandWhere(where func(Command) bool) {
	g.mutate(func(s *State) {
		s.Commands = slices.DeleteFunc(s.Commands, where)
	})
}
func (g *Game) SortCommands() {
	g.mutate(func(s *State) {
		rand.Shuffle(len(s.Commands), func(i, j int) {
			s.Commands[i], s.Commands[j] = s.Commands[j], s.Commands[i]
		})

		slices.SortStableFunc(s.Commands, func(a, b Command) int {
			if byPriority := cmp.Compare(b.Priority, a.Priority); byPriority != 0 {
				return byPriority
			}

			aSource, aOK := g.GetSource(a.Context)
			bSource, bOK := g.GetSource(b.Context)
			if !aOK && !bOK {
				return 0
			}
			if !aOK {
				return 1
			}
			if !bOK {
				return -1
			}

			return cmp.Compare(bSource.Stats[Speed], aSource.Stats[Speed])
		})
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
func (g *Game) SetPosition(actor_id uuid.UUID, position_id uuid.UUID) {
	actor, ok := g.GetActor(actor_id)
	if !ok {
		return
	}

	var evicted_id uuid.UUID
	g.mutate(func(s *State) {
		updated := false

		s.UpdatePlayer(actor.PlayerID, func(player Player) Player {
			updated = true

			if position_id != uuid.Nil {
				current_id, ok := player.Positions[position_id]
				if !ok {
					updated = false
					return player
				}

				if current_id != actor_id {
					evicted_id = current_id
				}

				player.Positions[position_id] = actor_id
			}

			if actor.IsActive() && player.Positions[actor.PositionID] == actor_id {
				player.Positions[actor.PositionID] = uuid.Nil
			}

			return player
		})

		if !updated {
			return
		}

		trigger_context := MakeContextFor(actor)
		if position_id == uuid.Nil {
			log := NewLog("$actor$ left the battle.", map[string]string{
				"$actor$": actor.Name,
			})
			g.PushLog(log.Bind(trigger_context))
			g.On(OnActorEnter, trigger_context)
		} else {
			log := NewLog("$actor$ joined the battle.", map[string]string{
				"$actor$": actor.Name,
			})
			g.PushLog(log.Bind(trigger_context))
			g.On(OnActorLeave, trigger_context)
		}

		s.UpdateActor(actor_id, func(a Actor) Actor {
			a.PositionID = position_id
			return a
		})
	})

	if evicted_id != uuid.Nil {
		g.SetPosition(evicted_id, uuid.Nil)
	}
}
func (g *Game) DamageTargets(context Context, damage float64) {
	for _, target := range g.GetTargets(context) {
		g.MutateActor(target.ID, func(a Actor) Actor {
			resolved, ok := g.GetActor(target.ID)
			if !ok || !resolved.IsAlive {
				return a
			}

			a.ApplyDamage(damage, resolved)
			log_ctx := MakeContextFor(a, a)

			if damage > 0 {
				g.PushLog(NewLog(
					fmt.Sprintf("$target$ lost %d HP.", int(damage)),
					map[string]string{
						"$target$": a.Name,
					},
				).Bind(log_ctx))
			}

			if damage < 0 {
				g.PushLog(NewLog(
					fmt.Sprintf("$target$ healed %d HP.", int(-damage)),
					map[string]string{
						"$target$": a.Name,
					},
				).Bind(log_ctx))
			}

			if !a.IsAlive && resolved.IsAlive {
				g.PushLog(NewLog(
					"$target$ died.",
					map[string]string{
						"$target$": a.Name,
					},
				).Bind(log_ctx))
			}

			return a
		})
	}
}
func (g *Game) IncrementActorTurns() {
	for _, actor := range g.State().Actors {
		g.MutateActor(actor.ID, func(a Actor) Actor {
			a.IncrementTurns()
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
	g.Turn++
	g.IncrementActorTurns()
	g.Phase = PhaseMain
}
func (g *Game) EndPhase() {
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

	g.PushTransactions(trig.ResolveTrigger(g))
}
func (g *Game) NextCommand() {
	g.SortCommands()
	cmd, err := g.state.Commands.Dequeue()
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

type GameJSON struct {
	ActiveContext *Context        `json:"active_context"`
	Actors        []actorJSON     `json:"actors"`
	Logs          []Bindable[Log] `json:"logs"`
	Modifiers     []Modifier      `json:"modifiers"`
	Phase         GamePhase       `json:"phase"`
	Players       []Player        `json:"players"`
	Status        GameStatus      `json:"status"`
	Turn          int             `json:"turn"`
}

func (g Game) ToJSON() GameJSON {
	state := g.State()
	actors := make([]actorJSON, len(state.Actors))
	for i, actor := range state.Actors {
		actors[i] = actor.ToJSON(g)
	}

	return GameJSON{
		ActiveContext: state.ActiveContext,
		Actors:        actors,
		Logs:          g.Logs,
		Modifiers:     g.meta.modifiers,
		Phase:         g.Phase,
		Players:       state.Players,
		Status:        g.Status,
		Turn:          g.Turn,
	}
}

// temp functions
func (g *Game) Flush() {
	for g.Next() {
		//	time.Sleep(time.Second / 5)
	}

	g.mutate(func(s *State) {
		s.ActiveContext = nil
	})
}

func (g Game) MarshalJSON() ([]byte, error) {
	return json.Marshal(g.ToJSON())
}
