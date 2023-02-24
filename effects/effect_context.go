package effects

import "GoCards/state"

type ICard interface {
	GetName() string
	GetColor() state.OfferingColor
}

// Provides Effects with the data needed to run
type EffectContext struct {
	Caster *state.Champion
	Target *state.Champion
	Battle *state.Battle
	Card   ICard
}
