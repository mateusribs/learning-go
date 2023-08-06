/*
Write a function that calculates the average of a float64 slice.
*/

package main

import "fmt"

func get_average(values []float64) (float64) {

	var sum float64 = 0.0

	switch len(values){
	case 0:
		average := 0.0
		return average
	
	default:
		for _, value := range values {
			sum += float64(value)
		}
	
		average := sum / float64(len(values))
		return average
	}

}

func main() {

	values := []float64{2, 2, 2, 222, 3}

	average := get_average(values)

	fmt.Printf("The average is %.2f \n", average)
	
}
