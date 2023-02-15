package main

import (
	. "GoCards/actions"
	. "GoCards/slice"
	"fmt"
)

type test struct {
}

func (t test) Notify(oldValue int, newValue int) {
	fmt.Println("Changed health from %d to %d", oldValue, newValue)
}

func main() {
	RemoveAt(make([]int, 2), 1)
	fmt.Println("hello world")

	var p1 = Player{}

	var listener ActionListener[StatChangeAction[int]] = ActionListenerFunction[StatChangeAction[int]]{
		Function: func(change StatChangeAction[int]) {
			fmt.Println()
		},
	}

	p1.health.OnChange.OnStart(&listener)

	p1.health.Change(10)

}
