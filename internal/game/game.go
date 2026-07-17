package game

import (
	"fmt"
	"slices"

	"github.com/google/uuid"
)

type gamestate uint8

const (
	unresolved gamestate = 0
	resolving  gamestate = 1
	resolved   gamestate = 2
)

type gamemeta struct {
	applied_modifiers map[uuid.UUID]map[uuid.UUID]struct{}
	log_depth         int
	modifiers         []Modifier
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
	GameStatusIdle    GameStatus = "idle"
	GameStatusRunning GameStatus = "running"
	GameStatusWaiting GameStatus = "waiting"
)

const (
	PhaseInit    GamePhase = "init"
	PhaseStart   GamePhase = "start"
	PhaseMain    GamePhase = "main"
	PhaseEnd     GamePhase = "end"
	PhaseCleanup GamePhase = "cleanup"
)

const maxLogCount = 30

type Game struct {
	state    State
	resolved State

	gamestate gamestate
	meta      gamemeta

	InstanceID uuid.UUID
	Phase      GamePhase
	Status     GameStatus
	Turn       int
	Logs       []Bindable[Log]
}

func NewGame(instanceID uuid.UUID) *Game {
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
		EffectActorsAll(EffectPriorityMapStages, func(g *Game, a Actor, ctx Context) Actor {
			a.mapStagedStats()
			return a
		}).Bind(NewContext()),
		// map base stats
		EffectActorsAll(EffectPriorityMapBaseStats, func(g *Game, a Actor, ctx Context) Actor {
			a.mapBaseStats()
			return a
		}).Bind(NewContext()),
	}

	state.Modifiers = append(state.Modifiers, system_modifiers...)

	return &Game{
		state:     state,
		resolved:  state,
		gamestate: unresolved,
		meta: gamemeta{
			applied_modifiers: map[uuid.UUID]map[uuid.UUID]struct{}{},
		},
		InstanceID: instanceID,
		Logs:       []Bindable[Log]{},
		Phase:      PhaseInit,
		Status:     GameStatusIdle,
		Turn:       0,
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
	g.Logs = append(g.Logs, log)
	if len(g.Logs) > maxLogCount {
		g.Logs = slices.Clone(g.Logs[len(g.Logs)-maxLogCount:])
	}
}
func (g *Game) PushLogDepth(log Bindable[Log], rank int) {
	log.Payload.Depth = rank
	g.PushLog(log)
}
func (g *Game) PushLogMeta(log Bindable[Log]) {
	g.PushLogDepth(log, g.meta.log_depth)
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
		if a.Active() && a.IsAlive {
			modifiers = append(modifiers, a.GetModifiers()...)
		}
	}

	slices.SortStableFunc(modifiers, func(a, b Modifier) int {
		return a.Priority - b.Priority
	})

	return modifiers
}
func (g *Game) GetModifier(id uuid.UUID) (Modifier, bool) {
	for _, mod := range g.GetModifiers() {
		if mod.ID == id {
			return mod, true
		}
	}
	return Modifier{}, false
}
func (g *Game) GetModifierByEffectID(effect_ID uuid.UUID) (Modifier, bool) {
	for _, mod := range g.GetModifiers() {
		if mod.Payload.ID == effect_ID {
			return mod, true
		}
	}
	return Modifier{}, false
}
func (g *Game) GetTriggers() []Trigger {
	triggers := []Trigger{}
	for _, mod := range g.GetModifiers() {
		triggers = append(triggers, mod.Payload.Triggers...)
	}

	return triggers
}
func (g *Game) GetPlayer(id uuid.UUID) (Player, bool) {
	state := g.State()
	return state.GetPlayer(id)
}
func (g *Game) GetActor(id uuid.UUID) (Actor, bool) {
	state := g.State()
	return state.GetActor(id)
}
func (g *Game) FindActors(where Filter[Actor], context Context) []Actor {
	state := g.State()
	return state.FindActors(g, where, context)
}
func (g *Game) GetActorsByPlayer(player_id uuid.UUID) []Actor {
	state := g.State()
	return state.FindActors(g, func(g *Game, a Actor, ctx Context) bool {
		return a.PlayerID == player_id
	}, NewContext())
}
func (g *Game) GetActionableActors() []Actor {
	return g.FindActors(CombineFilters(
		ActiveActors,
		AliveActors,
		NonStunnedActors,
	), NewContext())
}
func (g *Game) GetActionableActionsCount() int {
	count := 0
	for _, a := range g.GetActionableActors() {
		count += int(a.Stats[Actions])
	}

	return count
}
func (g *Game) GetActionableActionsByPlayer(player_ID uuid.UUID) int {
	count := 0
	for _, a := range g.GetActionableActors() {
		if a.PlayerID == player_ID {
			count += int(a.Stats[Actions])
		}
	}

	return count
}
func (g *Game) FindCommands(where Filter[Command], context Context) []Command {
	state := g.State()
	return state.FindCommands(g, where, context)
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

	modifiers := g.GetModifiers()
	g.meta.modifiers = modifiers
	for _, mod := range modifiers {
		mod.Resolve(g)
	}

	g.gamestate = resolved
}

func (g *Game) On(on TriggerOn, context Context) {
	triggers := []TriggerCommand{}
	for _, modifier := range g.GetModifiers() {
		for _, trigger := range modifier.Payload.Triggers {
			if trigger.On != on {
				continue
			}
			if trigger.Validate != nil && !trigger.Validate(g, context, modifier.Context) {
				continue
			}

			triggers = append(triggers, trigger.BindWithParent(context, modifier.Context))
		}
	}

	g.mutate(func(s *State) {
		s.Triggers = append(s.Triggers, triggers...)
	})
}

// validate
func (g *Game) Validate() bool {
	valid := true
	actors := g.State().Actors
	for _, actor := range actors {
		if actor.Active() && !actor.IsAlive {
			g.SetPosition(actor.ID, uuid.Nil)
			valid = false
		}
	}

	g.condensePositions()

	state := g.State()
	for _, player := range state.Players {
		open_positions := state.GetOpenPositionIDs(player.ID)
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

	if !valid {
		g.Status = GameStatusWaiting
	}
	return valid
}
