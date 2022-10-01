package extensions

type slice[T comparable] []T

func (s *slice[T]) Remove(element T) []T {

	newSlice := []T{}

	for _, i := range *s {
		if i != element {
			newSlice = append(newSlice, i)
		}
	}
	return newSlice
}

func Remove[T comparable](slice []T, element T) []T {
	newSlice := []T{}

	for _, i := range slice {
		if i != element {
			newSlice = append(newSlice, i)
		}
	}
	return newSlice
}