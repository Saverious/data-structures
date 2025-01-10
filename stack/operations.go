package main

import "fmt"

// create stack
func InitStack() *Stack {
	nodes := []*Node{}

	stack := &Stack{
		items: nodes,
		top:   -1,
	}

	return stack
}

// add item
func (stack *Stack) Push(node *Node) {
	stack.items = append(stack.items, node)
	stack.top += 1
	fmt.Println("new node pushed")
}

// remove item
func (stack *Stack) Pop() {
	if stack.IsEmpty() {
		fmt.Println("cannot pop empty stack")
		return
	}

	last := len(stack.items) - 1
	stack.items = stack.items[:last]
	stack.top = len(stack.items) - 1
	fmt.Println("node popped")
}

// get top item
func (stack *Stack) Peek() (*Node, interface{}) {
	if stack.IsEmpty() {
		return nil, nil
	}

	return stack.items[stack.top], nil
}

// check if empty
func (stack *Stack) IsEmpty() bool {
	return stack.top < 0
}

// log stack
func (stack *Stack) Print() (int, interface{}) {
	if stack.IsEmpty() {
		fmt.Println("stack is empty")
		return stack.top, nil
	}

	fmt.Printf("\n\n*** STACK START ***\n")
	defer fmt.Printf("*** STACK END ***\n\n")
	for index, node := range stack.items {
		defer fmt.Printf("%v. %v\n", index, *node)
	}

	return stack.top, nil
}
