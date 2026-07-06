package instance

import (
	"grimdark/internal/game"
	"grimdark/internal/game/actions"
	"grimdark/internal/game/effects"
	"grimdark/internal/game/items"
	"grimdark/internal/game/weapons"

	"github.com/google/uuid"
)

func SetupOpponent(g *game.Game) {
	monsters := g.FindActors(func(g game.Game, a game.Actor, ctx game.Context) bool {
		return a.Name == "Monster"
	}, game.NewContext())

	if len(monsters) == 0 {
		opp := game.NewPlayer(uuid.New())
		g.AddPlayers(opp)
		open_positions := g.State().GetOpenPositions(opp.ID)
		newMonster := func() game.Actor {
			opp_def := game.NewActorDef()
			opp_def.Name = "Monster"
			opp_def.Affinities = map[game.Affinity]struct{}{
				game.Poison: {},
			}
			return game.NewActor(opp.ID, opp_def)
		}
		for _, pos := range open_positions {
			mnstr := newMonster()
			g.AddActor(mnstr)
			g.SetPosition(mnstr.ID, pos.ID)
		}
	}
}

func SetupGame(g *game.Game, user game.User) {
	bypass := effects.StagesResetWhere(func(g game.Game, a game.Actor, ctx game.Context) bool {
		active_context := g.State().ActiveContext
		if active_context == nil {
			return false
		}

		if active_context.SourceID == ctx.ParentID {
			return active_context.HasTarget(a)
		}

		return false
	})
	bypass.Name = "bypass effect"

	player := game.NewPlayer(user.ID)
	player.User = user
	if len(g.State().Players) > 0 {
		player = g.Base().Players[0]
	}

	max_def := game.NewActorDef()
	max_def.Name = "Max"
	max_def.Affinities = map[game.Affinity]struct{}{
		game.Fire: {},
	}
	max := game.NewActor(player.ID, max_def)
	max.Effects = []game.Effect{bypass, effects.Intimidate()}
	max.Weapon = &weapons.SlashSword

	katie_def := game.NewActorDef()
	katie_def.Name = "Katie"
	katie_def.Affinities = map[game.Affinity]struct{}{
		game.Cryo:   {},
		game.Arcane: {},
	}

	katie := game.NewActor(player.ID, katie_def)
	katie.AuxStats[game.Speed] = 10
	katie.Actions = []game.Action{
		actions.SwordsDance,
	}
	katie.Weapon = &weapons.SlashSword
	katie.Item = game.P(items.TestItem())
	katie.AffinityImmunities = map[game.Affinity]float64{
		game.Kinetic: 0,
	}

	gabe_def := game.NewActorDef()
	gabe_def.Name = "gabe"
	gabe_def.Affinities = map[game.Affinity]struct{}{
		game.Cryo:   {},
		game.Arcane: {},
	}
	gabe := game.NewActor(player.ID, gabe_def)
	gabe.Actions = []game.Action{
		actions.SwordsDance,
	}
	gabe.Weapon = &weapons.SlashSword
	gabe.Item = game.P(items.TestItem())
	gabe.AffinityImmunities = map[game.Affinity]float64{
		game.Kinetic: 0,
	}

	if len(g.State().Players) == 0 {
		g.AddPlayers(player)
	}
	SetupOpponent(g)
	g.AddActor(max)
	g.AddActor(katie)
	g.AddActor(gabe)
}
