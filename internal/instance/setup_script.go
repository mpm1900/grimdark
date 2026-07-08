package instance

import (
	"grimdark/internal/game"
	"grimdark/internal/game/actors"
	"grimdark/internal/game/items"
	"grimdark/internal/game/weapons"

	"github.com/google/uuid"
)

func SetupOpponent(g *game.Game, user game.User) {
	monsters := g.FindActors(func(g game.Game, a game.Actor, ctx game.Context) bool {
		return a.PlayerID != user.ID
	}, game.NewContext())

	if len(monsters) == 0 {
		opp := game.NewPlayer(uuid.New())
		g.AddPlayers(opp)
		open_positions := g.State().GetOpenPositions(opp.ID)
		newMonster := func() game.Actor {
			return game.NewActor(actors.NecronWarrior())
		}
		for _, pos := range open_positions {
			mnstr := newMonster()
			g.AddActor(mnstr, opp.ID)
			g.SetPosition(mnstr.ID, pos.ID)
		}
		g.AddActor(newMonster(), opp.ID)
	}
}

func SetupGame(g *game.Game, user game.User) {

	player := game.NewPlayer(user.ID)
	player.User = user
	if len(g.State().Players) > 0 {
		player = g.Base().Players[0]
	}
	if len(g.State().Players) == 0 {
		g.AddPlayers(player)
	}

	max := game.NewActor(actors.NewUltramarine())
	max.WeaponL = &weapons.SlashSword
	max.WeaponR = &weapons.SlashSword

	katie := game.NewActor(actors.NewSisterOfBattle())
	katie.AuxStats[game.Speed] = 10
	katie.WeaponL = &weapons.SlashSword
	katie.Item = game.P(items.TestItem())

	gabe := game.NewActor(actors.NewTechPriest())
	gabe.WeaponL = &weapons.SlashSword
	gabe.Item = game.P(items.TestItem())

	other := game.NewActor(actors.NewBloodAngel())
	other.WeaponL = &weapons.SlashSword

	SetupOpponent(g, user)
	g.AddActor(max, player.ID)
	g.AddActor(katie, player.ID)
	g.AddActor(gabe, player.ID)
	g.AddActor(other, player.ID)
}
