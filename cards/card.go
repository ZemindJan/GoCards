package cards

import (
	"GoCards/effects"
	"GoCards/state"
)

type Card struct {
	Effect effects.Effect
	Color  state.OfferingColor
}
