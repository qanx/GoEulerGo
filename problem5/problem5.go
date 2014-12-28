package main

import (
	"fmt"
	"math"
)

/* Keeping Pop for posterity
func pop(inputSlice []int, element int) []int {
	inputSlice[element] = inputSlice[element+1]
	return append(inputSlice[:element+1], inputSlice[element+2:]...)
}*/

func makeBoundedListOfPrimes(maxValue int) []int {
	//Note: Intentionally not including 1 in list of prime factors since
	//all numbers divisible by 1.
	primeList := []int{2}
	for i := 2; i < maxValue; i++ {
		for j := 0; j < len(primeList); j++ {
			if i%primeList[j] == 0 {
				break
			} else if j == len(primeList)-1 {
				primeList = append(primeList, i)
			}
		}
	}
	return primeList
}

func makeListOfMultiples(sliceOfPrimes []int, maxValue int) []int {
	sliceOfMultiples := make([]int, len(sliceOfPrimes))
	for i := 0; i < len(sliceOfPrimes); i++ {
		exponent := 0
		for int(math.Pow(float64(sliceOfPrimes[i]), float64(exponent))) < maxValue {
			sliceOfMultiples[i] = int(math.Pow(float64(sliceOfPrimes[i]), float64(exponent)))
			exponent++
		}
	}
	return sliceOfMultiples
}

func productOfSliceElements(inputSlice []int) int {
	product := 1
	for i := 0; i < len(inputSlice); i++ {
		product = product * inputSlice[i]
	}
	return product
}

func main() {
	/* Prompt:
	2520 is the smallest number that can be divided by each of the numbers
	from 1 to 10 without any remainder.
	What is the smallest positive number that is evenly divisible by all of the
	numbers from 1 to 20?
	*/

	fmt.Println("The solution to Euler Problem 5 is: ")
	fmt.Println(productOfSliceElements(makeListOfMultiples(makeBoundedListOfPrimes(20), 20)))

}
