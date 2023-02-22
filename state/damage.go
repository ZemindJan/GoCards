package state

import "GoCards/actions"

type DamageAction struct {
	*actions.ActionBase
	Amount int
	Target *Champion
}

func (da DamageAction) Execute() {
	da.Target.Health.Set(
		da.Target.Health.Get() - da.Amount,
	)
}
