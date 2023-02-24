package effects

type CompositeEffect struct {
	effects []Effect
}

func (ce *CompositeEffect) Trigger(context EffectContext) {
	for _, effect := range ce.effects {
		effect.Trigger(context)
	}
}

func Compose(effects ...Effect) *CompositeEffect {
	return &CompositeEffect{
		effects: effects,
	}
}
