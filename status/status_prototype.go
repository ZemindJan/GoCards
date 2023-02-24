package status

import (
	"GoCards/ui"
)

type StatusPrototype interface {
	ui.Named
	Attach(context StatusContext)
	Remove(context StatusContext)
	Add(amt1 StatusAmount, amt2 StatusAmount) StatusAmount
}

type StatusPrototypeBase struct {
	Name string
}
