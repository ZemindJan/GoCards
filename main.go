package main

import (
	"GoCards/actions"
	"GoCards/content"
	"GoCards/content/basic"
	"GoCards/game"
	"GoCards/state"
	"GoCards/status"
)

func main() {
	(&game.Game{
		Controller1: &game.RandomController{},
		Deck1:       basic.BasicDeck(),
		Controller2: &game.RandomController{},
		Deck2:       basic.BasicDeck(),
		Battle: &state.Battle{
			Player1:     state.NewChampion(content.BasicChampion()),
			Player2:     state.NewChampion(content.BasicChampion()),
			OnDamage:    actions.NewBaseDispatcher[state.DamageAction](),
			TurnManager: state.MakeTurnManager(),
		},
		StatusManager: &status.StatusManager{ActiveStatuses: []status.Status{}},
	}).MainLoop()

}
