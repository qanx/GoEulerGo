package main

import "fmt"

func generatePythagoreanTriples(limit int) [][]int {
	sliceOfPythagoreanTriples := [][]int{{3, 4, 5}}
	for q := 1; q <= limit; q++ {
		for p := 1; p <= limit; p++ {
			a := 2 * p * q
			b := p*p - q*q
			c := p*p + q*q
			if a*a+b*b == c*c && a > 0 && b > 0 && c > 0 && a < b && b < c {
				sliceToAppend := [][]int{{a, b, c}}
				sliceOfPythagoreanTriples = append(sliceOfPythagoreanTriples, sliceToAppend...)
			}
		}
	}
	return sliceOfPythagoreanTriples
}

func findPythagoreanTripleWithSumOf(value int, triplets [][]int) []int {
	for i := 0; i < len(triplets); i++ {
		triplet := triplets[i]
		if triplet[0]+triplet[1]+triplet[2] == 1000 {
			return triplet
		}
	}
	return []int{3, 4, 5}
}

func main() {
	/* Prompt:
	A Pythagorean triplet is a set of three natural numbers, a < b < c, for which
	a^2 + b^2 = c^2
	For example:
	3^2 + 4^2 = 9 + 16 = 25 = 5^2
	There exists exactly one Pythagorean triplet for which a + b + c = 1000.
	Find the product abc.
	*/

	/* Formula for generating a, b, c from http://www.friesian.com/pythag.htm:
		   a = 2pq
		   b = p^2 - q^2
		   c = p^2 + q^2

		   For some integers p and q. Since b and c are the largest, we can bind these
	     to 500:
	     p^2 - q^2 < 500; p^2 + q^2 < 500; p> 0; q > 0
	     Therefore p < 22 and q < 22 (limited to integers).
	*/

	fmt.Println("The solution to Euler Problem 9 is: ")
	sliceOfPythagoreanTriples := generatePythagoreanTriples(22)
	tripletSliceOfInterest := findPythagoreanTripleWithSumOf(1000, sliceOfPythagoreanTriples)
	fmt.Println(tripletSliceOfInterest[0] * tripletSliceOfInterest[1] * tripletSliceOfInterest[2])
}
