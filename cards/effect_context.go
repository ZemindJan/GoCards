package cards

import "GoCards/state"

// Provides Effects with the data needed to run
type EffectContext struct {
	Caster *state.Champion
	Target *state.Champion
	Battle *state.Battle
	Card   *Card
}
