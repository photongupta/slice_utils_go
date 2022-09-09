package mapper

type Mapper[T any, R any] func(T) R

func Map[T any, R any](mapper Mapper[T,R], array []T) []R {
	mappedArray := make([]R, len(array))
	for index, element := range array {
		mappedArray[index] = mapper(element)
	}
	return mappedArray
}
