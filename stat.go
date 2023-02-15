package main

import (
	"GoCards/actions"
)

type Stat[T int | float32] struct {
	OnChange actions.ActionDispatcher[StatChangeAction[T]]
	Value    T
}

func (stat Stat[T]) Change(changeAmount T) {
	stat.OnChange.Dispatch(StatChangeAction[T]{
		stat:          &stat,
		OriginalValue: stat.Value,
		ChangeAmount:  changeAmount,
	})
}

type StatChangeAction[T int | float32] struct {
	actions.BaseAction
	stat          *Stat[T]
	OriginalValue T
	ChangeAmount  T
}

func (action StatChangeAction[T]) Execute() {
	action.OriginalValue += action.ChangeAmount
}
