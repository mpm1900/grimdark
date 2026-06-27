package instance

import (
	"grimdark/internal/game"
	"grimdark/internal/game/actions"
)

func TestGame(g *game.Game) {
	max := g.FindActors(func(g game.Game, a game.Actor, ctx game.Context) bool {
		return a.Name == "Max"
	}, game.NewContext())[0]
	katie := g.FindActors(func(g game.Game, a game.Actor, ctx game.Context) bool {
		return a.Name == "Katie"
	}, game.NewContext())[0]

	ctx0 := game.MakeContextFor(katie, max)
	ctx0.ActionID = actions.SwordsDance.ID
	ctx1 := game.MakeContextFor(katie, max)
	ctx1.ActionID = actions.Slash.ID
	ctx2 := game.MakeContextFor(max, katie)
	ctx2.ActionID = actions.Slash.ID

	if cmd0, ok := g.HydrateToCommand(ctx0); ok {
		g.PushCommand(cmd0)
	}
	if cmd1, ok := g.HydrateToCommand(ctx1); ok {
		g.PushCommand(cmd1)
	}
	if cmd2, ok := g.HydrateToCommand(ctx2); ok {
		g.PushCommand(cmd2)
	}
}
