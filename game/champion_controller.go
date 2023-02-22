package game

import "GoCards/state"

type ChampionController interface {
	GetOffering(myChampion *state.Champion, battle *state.Battle)
}
