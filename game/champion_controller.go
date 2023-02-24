package game

import (
	"GoCards/cards"
	"GoCards/state"
)

type ChampionController interface {
	GetOffering(myChampion *state.Champion, myDeck *cards.Deck, battle *state.Battle) state.Offering
	GetCard(myChampion *state.Champion, myDeck *cards.Deck, battle *state.Battle) *cards.Card
}
