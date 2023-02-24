package basic

import (
	"GoCards/cards"
	"GoCards/effects"
	"GoCards/state"
)

func Gadget() cards.Card {
	return cards.Card{
		Name:  "Gadget",
		Color: state.Green,
		Effect: &effects.ScaledEffect{
			OfferingColor: state.Green,
			Effect0:       effects.Damage(100),
			Effect1: effects.Compose(
				effects.Damage(100),
				effects.SapCharge(1),
			),
			Effect2: effects.Compose(
				effects.Damage(100),
				effects.SapCharge(2),
			),
			Effect3: effects.Compose(
				effects.Damage(100),
				effects.SapCharge(3),
			),
		},
	}
}
