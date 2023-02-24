package basic

import (
	"GoCards/cards"
	"GoCards/effects"
	"GoCards/state"
)

func Weapon() cards.Card {
	return cards.Card{
		Name:  "Weapon",
		Color: state.Red,
		Effect: &effects.ScaledEffect{
			OfferingColor: state.Red,
			Effect0:       effects.Damage(100),
			Effect1:       effects.Damage(125),
			Effect2:       effects.Damage(150),
			Effect3:       effects.Damage(200),
		},
	}
}
