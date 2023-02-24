package state

import "fmt"

type OfferingColor string

const (
	Red   OfferingColor = "Red"
	Green OfferingColor = "Green"
	Blue  OfferingColor = "Blue"
)

type Offering map[OfferingColor]int

func MakeOffering(red int, green int, blue int) Offering {
	return map[OfferingColor]int{
		Red:   red,
		Green: green,
		Blue:  blue,
	}
}

func (offering Offering) ToString() string {
	return fmt.Sprintf("{%d Red, %d Green, %d Blue}", offering[Red], offering[Green], offering[Blue])
}
