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
	applied_modifiers   map[uuid.UUID]map[uuid.UUID]struct{}
	modifier_immunities map[uuid.UUID]struct{}
	modifiers           []Modifier
}

func (gm *gamemeta) apply(modifierID uuid.UUID, actorID uuid.UUID) {
	_, ok := gm.applied_modifiers[modifierID]
	if !ok {
		gm.applied_modifiers[modifierID] = map[uuid.UUID]struct{}{}
	}

	gm.applied_modifiers[modifierID][actorID] = struct{}{}
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
		Positions:    []Position{},
		Transactions: Queue[Transaction]{},
		Modifiers:    []Modifier{},
		Commands:     []Command{},
		Triggers:     []TriggerCommand{},
		Prompts:      []PromptCommand{},
	}
	var system_modifiers = []Modifier{
		// map stat stages
		EffectActorsAll(EffectPriorityMapStages, func(g Game, a Actor, ctx Context) Actor {
			a.mapStagedStats()
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
			applied_modifiers: map[uuid.UUID]map[uuid.UUID]struct{}{},
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
func (g *Game) SetActiveContext(context Context) {
	g.mutate(func(s *State) {
		cloned := context.Clone()
		s.ActiveContext = &cloned
	})
}
func (g *Game) PushLog(log Bindable[Log]) {
	// fmt.Println(log.Payload.Resolve())
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
func (g *Game) AppliedModifiers(actor_id uuid.UUID) map[uuid.UUID]struct{} {
	effect_ids := map[uuid.UUID]struct{}{}
	for modifier_id, actors := range g.meta.applied_modifiers {
		_, ok := actors[actor_id]
		if ok {
			effect_ids[modifier_id] = struct{}{}
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
	return g.State().GetPlayer(id)
}
func (g *Game) GetActor(id uuid.UUID) (Actor, bool) {
	return g.State().GetActor(id)
}
func (g *Game) FindActors(where Filter[Actor], context Context) []Actor {
	return g.State().FindActors(*g, where, context)
}
func (g *Game) GetActionableActors() []Actor {
	return g.FindActors(CombineFilters(
		ActiveActors,
		AliveActors,
		NonStunnedActors,
	), NewContext())
}

func (g *Game) IsReadyToRun() bool {
	return len(g.State().Commands) == len(g.GetActionableActors())
}
func (g *Game) PromptsReady() bool {
	for _, prompt := range g.state.Prompts {
		if !prompt.Ready {
			return false
		}
	}

	return true
}

func (g *Game) resolve() {
	g.gamestate = resolving
	g.resolved = g.state.Clone()
	g.meta.applied_modifiers = map[uuid.UUID]map[uuid.UUID]struct{}{}
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
		for _, p := range players {
			s.Positions = append(s.Positions, Position{
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
				g.On(OnModifierAdd, mod.Context)
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
func (g *Game) RemoveModifiers(where Filter[Modifier], ctx Context) {
	g.mutate(func(s *State) {
		s.Modifiers = slices.DeleteFunc(s.Modifiers, func(m Modifier) bool {
			return where(*g, m, ctx)
		})
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
func (g *Game) PushPromptCommand(command PromptCommand) {
	g.mutate(func(s *State) {
		s.Prompts = append(s.Prompts, command)
	})
}
func (g *Game) UpdatePromptCommand(context Context) {
	g.mutate(func(s *State) {
		for i, cmd := range s.Prompts {
			is_player := cmd.Context.PlayerID == context.PlayerID
			is_source := cmd.Context.SourceID == context.SourceID
			is_action := cmd.Context.ActionID == context.ActionID
			if is_player && is_source && is_action {
				s.Prompts[i].Context = context
				s.Prompts[i].Ready = cmd.Payload.ValidateContext(*g, context)
			}
		}
	})
}
func (g *Game) On(on TriggerOn, context Context) {
	triggers := []TriggerCommand{}
	for _, modifier := range g.GetModifiers() {
		for _, trigger := range modifier.Payload.Triggers {
			if trigger.On != on {
				continue
			}
			if trigger.Validate != nil && !trigger.Validate(*g, context, modifier.Context) {
				continue
			}

			triggers = append(triggers, trigger.BindWithParent(context, modifier.Context))
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
		evicted_id, updated = s.SetPosition(position_id, actor)

		if !updated {
			return
		}

		s.UpdateActor(actor_id, func(a Actor) Actor {
			a.SetPosition(position_id)
			return a
		})

		trigger_context := MakeContextFor(actor)
		if position_id == uuid.Nil {
			s.Commands = slices.DeleteFunc(s.Commands, func(cmd Command) bool {
				return cmd.Context.ParentID == actor.ID
			})
			s.Modifiers = slices.DeleteFunc(s.Modifiers, func(mod Modifier) bool {
				return mod.Context.ParentID == actor.ID
			})

			if actor.IsAlive {
				log := NewLog("$source$ left the battle.", map[string]string{
					"$source$": actor.Name,
				})
				g.PushLog(log.Bind(trigger_context))
			} else {
				log := NewLog("$source$ died.", map[string]string{
					"$source$": actor.Name,
				})
				g.PushLog(log.Bind(trigger_context))
			}
			g.On(OnActorLeave, trigger_context)
		}
		if actor.PositionID == uuid.Nil {
			log := NewLog("$source$ joined the battle.", map[string]string{
				"$source$": actor.Name,
			})
			g.PushLog(log.Bind(trigger_context))
			g.On(OnActorEnter, trigger_context)
		}
	})

	if evicted_id != uuid.Nil {
		g.SetPosition(evicted_id, actor.PositionID)
	}
}
func (g *Game) PushForwards(actor_id uuid.UUID) {
	g.moveActor(actor_id, -1)
}
func (g *Game) PushToFront(actor_id uuid.UUID) {
	for g.moveActor(actor_id, -1) {
	}
}
func (g *Game) PushBackwards(actor_id uuid.UUID) {
	g.moveActor(actor_id, 1)
}
func (g *Game) PushToBack(actor_id uuid.UUID) {
	for g.moveActor(actor_id, 1) {
	}
}
func (g *Game) moveActor(actor_id uuid.UUID, direction int) bool {
	actor, ok := g.GetActor(actor_id)
	if !ok {
		return false
	}
	position, ok := g.state.GetPositionByActorID(actor_id)
	if !ok {
		return false
	}

	next, ok := nextPositionByRank(actor.PlayerID, g.State().Positions, position.Rank, direction)
	if !ok {
		return false
	}

	g.SetPosition(actor.ID, next.ID)
	log_ctx := MakeContextFrom(actor)
	log_ctx.PositionIDs = []uuid.UUID{next.ID}
	if direction > 0 {
		log := NewLog("$source$ moved backwards.", map[string]string{
			"$source$": actor.Name,
		})
		g.PushLog(log.Bind(log_ctx))
	}
	if direction < 0 {
		log := NewLog("$source$ moved forwards.", map[string]string{
			"$source$": actor.Name,
		})
		g.PushLog(log.Bind(log_ctx))
	}

	g.On(OnActorMove, log_ctx)
	return true
}

func (g *Game) DamageTargets(context Context, damage float64) {
	for _, target := range g.GetTargets(context) {
		g.MutateActor(target.ID, func(a Actor) Actor {
			resolved, ok := g.GetActor(target.ID)
			if !ok || !resolved.IsAlive {
				return a
			}

			if damage > resolved.GetRemainingHealth() {
				damage = resolved.GetRemainingHealth()
			}
			if damage < -a.Wounds {
				damage = -a.Wounds
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

// validate
func (g *Game) Validate() bool {
	valid := true
	actors := g.State().Actors
	for _, actor := range actors {
		if actor.IsActive() && !actor.IsAlive {
			g.SetPosition(actor.ID, uuid.Nil)
			valid = false
		}
	}

	g.condensePositions()

	for _, player := range g.State().Players {
		open_positions := g.State().GetOpenPositionIDs(player.ID)
		alive_inactive := g.FindActors(CombineFilters(AliveActors, Allies, InactiveActors), MakeContextPlayer(player.ID))
		size := min(len(open_positions), len(alive_inactive))
		if size > 0 {
			action := SwitchIn(size)
			ctx := NewContext()
			ctx.PlayerID = player.ID
			ctx.ActionID = action.ID
			ctx.PositionIDs = open_positions
			g.PushPromptCommand((action).ToPrompt().Bind(ctx))
		}
	}

	return valid
}

func (g *Game) condensePositions() bool {
	moved := false
	for {
		actorID, ok := g.nextActorBehindGap()
		if !ok {
			return moved
		}

		if !g.moveActor(actorID, -1) {
			return moved
		}
		moved = true
	}
}

func (g *Game) nextActorBehindGap() (uuid.UUID, bool) {
	for _, player := range g.State().Players {
		positions := g.State().GetPositionsByPlayerID(player.ID)
		slices.SortStableFunc(positions, func(a, b Position) int {
			return cmp.Compare(a.Rank, b.Rank)
		})

		foundGap := false
		for _, position := range positions {
			if position.ActorID == uuid.Nil {
				foundGap = true
				continue
			}

			if foundGap {
				return position.ActorID, true
			}
		}
	}

	return uuid.Nil, false
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
	if g.Turn > 0 {
		g.IncrementActorTurns()
	}
	g.Turn++
	g.Phase = PhaseMain
	log := NewLog(fmt.Sprintf("Turn %d", g.Turn), map[string]string{})
	log.Type = "turn"
	g.PushLog(log.Bind(NewContext()))
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

	if len(g.state.Triggers) > 0 {
		g.NextTrigger()
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

	if len(g.state.Commands) > 0 {
		g.NextCommand()
		return true
	}

	return false
}

type GameJSON struct {
	ActiveContext *Context               `json:"active_context"`
	Actors        []actorJSON            `json:"actors"`
	Commands      []Bindable[actionJSON] `json:"commands"`
	Logs          []Bindable[Log]        `json:"logs"`
	Modifiers     []Modifier             `json:"modifiers"`
	Phase         GamePhase              `json:"phase"`
	Positions     []Position             `json:"positions"`
	PlayerID      uuid.UUID              `json:"player_ID"`
	Players       []Player               `json:"players"`
	Prompts       []Bindable[actionJSON] `json:"prompts"`
	Status        GameStatus             `json:"status"`
	Turn          int                    `json:"turn"`
}

func (g Game) ToJSON() GameJSON {
	state := g.State()
	actors := make([]actorJSON, len(state.Actors))
	for i, actor := range state.Actors {
		actors[i] = actor.ToJSON(g)
	}

	prompts := []Bindable[actionJSON]{}
	for _, prompt := range state.Prompts {
		if prompt.Ready {
			continue
		}

		prompts = append(prompts, bind(prompt.Payload.ToJSON(g, Actor{}), prompt.Context))
	}
	commands := []Bindable[actionJSON]{}
	for _, command := range state.Commands {
		commands = append(commands, bind(command.Payload.ToJSON(g, Actor{}), command.Context))
	}

	return GameJSON{
		ActiveContext: state.ActiveContext,
		Actors:        actors,
		Commands:      commands,
		Logs:          g.Logs,
		Modifiers:     g.meta.modifiers,
		Phase:         g.Phase,
		Positions:     slices.Clone(g.State().Positions),
		PlayerID:      uuid.Nil,
		Players:       state.Players,
		Prompts:       prompts,
		Status:        g.Status,
		Turn:          g.Turn,
	}
}

func (json *GameJSON) ForPlayer(player_ID uuid.UUID) {
	json.PlayerID = player_ID
	prompts := slices.Clone(json.Prompts)
	prompts = slices.DeleteFunc(prompts, func(p Bindable[actionJSON]) bool {
		return p.Context.PlayerID != player_ID
	})
	commands := slices.Clone(json.Commands)
	commands = slices.DeleteFunc(commands, func(p Bindable[actionJSON]) bool {
		return p.Context.PlayerID != player_ID
	})
	json.Prompts = prompts
	json.Commands = commands

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
