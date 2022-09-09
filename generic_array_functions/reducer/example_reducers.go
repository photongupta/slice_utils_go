package reducer

import (contract "array_utils/generic_array_functions")

func Sum(sum int64, num int64) int64 {
	return sum + num
}

func NewPerson(name string) contract.Person {
	return contract.Person{
		Name: name,
		Age:  20,
	}
}

