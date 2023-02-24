package state

import (
	"GoCards/actions"
)

type TurnManager struct {
	OnStartTurn          actions.Dispatcher[StartTurn]
	OnEndTurn            actions.Dispatcher[EndTurn]
	OnStartOfferingPhase actions.Dispatcher[StartOfferingPhase]
	OnEndOfferingPhase   actions.Dispatcher[EndOfferingPhase]
	OnStartPlayPhase     actions.Dispatcher[StartPlayPhase]
	OnEndPlayPhase       actions.Dispatcher[EndPlayPhase]
}

func MakeTurnManager() TurnManager {
	return TurnManager{
		OnStartTurn:          actions.NewBaseDispatcher[StartTurn](),
		OnEndTurn:            actions.NewBaseDispatcher[EndTurn](),
		OnStartOfferingPhase: actions.NewBaseDispatcher[StartOfferingPhase](),
		OnEndOfferingPhase:   actions.NewBaseDispatcher[EndOfferingPhase](),
		OnStartPlayPhase:     actions.NewBaseDispatcher[StartPlayPhase](),
		OnEndPlayPhase:       actions.NewBaseDispatcher[EndPlayPhase](),
	}
}

type StartTurn struct {
	*actions.Event
	Battle *Battle
}

type StartOfferingPhase struct {
	*actions.Event
	Battle *Battle
}

type EndOfferingPhase struct {
	*actions.Event
	Battle *Battle
}

type StartPlayPhase struct {
	*actions.Event
	Battle *Battle
}

type EndPlayPhase struct {
	*actions.Event
	Battle *Battle
}

type EndTurn struct {
	*actions.Event
	Battle *Battle
}
