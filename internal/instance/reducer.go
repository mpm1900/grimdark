package instance

import (
	"grimdark/internal/game"
	"grimdark/internal/game/actors"
)

func postConnect(instance *Instance, request Request) {
	if request.TeamConfig == nil {
		return
	}

	actors.ApplyTeamConfig(instance.Game, request.ClientID, *request.TeamConfig)
	instance.PostConnectResponse(request.ClientID)
}
func lobbyReady(instance *Instance, request Request) {
	instance.Lobby.SetReady(request.ClientID)
	instance.BroadcastLobby()
	if instance.Lobby.Ready() {
		instance.Status = InstanceStatusRunning
		instance.BroadcastGameStart()
		instance.RunGameActions()
	}
}
func getTargets(instance *Instance, request Request) {
	action, ok := instance.Game.FindAction(request.Context)
	if !ok {
		instance.TargetIDsResponse(request.ClientID, request.Context)
		return
	}

	context := request.Context.Clone()
	context.ClearTargets()
	if action.TargetsPredicate == nil {
		instance.TargetIDsResponse(request.ClientID, context)
		return
	}

	actors := instance.Game.State().Actors
	for _, actor := range actors {
		if action.TargetsPredicate(instance.Game, actor, request.Context) && actor.Targetable() {
			context.AddTarget(actor)
		}
	}

	instance.TargetIDsResponse(request.ClientID, context)
}
func validateContext(instance *Instance, request Request) {
	action, ok := instance.Game.FindAction(request.Context)
	if !ok {
		instance.ValidateContextResponse(request.ClientID, request.Context, false)
		return
	}

	valid := action.ValidateContext(instance.Game, request.Context)
	instance.ValidateContextResponse(request.ClientID, request.Context, valid)
}

func turnReady(instance *Instance, request Request) {
	player, ok := instance.Game.GetPlayer(request.ClientID)
	if !ok {
		return
	}

	instance.Game.MutatePlayer(player.ID, func(p game.Player) game.Player {
		p.Ready = true
		return p
	})

	for _, player := range instance.Game.State().Players {
		if !player.Ready {
			instance.BroadcastGame()
			return
		}
	}

	if instance.Game.Status == game.GameStatusIdle {
		instance.RunGameActions()
	}
}
func pushAction(instance *Instance, request Request) {
	source, ok := instance.Game.GetSource(request.Context)
	if !ok {
		return
	}

	action, ok := source.GetActionByID(request.Context.ActionID)
	if !ok {
		return
	}

	instance.Game.PushCommand(source, action.Bind(request.Context))
	needed_actions := instance.Game.GetActionableActionsByPlayer(request.ClientID)
	state := instance.Game.State()
	commands := state.FindCommands(instance.Game, func(g *game.Game, c game.Command, ctx game.Context) bool {
		return c.Context.PlayerID == request.ClientID
	}, request.Context)
	if needed_actions == len(commands) {
		turnReady(instance, request)
		return
	}

	instance.BroadcastGame()
}
func cancelAction(instance *Instance, request Request) {
	instance.Game.DeleteCommandWhere(func(cmd game.Command) bool {
		if cmd.Context.PlayerID == request.ClientID {
			return cmd.Payload.ID == request.Context.ActionID
		}

		return false
	})

	instance.BroadcastGame()
}
func resolvePrompt(instance *Instance, request Request) {
	instance.Game.UpdatePromptCommand(request.Context)
	if instance.Game.PromptsReady() {
		instance.RunGameActions()
	}

	instance.BroadcastGame()
}

func Reducer(instance *Instance, request Request) {
	switch request.Type {
	case PostConnect:
		postConnect(instance, request)
		return
	case LobbyReady:
		lobbyReady(instance, request)
		return

	case GetTargets:
		getTargets(instance, request)
		return
	case ValidateContext:
		validateContext(instance, request)
		return

	case PushAction:
		pushAction(instance, request)
		return
	case CancelAction:
		cancelAction(instance, request)
		return
	case ResolvePrompt:
		resolvePrompt(instance, request)
		return
	case TurnReady:
		turnReady(instance, request)
		return
	default:
		return
	}
}
