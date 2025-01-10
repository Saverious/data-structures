package main

// node
type Node struct {
	country string
	next    *Node
}

// linked list
type LinkedList struct {
	head *Node
	size int
}
