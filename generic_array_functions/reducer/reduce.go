package reducer

type Reducer[R any, T any] func(R,T) R

func Reduce[T any, R any](reducer Reducer[R,T], array []T) R {
	acc := new(R)
	for _, element := range array {
		temp := reducer(*acc, element)
		acc = &temp
	}
	return *acc
}
