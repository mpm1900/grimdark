package actions

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

var CalledShot = game.Action{
	ID:   uuid.MustParse("019f85ac-e872-7fd0-9511-a7df176f402f"),
	Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
	Config: game.ActionConfig{
		Name:         "Called Shot",
		Description:  "This action cannot miss but is -1 priority. This action is only usable from the 3rd position.",
		Affinity:     game.Kinetic,
		Stat:         game.Ranged,
		Power:        95,
		Lifesteal:    0,
		Hits:         1,
		CritChance:   0,
		CritModifier: 2,
		TargetCount:  1,
		Priority:     -1,
	},
	Resolve:          game.MakeAttack(game.AttackConfig{}),
	ValidateContext:  game.ContextTargetLength(1),
	TargetsPredicate: game.CombineFilters(game.ActiveActors, game.OtherActors),
	DisabledCheck: func(g *game.Game, source game.Actor) bool {
		position, ok := g.GetPosition(source.PositionID)
		if !ok {
			return true
		}

		return position.Rank != 2
	},
}
