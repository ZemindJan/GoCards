package status

import (
	"GoCards/actions"
	"GoCards/state"
)

type StatusAmount int // -1 for none

type Status struct {
	Prototype StatusPrototype
	ID        actions.ID
	Host      *state.Champion
	Amount    StatusAmount
}
