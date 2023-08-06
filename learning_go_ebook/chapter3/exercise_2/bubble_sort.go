/*
Write a function that performs a bubble sort on a slice of ints.
*/

package main

import "fmt"

func bubble_sort_1(values []int) []int {

	for {
		swapped := false
		for i := 1; i <= len(values)-1; i++ {
			if values[i-1] > values[i] {
				values[i], values[i-1] = values[i-1], values[i]
				swapped = true
			}
		}

		if !swapped {
			return values
		}
	}
}

func bubble_sort_2(values []int) {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] > values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
}

func main() {

	values := []int{4, 2, 3, 3, 2, 1}

	fmt.Printf("Original sequence: %v \n", values)

	bubble_sort_2(values)

	fmt.Printf("Sorted sequence: %v \n", values)

}
