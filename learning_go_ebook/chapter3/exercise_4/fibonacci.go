/*
The Fibonacci sequence starts as follows: 1,1,2,3,5,8,13,…
Write a function that takes an int value and gives that many terms of the Fibonacci sequence
*/

package main

import "fmt"

func fibonacci(n int) []int {

	sequence := make([]int, n)

	for i := 0; i < n; i++ {

		if i < 2 {
			sequence[i] = 1
		} else {
			sequence[i] = sequence[i-1] + sequence[i-2]
		}
	}

	return sequence
}

func main() {

	n := 8

	fib_sequence := fibonacci(n)

	fmt.Printf("%v \n", fib_sequence)
}
