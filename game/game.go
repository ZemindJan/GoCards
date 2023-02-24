package game

import (
	"GoCards/actions"
	"GoCards/cards"
	"GoCards/effects"
	"GoCards/state"
	"GoCards/status"
	"GoCards/ui"
	"fmt"
)

type Game struct {
	Controller1   ChampionController
	Deck1         cards.Deck
	Controller2   ChampionController
	Deck2         cards.Deck
	Battle        *state.Battle
	StatusManager *status.StatusManager
}

func (gm *Game) MainLoop() state.VictoryState {
	for gm.Battle.GetVictoryState() == state.NoVictory {
		gm.StartTurn()
		gm.OfferingPhase()
		gm.PlayPhase()
		gm.EndTurn()
	}

	return gm.Battle.GetVictoryState()
}

func (gm *Game) StartTurn() {
	// Every phase opens with minimum 1 health
	if gm.Battle.Player1.Health.Get() < 1 {
		gm.Battle.Player1.Health.SetDirect(1)
	}

	if gm.Battle.Player2.Health.Get() < 1 {
		gm.Battle.Player2.Health.SetDirect(1)
	}

	gm.Battle.Player1.DamageTakenSinceLastTurn = 0
	gm.Battle.Player2.DamageTakenSinceLastTurn = 0

	gm.Battle.TurnManager.OnStartTurn.Dispatch(state.StartTurn{
		Event:  &actions.Event{},
		Battle: gm.Battle,
	})
}

func (gm *Game) OfferingPhase() {
	gm.Battle.TurnManager.OnStartOfferingPhase.Dispatch(state.StartOfferingPhase{
		Event:  &actions.Event{},
		Battle: gm.Battle,
	})

	offering1 := gm.Controller1.GetOffering(gm.Battle.Player1, &gm.Deck1, gm.Battle)
	offering2 := gm.Controller2.GetOffering(gm.Battle.Player2, &gm.Deck2, gm.Battle)

	gm.Battle.Player1.Offering = offering1
	gm.Battle.Player2.Offering = offering2

	fmt.Printf("Player1 made offering %s\n", offering1.ToString())
	fmt.Printf("Player2 made offering %s\n", offering2.ToString())

	gm.Battle.TurnManager.OnEndOfferingPhase.Dispatch(state.EndOfferingPhase{
		Event:  &actions.Event{},
		Battle: gm.Battle,
	})

	ui.PressEnterToContinue()
}

func (gm *Game) PlayPhase() {
	gm.Battle.TurnManager.OnStartPlayPhase.Dispatch(state.StartPlayPhase{
		Event:  &actions.Event{},
		Battle: gm.Battle,
	})

	card1 := gm.Controller1.GetCard(gm.Battle.Player1, &gm.Deck1, gm.Battle)
	card2 := gm.Controller2.GetCard(gm.Battle.Player2, &gm.Deck2, gm.Battle)

	card1.Effect.Trigger(effects.EffectContext{
		Caster: gm.Battle.Player1,
		Target: gm.Battle.Player2,
		Battle: gm.Battle,
		Card:   card1,
	})

	card2.Effect.Trigger(effects.EffectContext{
		Caster: gm.Battle.Player2,
		Target: gm.Battle.Player1,
		Battle: gm.Battle,
		Card:   card2,
	})

	gm.Battle.TurnManager.OnEndPlayPhase.Dispatch(state.EndPlayPhase{
		Event:  &actions.Event{},
		Battle: gm.Battle,
	})
}

func (gm *Game) EndTurn() {
	gm.Battle.TurnManager.OnEndTurn.Dispatch(state.EndTurn{
		Event:  &actions.Event{},
		Battle: gm.Battle,
	})
}
