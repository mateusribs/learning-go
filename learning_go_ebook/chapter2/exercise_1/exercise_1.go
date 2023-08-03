/*
Create a loop with the for construct.
Make it loop 10 times and print out the loop counter with the fmt package.
*/

package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Loop number: %d \n", i+1)
	}
}
