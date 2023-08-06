/*
Write a simple map()-function in Go. It is sufficient for this function only to work for ints.
*/

package main

import "fmt"

func push (value int, stack []int) []int{
	
	if len(stack) < cap(stack) {
		stack = append(stack, value)
		return stack
	}

	fmt.Println("The stack is full")

	return stack
}

func pop(stack []int) (int, []int) {
	new_stack := make([]int, 0)
	if len(stack) > 0 {
		value := stack[len(stack) - 1]
		new_stack = append(new_stack, stack[:len(stack) - 1]...)
		return value, new_stack
	}

	fmt.Println("Stack is empty")

	return -1, stack
}

func main() {
	
	var value int
	stack := make([]int, 0, 5)

	fmt.Printf("%v\n", stack)

	stack = push(1, stack[:])
	stack = push(2, stack[:])
	stack = push(3, stack[:])
	stack = push(4, stack[:])
	stack = push(5, stack[:])
	stack = push(6, stack[:])

	fmt.Printf("%v\n", stack)

	value, stack = pop(stack)
	fmt.Printf("%v\n %d\n", stack, value)

	value, stack = pop(stack)
	fmt.Printf("%v\n %d\n", stack, value)

	value, stack = pop(stack)
	fmt.Printf("%v\n %d\n", stack, value)

	value, stack = pop(stack)
	fmt.Printf("%v\n %d\n", stack, value)

	value, stack = pop(stack)
	fmt.Printf("%v\n %d\n", stack, value)

	value, stack = pop(stack)
	fmt.Printf("%v\n %d\n", stack, value)
		

}
