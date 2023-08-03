/*
Rewrite the loop again so that it fills an array and then prints that array to the screen.
*/

package main

import "fmt"

func main() {
	var arr [10]int

	for i := 0; i < 10; i++ {
		arr[i] = i
	}

	fmt.Println(arr)
}
