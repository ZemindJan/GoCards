package effects

import (
	"GoCards/actions"
	"GoCards/state"
)

type DamageEffect struct {
	Amount int
}

func (de *DamageEffect) Trigger(context EffectContext) {
	context.Battle.OnDamage.Dispatch(state.DamageAction{
		ActionBase: actions.NewActionBase(),
		Amount:     de.Amount,
		Target:     context.Target,
	})
}

func Damage(amount int) *DamageEffect {
	return &DamageEffect{
		Amount: amount,
	}
}
