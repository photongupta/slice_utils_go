package main

import (
	contract "array_utils/generic_array_functions"
	"array_utils/generic_array_functions/filter"
	"array_utils/generic_array_functions/mapper"
	"array_utils/generic_array_functions/reducer"
	"fmt"
)

func main() {
	fmt.Println("Map.......")
	runMap()
	fmt.Println("Filter.......")
	runFilter()
	fmt.Println("Reduce.......")
	runReduce()
}

func runMap() {
	//int64 to int64
	numbers := []int64{1, 2, 3, 4, 54, 9}
	mappedArray1 := mapper.Map[int64](mapper.IncrementByOne, numbers)
	fmt.Println(mappedArray1)

	//string to struct
	names := []string{"John", "Kale", "Peter", "Jane"}
	mappedArray2 := mapper.Map[string,contract.Person](mapper.NewPerson, names)
	fmt.Println(mappedArray2)
}

func runFilter() {
	//int64
	numbers := []int64{1, 2, 3, 4, 54, 9}
	filteredArray1 := filter.Filter[int64](filter.IsEven, numbers)
	fmt.Println(filteredArray1)

	//string
	words := []string{"apple", "", "ball", "", ""}
	filteredArray2 := filter.Filter[string](filter.IsNotEmpty, words)
	fmt.Println(filteredArray2)

	//struct
	persons := []contract.Person{{
		Name: "John",
		Age:  35,
	},
		{
			Name: "Jane",
			Age:  10,
		},
		{
			Name: "Peter",
			Age:  50,
		}}
	filteredArray3 := filter.Filter[contract.Person](filter.EligibleForVote, persons)
	fmt.Println(filteredArray3)
}

func runReduce() {
	//int64 to int64
	numbers := []int64{1, 2, 3, 4, 5, 10}
	sum := reducer.Reduce[int64](reducer.Sum, numbers)
	fmt.Println(sum)
}
