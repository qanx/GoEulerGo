package main

import "fmt"

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

func sumOfElements(input []int) int {
	sum := 0
	for i := 0; i < len(input); i++ {
		sum += input[i]
	}
	return sum
}

func main() {
	/* Prompt:
	The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
	Find the sum of all the primes below 2000000.
	*/
	fmt.Println("The solution to Euler Problem 10 is: ")
	fmt.Println(sumOfElements(makeBoundedListOfPrimes(2000000)))
}
