package main

// node
type Message string

// queue
type Queue struct {
	nodes []*Message
	front int
	rear  int
	size  int
}
