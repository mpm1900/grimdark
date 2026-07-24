package effects

import "grimdark/internal/game"

func Burned() game.Effect {
	effect := game.EffectTargets(
		game.EffectPriorityStatus,
		func(g *game.Game, a game.Actor, ctx game.Context) game.Actor {
			a.Status = game.StatusBurned
			return a
		},
	)
	effect.Name = "Burned"
	effect.Description = "On turn end, this actor takes 8% damage."
	effect.CheckSuccess = game.EffectGainTargetsOnSuccess
	effect.Triggers = append(effect.Triggers, game.Trigger{
		On: game.OnTurnEnd,
		Action: game.Action{
			Config: game.ActionConfig{
				Name: "Burned",
			},
			Resolve: func(g *game.Game, ctx game.Context, this game.ActionContext) []game.Transaction {
				dmg_context := game.MakeContextFor(this.Source, this.Source)
				health := this.Source.Stats[game.Health]
				amount := health * 0.08
				this.Push(game.DamageTargets(amount, false).Bind(dmg_context))
				return this.Done()
			},
		},
	})
	effect.Check = func(g *game.Game, ctx game.Context) bool {
		targets := g.GetTargets(ctx)
		for _, target := range targets {
			if target.Status != game.StatusNone {
				return false
			}
		}

		return true
	}

	return effect
}
