package mapper

import (contract "array_utils/generic_array_functions")

func IncrementByOne(num int64) int64 {
	num++
	return num
}

func NewPerson(name string) contract.Person {
	return contract.Person{
		Name: name,
		Age:  20,
	}
}

