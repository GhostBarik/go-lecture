package main

import "fmt"


func main() {

	// run task 3
	task3()

	// run task 4
	//task4()

	// run task 5
	//task5()
}

// TASK 3: implement function filter(...), that will read input collection,
//         call predicate function on each element and will return
//         new output collection, containing only elements for which predicate
//         function returns true

// HINT: use build-in function make(..) to create an new empty collection
// HINT: use built-in function append(..) too add new element to the collection

type Predicate = func (int) bool

func filter(predicate Predicate, collection []int) []int {
	// TODO: implement missing part
	return collection
}

func task3() {

	// create input collection (slice)
	input := []int {1, 2, 3, 4}

	// create test predicate function to test if value is > 2
	predicate := func(x int) bool {
		return x > 2
	}

	// apply filter function to filter some elements out
	// and store result in output variable
	output := filter(predicate, input)

	// print results
	fmt.Printf("after filtering: %v \n", output)
}

/////////////////////////////////////////////////////////////////////////////

// TASK 4: implement function transform(...), that will read input collection,
//         apply transformation function on each element and will return
//         new output collection, containing updated elements
//
// HINT: use build-in function make(..) to create an new empty collection
// HINT: use built-in function append(..) too add new element to the collection

type Mapper = func (int) int

func transform(mapper Mapper, collection []int) []int {
	// TODO: implement missing part
	return collection
}

func task4() {

	// create input collection (slice)
	input := []int {1, 2, 3, 4}

	// create test transformation function, that adds +1 to input integer number
	mapper := func (x int) int {
		return x + 1
	}

	// transform input collection and store result in output variable
	output := transform(mapper, input)

	// print results
	fmt.Printf("after filtering: %v \n", output)
}

type Reducer = func (int, int) int

func reduce(reducer Reducer, initial int, collection []int) int {
	// TODO: implement missing part
	return initial
}

// TASK 5: (BONUS TASK) implement reduce(...) function, that will that take input collection
//         along with some initial accumulator value, will iterate over each element of input collection
//         and combine each element with accumulator value
//
// HINT: use build-in function make(..) to create an new empty collection
// HINT: use built-in function append(..) too add new element to the collection

func task5() {

	// create input collection (slice)
	input := []int {1, 2, 3, 4}

	// create test combine function, that adds two integer values together
	combine := func (accumulator int, x int) int {
		return accumulator + x
	}

	// reduce input collection (sum up all numbers in the input collection)
	output2 := reduce(combine, 0, input)

	// print results
	fmt.Printf("after reduce (+): %v \n", output2)
}