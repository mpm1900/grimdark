package actions

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

var RecklessStrike = game.Action{
	ID:   uuid.MustParse("019f8fe9-4e63-7343-8c21-94bca920fc18"),
	Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
	Config: game.ActionConfig{
		Name:         "Reckless Strike",
		Description:  "Deals heavy damage, but user takes 1/3 of damage dealt as recoil. This action is only usable in 1st position.",
		Affinity:     game.Kinetic,
		Stat:         game.Melee,
		Accuracy:     game.P(0.90),
		Power:        130,
		Lifesteal:    0,
		Recoil:       1.0 / 3.0,
		Hits:         1,
		CritStage:    0,
		CritModifier: 1.5,
		TargetCount:  1,
		Range:        game.P(1),
	},
	Resolve:          game.MakeAttack(game.AttackConfig{}),
	ValidateContext:  game.ContextTargetLength(1),
	TargetsPredicate: game.CombineFilters(game.ActiveActors, game.OtherActors, game.ActionRange(1)),
	DisabledCheck: func(g *game.Game, source game.Actor) bool {
		position, ok := g.GetPosition(source.PositionID)
		if !ok {
			return true
		}

		return position.Rank != 0
	},
}
