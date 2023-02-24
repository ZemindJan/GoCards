package cards

import (
	"GoCards/effects"
	"GoCards/state"
)

type Card struct {
	Name   string
	Effect effects.Effect
	Color  state.OfferingColor
}

func (card *Card) GetName() string {
	return card.Name
}

func (card *Card) GetColor() state.OfferingColor {
	return card.Color
}
