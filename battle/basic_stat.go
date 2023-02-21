package battle

import "GoCards/actions"

// Stats represent some value in the game state
type BasicStat[T StatValue] struct {
	actions.Dispatcher[StatChange[T]]
	value T
}

func (stat *BasicStat[T]) Get() T {
	return stat.value
}

func (stat *BasicStat[T]) Set(value T) {
	stat.Dispatch(StatChange[T]{
		ActionBase: actions.NewActionBase(),
		NewValue:   value,
		OldValue:   stat.value,
		Stat:       stat,
	})
}

func (stat *BasicStat[T]) SetDirect(value T) {
	stat.value = value
}
