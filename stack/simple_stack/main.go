package main

import "fmt"

func main() {
	stack := []int{}

	stack_push := push(stack, 1)
	fmt.Println("Stack push:", stack_push)

	stack_push_two := push(stack_push, 5)
	stack_pop := pop(stack_push_two)
	fmt.Println("Before Stack pop:", stack_push_two, "After Stack pop: ", stack_pop)

	stack_top := top(stack_push_two)
	fmt.Println("Stack top:", stack_top)

	stack_pop_empty := pop(stack)
	fmt.Println("Stack pop in empty slice:", stack_pop_empty)

	stack_top_empty := top(stack)
	fmt.Println("Stack top in empty slice:", stack_top_empty)
}

func push(stack []int, value int) []int {
	return append(stack, value)
}

func pop(stack []int) []int {
	if empty(stack) {
		return stack
	} else {
		last_position := len(stack) - 1
		return append(stack[:last_position], stack[last_position+1:]...)
	}
}

func top(stack []int) []int {
	if empty(stack) {
		return stack
	} else {
		return stack[1:]
	}
}

func empty(stack []int) bool {
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}
