package game

import "github.com/google/uuid"

type ActionContext struct {
	Action         Action
	Source         Actor
	transactions   []Transaction
	pending_damage map[uuid.UUID]float64
}

func (ac *ActionContext) Push(transactions ...Transaction) {
	ac.transactions = append(ac.transactions, transactions...)
}

func (ac *ActionContext) Concat(transactions []Transaction) {
	ac.transactions = append(ac.transactions, transactions...)
}

func (ac *ActionContext) Done() []Transaction {
	return ac.transactions
}

func (ac *ActionContext) RecordDamage(target uuid.UUID, damage float64) {
	ac.pending_damage[target] = damage
}
func (ac *ActionContext) PendingDamage(target uuid.UUID) float64 {
	return ac.pending_damage[target]
}
