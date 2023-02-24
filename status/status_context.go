package status

import "GoCards/state"

type StatusContext struct {
	Host   *state.Champion
	Battle *state.Battle
	Status *Status
}
