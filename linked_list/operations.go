package main

import "fmt"

// initialize list
func InitLinkedList() *LinkedList {
	linkedList := &LinkedList{
		head: nil,
		size: 0,
	}

	return linkedList
}

// get size of linked list
func (linkedList *LinkedList) Size() int {
	return linkedList.size
}

// check for empty list
func (linkedList *LinkedList) IsEmpty() bool {
	return linkedList.head == nil
}

// get head
func (linkedList *LinkedList) FirstNode() *Node {
	return linkedList.head
}

// log list
func (linkedList *LinkedList) Print() {
	head := linkedList.head
	if linkedList.IsEmpty() {
		fmt.Printf("\n*** EMPTY LIST ***\n")
		return
	}

	fmt.Printf("\n*** PRINT LIST START ***\n")
	i := 1
	for head != nil {
		fmt.Printf("%v. node = %v   address = %p  next = %p\n", i, head, head, head.next)
		head = head.next
		i++
	}
	fmt.Printf("\n*** PRINT LIST STOP ***\n")
}

// insert node in front
func (linkedList *LinkedList) Prepend(node **Node) {
	(*node).next = linkedList.head
	linkedList.head = *node
	linkedList.size += 1
	fmt.Println("node prepended")
}

// insert node in the middle
func (linkedList *LinkedList) InsertAfter(index int, node **Node) {
	if linkedList.size == index {
		linkedList.Append(node)
		return
	}

	if linkedList.IsEmpty() {
		linkedList.Prepend(node)
		return
	}

	i := 1
	current := linkedList.head
	for current != nil {
		if i == index {
			(*node).next = current.next
			current.next = *node
			linkedList.size += 1
			fmt.Println("node inserted")
			return
		}

		current = current.next
		i++
	}
}

// insert node at the end
func (linkedList *LinkedList) Append(node **Node) {
	if linkedList.IsEmpty() {
		linkedList.Prepend(node)
		return
	}

	current := linkedList.head
	for current.next != nil {
		current = current.next
	}

	current.next = *node
	(*node).next = nil
	linkedList.size += 1
	fmt.Println("node appended")
}
