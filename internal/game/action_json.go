package game

import "github.com/google/uuid"

type actionJSON struct {
	ID           uuid.UUID    `json:"ID"`
	Config       ActionConfig `json:"config"`
	Cooldown     int          `json:"cooldown"`
	Entity       Entity       `json:"entity"`
	IsDisabled   bool         `json:"is_disabled"`
	Tags         []ActionTag  `json:"tags"`
	Uncancelable bool         `json:"uncancelable"`
	Uses         int          `json:"uses"`
}

func (a Action) ToJSON(g *Game, source Actor) actionJSON {
	state := source.ActionStates[a.ID]
	config := a.Config
	config.Cooldown = config.Cooldown + state.CooldownBonus
	config.Priority = config.Priority + state.PriorityBonus
	if config.Range != nil {
		config.Range = P(*config.Range + state.RangeBonus)
	}
	json := actionJSON{
		ID:           a.ID,
		Config:       config,
		Cooldown:     state.Cooldown,
		Entity:       a.Entity,
		IsDisabled:   a.Disabled(g, source),
		Tags:         a.Tags,
		Uncancelable: a.Uncancelable,
		Uses:         state.Uses,
	}

	if json.Config.Accuracy != nil {
		acc := *json.Config.Accuracy * source.Stats[Accuracy]
		json.Config.Accuracy = &acc
	}

	json.Config.CritChance = GetCriticalChance(json.Config.CritStage + source.Stages[CriticalChance])
	json.Config.CritModifier = json.Config.CritModifier * source.Stats[CriticalDamage]

	return json
}
func (a Action) ToJSONStatic() actionJSON {
	return actionJSON{
		ID:           a.ID,
		Config:       a.Config,
		Cooldown:     0,
		Entity:       a.Entity,
		IsDisabled:   false,
		Tags:         a.Tags,
		Uncancelable: a.Uncancelable,
		Uses:         0,
	}
}
