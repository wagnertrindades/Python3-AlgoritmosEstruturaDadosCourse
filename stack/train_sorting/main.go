package main

import "fmt"

type stack struct {
	stack []rune
	len   int
}

func main() {
	input_stack := stack{[]rune{'c', 'a', 'd', 'a', 'c', 'b', 'a'}, 7}
	holding_stack1 := stack{[]rune{}, 0}
	holding_stack2 := stack{[]rune{}, 0}
	output_stack := stack{[]rune{}, 0}

	fmt.Println("Input stack:", string(input_stack.stack))
	fmt.Println("Holding stack 1:", string(holding_stack1.stack))
	fmt.Println("Holding stack 2:", string(holding_stack2.stack))
	fmt.Println("Output stack:", string(output_stack.stack))
	fmt.Println("--------------------------------")

	holding_stack1.push(input_stack.pop())
	fmt.Println("Input stack:", string(input_stack.stack))
	fmt.Println("Holding stack 1:", string(holding_stack1.stack))
	fmt.Println("Holding stack 2:", string(holding_stack2.stack))
	fmt.Println("Output stack:", string(output_stack.stack))
	fmt.Println("--------------------------------")

	holding_stack2.push(input_stack.pop())
	fmt.Println("Input stack:", string(input_stack.stack))
	fmt.Println("Holding stack 1:", string(holding_stack1.stack))
	fmt.Println("Holding stack 2:", string(holding_stack2.stack))
	fmt.Println("Output stack:", string(output_stack.stack))
	fmt.Println("--------------------------------")

	output_stack.push(input_stack.pop())
	fmt.Println("Input stack:", string(input_stack.stack))
	fmt.Println("Holding stack 1:", string(holding_stack1.stack))
	fmt.Println("Holding stack 2:", string(holding_stack2.stack))
	fmt.Println("Output stack:", string(output_stack.stack))
	fmt.Println("--------------------------------")

	holding_stack1.push(input_stack.pop())
	fmt.Println("Input stack:", string(input_stack.stack))
	fmt.Println("Holding stack 1:", string(holding_stack1.stack))
	fmt.Println("Holding stack 2:", string(holding_stack2.stack))
	fmt.Println("Output stack:", string(output_stack.stack))
	fmt.Println("--------------------------------")

	holding_stack2.push(input_stack.pop())
	fmt.Println("Input stack:", string(input_stack.stack))
	fmt.Println("Holding stack 1:", string(holding_stack1.stack))
	fmt.Println("Holding stack 2:", string(holding_stack2.stack))
	fmt.Println("Output stack:", string(output_stack.stack))
	fmt.Println("--------------------------------")

	holding_stack1.push(input_stack.pop())
	fmt.Println("Input stack:", string(input_stack.stack))
	fmt.Println("Holding stack 1:", string(holding_stack1.stack))
	fmt.Println("Holding stack 2:", string(holding_stack2.stack))
	fmt.Println("Output stack:", string(output_stack.stack))
	fmt.Println("--------------------------------")

	output_stack.push(input_stack.pop())
	fmt.Println("Input stack:", string(input_stack.stack))
	fmt.Println("Holding stack 1:", string(holding_stack1.stack))
	fmt.Println("Holding stack 2:", string(holding_stack2.stack))
	fmt.Println("Output stack:", string(output_stack.stack))
	fmt.Println("--------------------------------")

	input_stack.push(holding_stack2.pop())
	fmt.Println("Input stack:", string(input_stack.stack))
	fmt.Println("Holding stack 1:", string(holding_stack1.stack))
	fmt.Println("Holding stack 2:", string(holding_stack2.stack))
	fmt.Println("Output stack:", string(output_stack.stack))
	fmt.Println("--------------------------------")

	input_stack.push(output_stack.pop())
	fmt.Println("Input stack:", string(input_stack.stack))
	fmt.Println("Holding stack 1:", string(holding_stack1.stack))
	fmt.Println("Holding stack 2:", string(holding_stack2.stack))
	fmt.Println("Output stack:", string(output_stack.stack))
	fmt.Println("--------------------------------")

	input_stack.push(output_stack.pop())
	fmt.Println("Input stack:", string(input_stack.stack))
	fmt.Println("Holding stack 1:", string(holding_stack1.stack))
	fmt.Println("Holding stack 2:", string(holding_stack2.stack))
	fmt.Println("Output stack:", string(output_stack.stack))
	fmt.Println("--------------------------------")

	input_stack.push(holding_stack2.pop())
	fmt.Println("Input stack:", string(input_stack.stack))
	fmt.Println("Holding stack 1:", string(holding_stack1.stack))
	fmt.Println("Holding stack 2:", string(holding_stack2.stack))
	fmt.Println("Output stack:", string(output_stack.stack))
	fmt.Println("--------------------------------")

	input_stack.push(holding_stack1.pop())
	fmt.Println("Input stack:", string(input_stack.stack))
	fmt.Println("Holding stack 1:", string(holding_stack1.stack))
	fmt.Println("Holding stack 2:", string(holding_stack2.stack))
	fmt.Println("Output stack:", string(output_stack.stack))
	fmt.Println("--------------------------------")

	input_stack.push(holding_stack1.pop())
	fmt.Println("Input stack:", string(input_stack.stack))
	fmt.Println("Holding stack 1:", string(holding_stack1.stack))
	fmt.Println("Holding stack 2:", string(holding_stack2.stack))
	fmt.Println("Output stack:", string(output_stack.stack))
	fmt.Println("--------------------------------")

	input_stack.push(holding_stack1.pop())
	fmt.Println("Input stack:", string(input_stack.stack))
	fmt.Println("Holding stack 1:", string(holding_stack1.stack))
	fmt.Println("Holding stack 2:", string(holding_stack2.stack))
	fmt.Println("Output stack:", string(output_stack.stack))
	fmt.Println("--------------------------------")
}

func (s *stack) push(value rune) {
	s.stack = append(s.stack, value)
	s.len++
}

func (s *stack) pop() rune {
	if s.empty() {
		panic("This stack is empty!")
	} else {
		last_position := s.len - 1
		last_position_value := s.top()
		s.stack = append(s.stack[:last_position], s.stack[last_position+1:]...)
		s.len--

		return last_position_value
	}
}

func (s stack) top() rune {
	if s.empty() {
		panic("This stack is empty!")
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
