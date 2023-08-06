/*
Write a simple map()-function in Go. It is sufficient for this function only to work for ints.
*/

package main

import "fmt"

func map_function(f func(int) int, list []int) []int {
	mapped_list := make([]int, len(list))

	for index, value := range list {
		mapped_list[index] = f(value)
	}

	return mapped_list
}

func main() {
	
	list := []int{1, 2, 3, 4, 5, 6, 7}

	p2 := func(i int) int {
		return i + 2
	}

	mapped_list := map_function(p2, list)

	fmt.Printf("Mapped values %v\n", mapped_list)

}
