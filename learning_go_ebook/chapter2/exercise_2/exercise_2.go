/*
Rewrite the loop from 1 to use goto. The keyword for may not be used.
*/

package main

import "fmt"

func main() {
	i := 0

Loop:
	fmt.Printf("Loop number: %d \n", i + 1)
	i++
	if i < 10 {
		goto Loop
	}
}
