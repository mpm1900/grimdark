package instance

import (
	"grimdark/internal/game"
	"grimdark/internal/game/actors"
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
			opp_def.SpriteURL = "/img/nec.png"
			return game.NewActor(opp.ID, opp_def)
		}
		for _, pos := range open_positions {
			mnstr := newMonster()
			g.AddActor(mnstr)
			g.SetPosition(mnstr.ID, pos.ID)
		}
		g.AddActor(newMonster())
	}
}

func SetupGame(g *game.Game, user game.User) {

	player := game.NewPlayer(user.ID)
	player.User = user
	if len(g.State().Players) > 0 {
		player = g.Base().Players[0]
	}

	max_def := actors.NewUltramarine()
	max := game.NewActor(player.ID, max_def)
	max.WeaponL = &weapons.SlashSword
	max.WeaponR = &weapons.SlashSword

	katie_def := actors.NewSisterOfBattle()
	katie := game.NewActor(player.ID, katie_def)
	katie.AuxStats[game.Speed] = 10
	katie.WeaponL = &weapons.SlashSword
	katie.Item = game.P(items.TestItem())

	gabe_def := actors.NewTechPriest()
	gabe := game.NewActor(player.ID, gabe_def)
	gabe.WeaponL = &weapons.SlashSword
	gabe.Item = game.P(items.TestItem())

	other := game.NewActor(player.ID, actors.NewBloodAngel())
	other.WeaponL = &weapons.SlashSword

	if len(g.State().Players) == 0 {
		g.AddPlayers(player)
	}
	SetupOpponent(g)
	g.AddActor(max)
	g.AddActor(katie)
	g.AddActor(gabe)
	g.AddActor(other)
}
