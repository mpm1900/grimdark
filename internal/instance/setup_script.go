package instance

import (
	"grimdark/internal/game"
	"grimdark/internal/game/actors"
	"grimdark/internal/game/weapons"

	"github.com/google/uuid"
)

func SetupPlayer(g *game.Game, user game.User, config game.TeamConfig) {
	player := game.NewPlayer(user)
	g.AddPlayers(player)
	actors.ApplyTeamConfig(g, player.ID, config)
}

func SetupOpponent(g *game.Game) {
	user := game.User{
		ID: uuid.New(),
	}
	config := game.TeamConfig{
		Actors: []game.ActorConfig{
			{
				Class:   game.P(actors.NecronWarrior.ID),
				WeaponR: game.P(weapons.SlashSword.ID),
			},
			{
				Class:   game.P(actors.NecronWarrior.ID),
				WeaponR: game.P(weapons.SlashSword.ID),
			},
			{
				Class:   game.P(actors.NecronWarrior.ID),
				WeaponR: game.P(weapons.SlashSword.ID),
			},
			{
				Class:   game.P(actors.NecronWarrior.ID),
				WeaponR: game.P(weapons.SlashSword.ID),
			},
		},
	}

	SetupPlayer(g, user, config)
	for _, a := range g.GetActorsByPlayer(user.ID) {
		open_positions := g.State().GetOpenPositions(user.ID)
		if len(open_positions) > 0 {
			pos := open_positions[0]
			g.SetPosition(a.ID, pos.ID)
		}
	}
}

func SetupGame(g *game.Game, user game.User) {
	SetupOpponent(g)
}
