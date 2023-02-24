package effects

import "GoCards/state"

// An effect that scales with the amount of offering
type ScaledEffect struct {
	OfferingColor state.OfferingColor
	Effect0       Effect
	Effect1       Effect
	Effect2       Effect
	Effect3       Effect
}

func (se *ScaledEffect) Trigger(context EffectContext) {
	offeringAmount := context.Caster.Offering[se.OfferingColor]
	map[int]Effect{
		0: se.Effect0,
		1: se.Effect1,
		2: se.Effect2,
		3: se.Effect3,
	}[offeringAmount].Trigger(context)
}
