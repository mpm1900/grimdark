package game

type TriggerOn string

const (
	OnActorEnter    TriggerOn = "on-actor-enter"
	OnActorLeave    TriggerOn = "on-actor-enter"
	OnDamageRecieve TriggerOn = "on-damage-recieve"
)

type Trigger struct {
	Action
	On       TriggerOn
	Validate func(g Game, t_context Context, m_context Context) bool
}
