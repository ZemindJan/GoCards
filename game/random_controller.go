package game

import (
	"GoCards/cards"
	"GoCards/state"
	"math/rand"
)

type RandomController struct {
}

func (rc *RandomController) GetOffering(myChampion *state.Champion, myDeck *cards.Deck, battle *state.Battle) state.Offering {
	return []state.Offering{
		{state.Red: 3, state.Green: 0, state.Blue: 0},
		{state.Red: 0, state.Green: 3, state.Blue: 0},
		{state.Red: 0, state.Green: 0, state.Blue: 3},
		{state.Red: 2, state.Green: 1, state.Blue: 0},
		{state.Red: 2, state.Green: 0, state.Blue: 1},
		{state.Red: 1, state.Green: 2, state.Blue: 0},
		{state.Red: 0, state.Green: 2, state.Blue: 1},
		{state.Red: 0, state.Green: 1, state.Blue: 2},
		{state.Red: 1, state.Green: 0, state.Blue: 2},
		{state.Red: 1, state.Green: 1, state.Blue: 1},
	}[rand.Int()%10]
}

func (rc *RandomController) GetCard(myChampion *state.Champion, myDeck *cards.Deck, battle *state.Battle) *cards.Card {
	return &myDeck.Cards[rand.Int()%len(myDeck.Cards)]
}
