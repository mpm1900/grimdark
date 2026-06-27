package game

type TriggerOn string

const (
	OnActorEnter    TriggerOn = "on-actor-enter"
	OnActorLeave    TriggerOn = "on-actor-leave"
	OnDamageSend    TriggerOn = "on-damage-recieve"
	OnDamageRecieve TriggerOn = "on-damage-recieve"
	OnTurnEnd       TriggerOn = "on-turn-end"
)

type Trigger struct {
	Action
	On       TriggerOn
	Validate func(g Game, t_context Context, m_context Context) bool
}
