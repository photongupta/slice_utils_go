package filter

import "array_utils/generic_array_functions"

func IsEven(num int64) bool {
	return num%2 == 0
}

func IsNotEmpty(value string) bool {
	return value != ""
}

func EligibleForVote(p generic_array_functions.Person) bool {
	return p.Age > 18
}

