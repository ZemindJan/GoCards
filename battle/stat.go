package battle

import "GoCards/actions"

// Value of stats
type StatValue interface {
	int | float32
}

type ValueProvider[T StatValue] interface {
	Get() T
}

// Action to change stat's value
type StatChange[T StatValue] struct {
	*actions.ActionBase
	NewValue T
	OldValue T
	Stat     Stat[T]
}

func (sc StatChange[T]) Execute() {
	sc.Stat.SetDirect(sc.NewValue)
}

type Stat[T StatValue] interface {
	ValueProvider[T]   // Allows stat to be read
	Set(value T)       // Attempt to set the stat via a StatChange action
	SetDirect(value T) // Actually modifies the stat
}
