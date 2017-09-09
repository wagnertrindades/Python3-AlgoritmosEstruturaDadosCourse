package main

import "fmt"

type deque struct {
	values []string
	len int
}

func main() {
	my_deque := deque{[]string{"salsi", "fufu"}, 2}

	fmt.Println("The init of the deque:", my_deque)

	my_deque.push_front("wagner")
	fmt.Println("Push front in deque:", my_deque)

	my_deque.push_back("timMaia")
	fmt.Println("Push back in deque:", my_deque)

	removed_deque_front := my_deque.pop_front()
	fmt.Println("Pop front in deque:", my_deque, "Removed:", removed_deque_front)

	removed_deque_back := my_deque.pop_back()
	fmt.Println("Pop back in deque:", my_deque, "Removed:", removed_deque_back)
}

func (d *deque) push_front(value string) {
	d.values = append([]string{value}, d.values...)
	d.len++
}

func (d *deque) push_back(value string) {
	d.values = append(d.values, value)
	d.len++
}

func (d *deque) pop_front() string {
	if d.empty() {
		panic("This deque is empty!")
	} else {
		first_value := d.front()
		d.values = append(d.values[:0], d.values[1:]...)
		d.len--

		return first_value
	}
}

func (d *deque) pop_back() string {
	if d.empty() {
		panic("This deque is empty!")
	} else {
		last_value := d.back()
		d.values = append(d.values[:d.len - 1], d.values[d.len:]...)
		d.len--

		return last_value
	}
}

func (d deque) front() string {
	if d.empty() {
		panic("This deque is empty!")
	} else {
		return d.values[0]
	}
}

func (d deque) back() string {
	if d.empty() {
		panic("This deque is empty!")
	} else {
		return d.values[d.len - 1]
	}
}

func (d deque) empty() bool {
	if d.len == 0 {
		return true
	} else {
		return false
	}
}
