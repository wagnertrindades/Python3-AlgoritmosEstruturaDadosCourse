package main

import "fmt"

type stack struct {
	stack []int
	len   int
}

func main() {
	my_stack := stack{[]int{}, 0}

	my_stack.push(1)
	fmt.Println("Stack push:", my_stack)

	my_stack.push(5)
	fmt.Println("Before Stack pop:", my_stack)
	my_stack.pop()
	fmt.Println("After Stack pop: ", my_stack)

	fmt.Println("Stack top:", my_stack.top())
}

func (s *stack) push(value int) {
	s.stack = append(s.stack, value)
	s.len++
}

func (s *stack) pop() {
	if s.empty() {
		fmt.Println("The stack is empty!")
	} else {
		last_position := s.len - 1
		s.stack = append(s.stack[:last_position], s.stack[last_position+1:]...)
		s.len--
	}
}

func (s stack) top() int {
	if s.empty() {
		return s.len
	} else {
		return s.stack[s.len-1]
	}
}

func (s stack) empty() bool {
	if s.len == 0 {
		return true
	} else {
		return false
	}
}
