package main

import "fmt"


func main() {

	// run task 1
	task1()

	// run task 2 (bonus task)
	//task12()
}

func task1() {

	// TASK 1: print all integer numbers from 1 to 50, that are divisible by 5 (e.g. 5, 10, 15...)
	// HINT: use fmt.Println(...) to print number to console

	// TODO: add missing loop
	// TODO: use if for checking if number is divisible by 5
	fmt.Println(5)
}

func task2() {

	// BONUS TASK: print all prime numbers from 1 to 50
	// HINT: use another inner loop

	for i := 1; i < 30; i++ {

		isPrime := false

		// TODO: fill missing code

		if isPrime {
			fmt.Printf("%v is a prime number\n", i)
		}
	}
}
