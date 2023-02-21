package utils

import (
	"golang.org/x/exp/constraints"
)

// Constraints a value to an inclusive lower and upper bound
// It is up to the user to ensure lower < upper
func Clamp[T constraints.Ordered](value T, lower T, upper T) T {
	if value > upper {
		return upper
	} else if value < lower {
		return lower
	} else {
		return value
	}
}
