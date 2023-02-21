package utils

func Contains[T comparable](slice []T, item T) bool {
	for _, other := range slice {
		if other == item {
			return true
		}
	}

	return false
}

// Returns nil if no item matches predicate
func First[T any](slice []T, predicate func(T) bool) *T {
	for _, item := range slice {
		if predicate(item) {
			return &item
		}
	}

	return nil
}

// Returns -1 if no item matches predicate
func FirstIndex[T any](slice []T, predicate func(T) bool) int {
	for index, item := range slice {
		if predicate(item) {
			return index
		}
	}

	return -1
}

// Returns new slice with index removed
func RemoveAt[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

// Returns new slice with first item matching predicate removed
func RemoveFirst[T any](slice []T, predicate func(T) bool) []T {
	index := FirstIndex(slice, predicate)

	// Found none
	if index == -1 {
		return slice
	}

	return RemoveAt(slice, index)
}

// Returns new slice with all matching elements removed
func RemoveAll[T any](slice []T, predicate func(T) bool) []T {
	index := FirstIndex(slice, predicate)

	for index != -1 {
		slice = RemoveAt(slice, index)
		index = FirstIndex(slice, predicate)
	}

	return slice
}
