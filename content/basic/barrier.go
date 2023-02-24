package basic

import (
	"GoCards/cards"
	"GoCards/effects"
	"GoCards/state"
)

func Barrier() cards.Card {
	return cards.Card{
		Name:  "Barrier",
		Color: state.Blue,
		Effect: &effects.ScaledEffect{
			OfferingColor: state.Blue,
			Effect0:       effects.Damage(100),
			Effect1: effects.Compose(
				effects.Damage(100),
				effects.IncreaseCharge(1),
			),
			Effect2: effects.Compose(
				effects.Damage(100),
				effects.IncreaseCharge(2),
			),
			Effect3: effects.Compose(
				effects.Damage(100),
				effects.IncreaseCharge(3),
			),
		},
	}
}
