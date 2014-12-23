package main

import (
	"fmt"
	"sort"
	"strconv"
)

func reverseString(valueAsString string) string {
	valueAsStringReversed := ""
	for i := len(valueAsString) - 1; i >= 0; i-- {
		valueAsStringReversed = valueAsStringReversed + string([]rune(valueAsString)[i])
	}
	return valueAsStringReversed
}

func integerIsAPalindrome(value int) bool {
	valueAsString := strconv.Itoa(value)
	return valueAsString == reverseString(valueAsString)
}

func main() {
	/* Prompt:
	A palindromic number reads the same both ways. The largest palindrome made
	from the product of two 2-digit numbers is 9009 = 91 × 99
	Find the largest palindrome made from the product of two 3-digit numbers.
	*/
	fmt.Println("The solution to Euler Problem 4 is: ")
	palindromeList := []int{}
	for i := 999; i >= 100; i-- {
		for j := 999; j >= 100; j-- {
			product := i * j
			if integerIsAPalindrome(product) {
				palindromeList = append(palindromeList, product)
			}
		}
	}
	sort.Ints(palindromeList)
	fmt.Println(palindromeList[len(palindromeList)-1])
}
