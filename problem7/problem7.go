package main

import "fmt"

func makeBoundedListOfPrimes(lengthOfList int) []int {
	//Note: Intentionally not including 1 in list of prime factors since
	//all numbers divisible by 1.
	primeList := []int{2}
	i := 1
	for len(primeList) < lengthOfList {
		i++
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

func main() {
	/* Prompt:
	By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see
	that the 6th prime is 13.

	What is the 10001st prime number?
	*/

	fmt.Println("The solution to Euler Problem 6 is: ")
	sliceOfPrimes := makeBoundedListOfPrimes(10001)
	fmt.Println(sliceOfPrimes[len(sliceOfPrimes)-1])
}
