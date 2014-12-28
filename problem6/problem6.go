package main

import (
	"fmt"
	"math"
)

func setSequentialSlice(maxValue int) []int {
	sequentialSlice := []int{1}
	for i := 1; i < maxValue; i++ {
		sequentialSlice = append(sequentialSlice, i+1)
	}
	return sequentialSlice
}

func setSumOfSquareOfElements(inputSlice []int) int {
	sum := 0
	for i := 0; i < len(inputSlice); i++ {
		sum += int(math.Pow(float64(inputSlice[i]), 2))
	}
	return sum
}

func setSumOfElements(inputSlice []int) int {
	sum := 0
	for i := 0; i < len(inputSlice); i++ {
		sum += inputSlice[i]
	}
	return sum
}

func main() {
	/* Prompt:
	  The sum of the squares of the first ten natural numbers is:
		(1^2) + (2^2) + ... + (10^2) = 385
		The square of the sum of the first ten natural numbers is:
		(1 + 2 + ... + 10)^2 = 3025
		Hence the difference between the sum of the squares of the first ten natural
		numbers and the square of the sum is 3025 − 385 = 2640.

		Find the difference between the sum of the squares of the first one hundred
		natural numbers and the square of the sum.
	*/
	fmt.Println("The solution to Euler Problem 6 is: ")
	sequentialSlice := setSequentialSlice(100)
	sumOfSquaredElements := setSumOfSquareOfElements(sequentialSlice)
	sumOfElementsSquared := int(math.Pow(float64(setSumOfElements(sequentialSlice)), 2))
	fmt.Println(sumOfElementsSquared - sumOfSquaredElements)
}
