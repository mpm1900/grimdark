package instance

import (
	"grimdark/internal/game"
	"grimdark/internal/game/actors"
	"grimdark/internal/game/items"
	"grimdark/internal/game/weapons"

	"github.com/google/uuid"
)

func SetupPlayer(g *game.Game, config game.TeamConfig) {
	player := game.NewPlayer(config.User)
	g.AddPlayers(player)

	for _, a_config := range config.Actors {
		actor, ok := actors.HydrateActorConfig(a_config)
		if ok {
			g.AddActor(actor, player.ID)
		}
	}
}

func SetupOpponent(g *game.Game, user game.User) {
	config := game.TeamConfig{
		User: game.User{
			ID: uuid.New(),
		},
		Actors: []game.ActorConfig{
			{
				Class:   actors.NecronWarrior.ID,
				WeaponR: weapons.SlashSword.ID,
			},
			{
				Class:   actors.NecronWarrior.ID,
				WeaponR: weapons.SlashSword.ID,
			},
			{
				Class:   actors.NecronWarrior.ID,
				WeaponR: weapons.SlashSword.ID,
			},
			{
				Class:   actors.NecronWarrior.ID,
				WeaponR: weapons.SlashSword.ID,
			},
		},
	}

	SetupPlayer(g, config)
	for _, a := range g.GetActorsByPlayer(config.User.ID) {
		open_positions := g.State().GetOpenPositions(config.User.ID)
		if len(open_positions) > 0 {
			pos := open_positions[0]
			g.SetPosition(a.ID, pos.ID)
		}
	}
}

func SetupGame(g *game.Game, user game.User) {
	config := game.TeamConfig{
		User: user,
		Actors: []game.ActorConfig{
			{
				Class:   actors.Ultramarine.ID,
				WeaponL: weapons.SlashSword.ID,
				WeaponR: weapons.SlashSword.ID,
			},
			{
				Class:   actors.TechPriest.ID,
				WeaponR: weapons.SlashSword.ID,
			},
			{
				Class:   actors.SisterOfBattle.ID,
				Items:   []uuid.UUID{items.TestItem.ID},
				WeaponR: weapons.SlashSword.ID,
			},
			{
				Class:   actors.BloodAngel.ID,
				WeaponR: weapons.SlashSword.ID,
			},
		},
	}

	SetupPlayer(g, config)
	SetupOpponent(g, user)
}
