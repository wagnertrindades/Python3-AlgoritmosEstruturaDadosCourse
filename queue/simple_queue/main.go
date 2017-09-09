package main

import "fmt"

type queue struct {
	values []string
	len		int
}

func main() {
	my_queue := queue{[]string{}, 0}

	my_queue.push("wagner")
	fmt.Println("Push in queue:", my_queue)

	my_queue.push("sonny boy")
	fmt.Println("Push in queue:", my_queue)

	pop_in_queue := my_queue.pop()
	fmt.Println("Pop in queue:", my_queue, "Value removed:", pop_in_queue)
}

func (q *queue) push(value string) {
	q.values = append(q.values, value)
	q.len++
}

func (q *queue) pop() string {
	if q.empty() {
		panic("This queue is empty!")
	} else {
		first_value := q.front()
		q.values = append(q.values[:0], q.values[1:]...)
		q.len--

		return first_value
	}
}

func (q queue) front() string {
	if q.empty() {
		panic("This queue is empty!")
	} else {
		return q.values[0]
	}
}

func (q queue) empty() bool {
	if q.len == 0 {
		return true
	} else {
		return false
	}
}
