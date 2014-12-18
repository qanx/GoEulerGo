package main

import "fmt"

func setListOfMultiples(upperLimit int, multipleA int, multipleB int) []int {
	listOfMultiples := []int{}
	for i := 1; i < upperLimit; i++ {
		if i%multipleA == 0 || i%multipleB == 0 {
			listOfMultiples = append(listOfMultiples, i)
		}
	}
	return listOfMultiples
}

func sumOfArrayElements(integerArray []int) int {
	total := 0
	for _, num := range integerArray {
		total += num
	}
	return total
}

func main() {
	// Prompt:
	//If we list all the natural numbers below 10 that are multiples of 3 or 5, we
	//get 3, 5, 6 and 9. The sum of these multiples is 23.
	//Find the sum of all the multiples of 3 or 5 below 1000.

	fmt.Printf("The solution to Euler Problem 1 is: %d\n", sumOfArrayElements(setListOfMultiples(1000, 3, 5)))

}
