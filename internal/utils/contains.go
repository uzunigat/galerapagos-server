package utils

func ArrayContains[T comparable](array []T, element T) bool {
	for _, arrayElement := range array {
		if arrayElement == element {
			return true
		}
	}
	return false
}
