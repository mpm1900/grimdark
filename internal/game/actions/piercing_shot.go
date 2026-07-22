package actions

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

var PiercingShot = game.Action{
	ID:   uuid.MustParse("019f87df-76c5-78b1-a12b-0d485ee54dad"),
	Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
	Config: game.ActionConfig{
		Name:         "Piercing Shot",
		Description:  "This action also damages all enemy actors positioned behind the target. This action is only usable from 2nd or 3rd position.",
		Affinity:     game.Kinetic,
		Stat:         game.Ranged,
		Accuracy:     game.P(0.80),
		Power:        85,
		Lifesteal:    0,
		Hits:         1,
		CritStage:    0,
		CritModifier: 1.5,
		TargetCount:  1,
	},
	Resolve:          game.MakeAttack(game.AttackConfig{}),
	MapContext:       game.CtxTargetPostCollateral(),
	ValidateContext:  game.ContextTargetLength(1),
	TargetsPredicate: game.CombineFilters(game.ActiveActors, game.Enemies),
	DisabledCheck: func(g *game.Game, source game.Actor) bool {
		position, ok := g.GetPosition(source.PositionID)
		if !ok {
			return true
		}

		return position.Rank == 1
	},
}
