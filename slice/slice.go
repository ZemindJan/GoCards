package slice

func IndexOf[T comparable](slice []T, element T) int {
	for i, e := range slice {
		if e == element {
			return i
		}
	}

	return -1
}

func Contains[T comparable](slice []T, element T) bool {
	return IndexOf(slice, element) != -1
}

func RemoveAt[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

func RemoveFirst[T comparable](slice []T, element T) []T {
	return RemoveAt(slice, IndexOf(slice, element))
}
