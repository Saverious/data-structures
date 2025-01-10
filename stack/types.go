package main

// node
type Node struct {
	name string
	age  int
}

// stack
type Stack struct {
	items []*Node
	top   int
}
