/*
Write a function that takes a variable number of ints and print each integer on a separate line.
*/

package main

import "fmt"

func variadic_func(args ...int) {
	for _, n := range args {
		fmt.Println(n)
	}
}

func main() {
	arguments := []int{1, 2, 3, 4, 5, 6, 7, 8, 10}

	variadic_func(arguments...)
}
