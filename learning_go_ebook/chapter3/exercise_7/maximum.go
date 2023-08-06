/*
Write a function that finds the maximum value in an int slice ([]int).
*/

package main

import "fmt"

func find_biggest_integer(values []int) int {
	biggest := -100000000

	for _, value := range values {
		if value > biggest {
			biggest = value
		}
	}

	return biggest
}

func main() {
	
	values := []int{-100, -2, -4, -2000, -3253}

	biggest := find_biggest_integer(values)

	fmt.Printf("The biggest integer is %d\n", biggest)

}
