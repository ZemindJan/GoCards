package content

import "GoCards/state"

func BasicChampion() *state.ChampionPrototype {
	return &state.ChampionPrototype{
		Name:      "Basic Dude",
		MaxHealth: 80,
		MaxCharge: 8,
	}
}
