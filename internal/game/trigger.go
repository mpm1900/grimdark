package game

type TriggerOn string

const (
	OnDamageRecieve TriggerOn = "on-damage-recieve"
)

type Trigger struct {
	Action
	On       TriggerOn
	Validate func(g Game, t_context Context, m_context Context) bool
}
