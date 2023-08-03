/*
Write code to calculate the average of a float64 slice.
In a later exercise you will make it into a function.
*/

package main

import "fmt"

func main() {

	var sum float64 = 0.0

	values := []float64{2, 2, 2, 2, 2}

	switch len(values){
	case 0:
		average := 0.0
		fmt.Printf("The average is %.2f \n", average)
	
	default:
		for _, value := range values {
			sum += float64(value)
		}
	
		average := sum / float64(len(values))
	
		fmt.Printf("The average is %.2f \n", average)
	}
}
