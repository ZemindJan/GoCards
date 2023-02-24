package basic

import (
	"GoCards/cards"
)

func BasicDeck() cards.Deck {
	return cards.Deck{
		Cards: []cards.Card{
			Weapon(),
			Gadget(),
			Barrier(),
		},
	}
}
