package state

import "GoCards/actions"

type Battle struct {
	Player1     *Champion
	Player2     *Champion
	turn        int
	OnDamage    actions.Dispatcher[DamageAction]
	TurnManager TurnManager
}

// Returns the oppponent of a champion
func (btl *Battle) Opponent(champion *Champion) *Champion {
	switch champion {
	case btl.Player1:
		return btl.Player2
	case btl.Player2:
		return btl.Player2
	default:
		return nil
	}
}

func (btl *Battle) GetVictoryState() VictoryState {
	if btl.Player1.Health.Get() > 0 && btl.Player2.Health.Get() > 0 {
		return NoVictory
	} else if btl.Player1.Health.Get() > 0 {
		return Player1Victory
	} else if btl.Player2.Health.Get() > 0 {
		return Player2Victory
	} else if btl.Player1.DamageTakenSinceLastTurn == btl.Player2.DamageTakenSinceLastTurn {
		return NoVictory
	} else if btl.Player1.DamageTakenSinceLastTurn < btl.Player2.DamageTakenSinceLastTurn {
		return Player1Victory
	} else {
		return Player2Victory
	}
}
