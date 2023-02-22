package state

// Value of stats
type ValueType interface {
	int | float32
}

type Value[T ValueType] interface {
	Get() T
}
