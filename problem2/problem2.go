package main

import "fmt"

func makeBoundedFibonacciArray(maxValue int) []int {
	fibonacciArray := []int{1, 2}
	for i := 2; i < maxValue; i++ {
		currentElement := fibonacciArray[i-2] + fibonacciArray[i-1]
		if currentElement > maxValue {
			return fibonacciArray
		}
		fibonacciArray = append(fibonacciArray, currentElement)
	}
	return fibonacciArray
}

func setOddElementsToZero(myArray []int) []int {
	for i := 0; i < len(myArray); i++ {
		if myArray[i]%2 == 1 {
			myArray[i] = 0
		}
	}
	return myArray
}

func sumOfElements(myArray []int) int {
	sum := 0
	for i := 0; i < len(myArray); i++ {
		sum += myArray[i]
	}
	return sum
}

func main() {
	/*Prompt:
	  Each new term in the Fibonacci sequence is generated by adding the previous
	  two terms. By starting with 1 and 2, the first 10 terms will:
	  1, 2, 3, 5, 8, 13, 21, 34, 55, 89 ...
	  By considering the terms in the Fibonacci sequence whose values do not exceed
	  four million, find the sum of the even-valued terms.
	*/
	fmt.Println("The solution to Euler Problem 2 is: ")
	fmt.Println(sumOfElements(setOddElementsToZero(makeBoundedFibonacciArray(4000000))))
}
