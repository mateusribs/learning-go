/*
Write a function that returns a function that performs a +2
 on integers. Name the function plusTwo.
 Generalize the function from above and create a plusX(x) which returns
 functions that add x to an integer.
*/

package main

import "fmt"


func main() {
	
	p2 := plus_two()
	px := plus_x(5)

	fmt.Printf("%v \n", p2(2))

	fmt.Printf("%v \n", px(2))

}

func plus_two() func(int) int {
	return func(i int) int {
		return i + 2
	}
}

func plus_x(x int) func(int) int {
	return func(i int) int {
		return i + x
	}
} 