package main

import "fmt"

type node struct {
	label string
	next *node
}

type linkedList struct {
	first node
	len int
}

func main() {
	fmt.Println("---- Node ----")

	my_node := node{"Angra", nil}

	label := my_node.getLabel()
	fmt.Println("My label is:", label)

	my_node.setLabel("Dr. Sin")
	fmt.Println("After setLabel:", my_node.label)

	nextNode := my_node.getNext()
	fmt.Println("Next node is:", nextNode)

	the_second_node := node{"Iron Maiden", nil}
	my_node.setNext(the_second_node)
	fmt.Println("My first node before set next:", my_node)
	fmt.Println("My first node after set next:", my_node, my_node.next)

	fmt.Println("---- Linked List ----")

	my_linked_list := linkedList{}
	fmt.Println("Init my linked list:", my_linked_list)

	my_linked_list.add("Symbols", 0)
	fmt.Println("Add first node:", my_linked_list)
	fmt.Println("First:", my_linked_list.first, "Lenght:", my_linked_list.len)

	fmt.Println("-----------")

	my_linked_list.add("Angra", 0)
	fmt.Println("Add second node:", my_linked_list)
	fmt.Println("First:", my_linked_list.first, "Lenght:", my_linked_list.len)
	fmt.Println("Get next:", my_linked_list.first.getNext())

	fmt.Println("-----------")

	my_linked_list.add("Guns n' Roses", 2)
	fmt.Println("Add third node:", my_linked_list)
	fmt.Println("First:", my_linked_list.first, "Lenght:", my_linked_list.len)
	fmt.Println("Get next:", my_linked_list.first.getNext())

	fmt.Println("-----------")

	my_linked_list.add("Helloween", 2)
	fmt.Println("Add fourth node:", my_linked_list)
	fmt.Println("First:", my_linked_list.first, "Lenght:", my_linked_list.len)

	my_linked_list.show()

	fmt.Println("-----------")

	my_linked_list.remove(2)
	fmt.Println("Remove first node:", my_linked_list)
	fmt.Println("First:", my_linked_list.first, "Lenght:", my_linked_list.len)

	my_linked_list.show()
}

func (n node) getLabel() string {
	return n.label
}

func (n *node) setLabel(value string) {
	n.label = value
}

func (n node) getNext() node {
	if n.next == nil {
		return node{}
	} else {
		return *n.next
	}
}

func (n *node) setNext(value node) {
	n.next = &value
}

func (list *linkedList) add(label string, index int) {

	if index >= 0 {
		list_node := node{label, nil}

		if	list.empty() {
			list.first = list_node
		} else {
			if index == 0 {
				// insert in first
				list_node.setNext(list.first)
				list.first = list_node
			} else {
				prev_node := list.first
				curr_node := list.first.getNext()
				curr_index := 1

				for curr_node.next != nil {

					if curr_index == index {
						list_node.setNext(curr_node)
						prev_node.setNext(list_node)
						break
					}

					prev_node = curr_node
					curr_node = curr_node.getNext()
					curr_index++
				}
			}
		}
		list.len++
	}
}

func (list linkedList) remove(index int) {

	if !list.empty() && index >= 0 && index < list.len {
		flag_remove := false

		if list.first.next == nil {
			list.first = node{}
			flag_remove = true
		} else if index == 0 {
			list.first = list.first.getNext()
			flag_remove = true
		} else {
			prev_node := list.first
			curr_node := list.first.getNext()
			curr_index := 1

			for curr_node.next != nil {
				if index == curr_index {
					next_node := curr_node.getNext()
					prev_node.setNext(next_node)
					curr_node.setNext(node{})
					flag_remove = true
					break
				}

				prev_node = curr_node
				curr_node = curr_node.getNext()
				curr_index++
			}
		}

		if flag_remove == true {
			list.len--
		}
	}
}

func (list linkedList) empty() bool {
	if list.len == 0 {
		return true
	} else {
		return false
	}
}

func(list linkedList) show() {
	curr_node := list.first

	for curr_node.next != nil {
		fmt.Println(curr_node.getLabel())
		curr_node = curr_node.getNext()
		fmt.Println("Final:", curr_node)
	}

}
