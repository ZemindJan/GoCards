package state

import "GoCards/actions"

// Action to change stat's value
type StatChange[T ValueType] struct {
	*actions.ActionBase
	NewValue T
	OldValue T
	Stat     Stat[T]
}

func (sc StatChange[T]) Execute() {
	sc.Stat.setDirect(sc.NewValue)
}

type Stat[T ValueType] interface {
	Value[T]           // Allows stat to be read
	Set(value T)       // Attempt to set the stat via a StatChange action
	setDirect(value T) // Actually modifies the stat
}
