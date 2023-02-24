package effects

type IncreaseChargeEffect struct {
	Amount int
}

func (inc *IncreaseChargeEffect) Trigger(context EffectContext) {
	context.Caster.Charge.Set(context.Caster.Charge.Get() + inc.Amount)
}

func IncreaseCharge(amount int) *IncreaseChargeEffect {
	return &IncreaseChargeEffect{
		Amount: amount,
	}
}
