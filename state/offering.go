package state

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
