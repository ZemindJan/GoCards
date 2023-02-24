package effects

type SapChargeEffect struct {
	Amount int
}

func (sap *SapChargeEffect) Trigger(context EffectContext) {
	context.Target.Charge.Set(context.Caster.Charge.Get() - sap.Amount)
}

func SapCharge(amount int) *SapChargeEffect {
	return &SapChargeEffect{
		Amount: amount,
	}
}
