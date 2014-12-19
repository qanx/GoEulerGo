package main

import (
	"fmt"
	"math"
)

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

func makeListOfPrimeFactors(value int, primeList []int) []int {
	factorList := []int{}
	for i := 0; i < len(primeList); i++ {
		if value%primeList[i] == 0 {
			factorList = append(factorList, primeList[i])
		}
	}
	return factorList
}

func main() {
	/* Prompt:
	The prime factors of 13195 are 5, 7, 13 and 29.
	What is the largest prime factor of the number 600851475143 ?
	*/

	value := 600851475143
	fmt.Println("The solution to Euler Problem 3 is: ")
	primeFactors := makeListOfPrimeFactors(value, makeBoundedListOfPrimes(int(math.Floor(math.Sqrt(float64(value))))))
	fmt.Println(primeFactors[len(primeFactors)-1])
}
