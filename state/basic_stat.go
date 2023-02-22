package state

import "GoCards/actions"

// Stats represent some value in the game state
type BasicStat[T ValueType] struct {
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

func (stat *BasicStat[T]) setDirect(value T) {
	stat.value = value
}

func MakeBasicStat[T ValueType](value T) BasicStat[T] {
	return BasicStat[T]{
		actions.NewBaseDispatcher[StatChange[T]](),
		value,
	}
}

func NewBasicStat[T ValueType](value T) *BasicStat[T] {
	stat := MakeBasicStat(value)
	return &stat
}
