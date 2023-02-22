package state

import (
	"GoCards/actions"
	"GoCards/utils"

	"golang.org/x/exp/constraints"
)

// A bounded stat must be ordered to make sense
type BoundedStatValue interface {
	ValueType
	constraints.Ordered
}

type BoundedStat[T BoundedStatValue] struct {
	actions.BaseDispatcher[StatChange[T]]
	value   T
	Maximum Stat[T]
	Minimum Stat[T]
}

func (stat *BoundedStat[T]) Get() T {
	return stat.value
}

func (stat *BoundedStat[T]) Set(value T) {
	stat.Dispatch(StatChange[T]{
		ActionBase: actions.NewActionBase(),
		NewValue:   value,
		OldValue:   stat.value,
		Stat:       stat,
	})
}

func (stat *BoundedStat[T]) setDirect(value T) {
	stat.value = utils.Clamp(value, stat.Minimum.Get(), stat.Maximum.Get())
}

func MakeBoundedStat[T BoundedStatValue](value T, min T, max T) BoundedStat[T] {
	return BoundedStat[T]{
		actions.BaseDispatcher[StatChange[T]]{},
		value,
		NewBasicStat(min),
		NewBasicStat(max),
	}
}

func NewBoundedStat[T BoundedStatValue](value T, min T, max T) *BoundedStat[T] {
	val := MakeBoundedStat(value, min, max)
	return &val
}
