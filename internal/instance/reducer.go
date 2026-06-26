package instance

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

func findAction(g *game.Game, request Request) (game.Action, bool) {
	actor, ok := g.GetSource(request.Context)
	if !ok {
		return game.Action{}, false
	}

	action, ok := actor.GetActionByID(request.Context.ActionID)
	if !ok {
		return game.Action{}, false
	}

	return action, ok
}

func getTargets(instance *Instance, request Request) int {
	action, ok := findAction(&instance.Game, request)
	if !ok {
		instance.TargetIDsResponse(request.ClientID, request.Context, nil)
		return none
	}

	actors := instance.Game.State().Actors
	targetIDs := make([]uuid.UUID, 0, len(actors))
	for i := range actors {
		if action.TargetsPredicate(instance.Game, actors[i], request.Context) {
			targetIDs = append(targetIDs, actors[i].ID)
		}
	}

	instance.TargetIDsResponse(request.ClientID, request.Context, targetIDs)
	return none
}

func validateContext(instance *Instance, request Request) int {
	action, ok := findAction(&instance.Game, request)
	if !ok {
		instance.ValidateContextResponse(request.ClientID, request.Context, false)
		return none
	}

	valid := action.ValidateContext(instance.Game, request.Context)
	instance.ValidateContextResponse(request.ClientID, request.Context, valid)
	return none
}
func pushAction(instance *Instance, request Request) int {
	actor, ok := instance.Game.GetSource(request.Context)
	if !ok {
		return none
	}

	action, ok := actor.GetActionByID(request.Context.ActionID)
	if !ok {
		return none
	}

	//if action.State.Disabled {
	//	return none
	//}

	instance.Game.PushCommand(action.Bind(request.Context))

	if false {
		instance.RunGameActions()
	}

	return state
}
func cancelAction(instance *Instance, request Request) int {
	instance.Game.DeleteCommandWhere(func(cmd game.Command) bool {
		if cmd.Context.PlayerID == request.ClientID {
			return cmd.Payload.ID == request.Context.ActionID
		}

		return false
	})

	return state
}
func runGameActions(instance *Instance) int {
	if instance.Game.Status == game.GameStatusRunning {
		return none
	}

	instance.RunGameActions()
	return state
}

func Reducer(instance *Instance, request Request) int {
	switch request.Type {
	case GetTargets:
		return getTargets(instance, request)
	case ValidateContext:
		return validateContext(instance, request)
	case Reset:
		// instance.Game.Reset()
		return state
	case PushAction:
		return pushAction(instance, request)
	case CancelAction:
		return cancelAction(instance, request)
	case RunGameActions:
		return runGameActions(instance)
	default:
		return none
	}
}
