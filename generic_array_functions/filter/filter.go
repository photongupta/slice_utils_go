package filter

type Predicate[T any] func(T) bool

func Filter[T any](predicate Predicate[T], array []T) []T {
	filteredArray := make([]T, 0)
	for _, element := range array {
		if predicate(element) {
			filteredArray = append(filteredArray, element)
		}
	}
	return filteredArray
}



