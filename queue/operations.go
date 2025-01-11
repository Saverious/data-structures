package main

import "fmt"

// initialize queue
func InitQueue() *Queue {
	queue := &Queue{
		nodes: nil,
		front: -1,
		rear:  -1,
		size:  0,
	}

	return queue
}

// view queue
func (queue *Queue) Print() {
	if queue.IsEmpty() {
		fmt.Println("Queue is empty")
		return
	}

	for index, node := range queue.nodes {
		fmt.Printf("%v. message = %v\n", index, *node)
	}
}

// check if empty
func (queue *Queue) IsEmpty() bool {
	return queue.size == 0
}

// add node
func (queue *Queue) Enqueue(node *Message) {
	queue.nodes = append(queue.nodes, node)
	if queue.IsEmpty() {
		queue.front++
	}

	queue.rear++
	queue.size++
	fmt.Println("node has been enqueued")
}

// remove node
func (queue *Queue) Dequeue() {
	if queue.IsEmpty() {
		fmt.Println("Cannot dequeue. Queue is empty")
		return
	}

	if queue.size == 1 {
		queue.front = -1
		queue.rear = -1
		queue.size = 0
	} else {
		queue.nodes = queue.nodes[queue.front+1:]
		queue.rear--
		queue.size--
	}

	fmt.Println("First node removed")
}

// select front node
func (queue *Queue) Peek() *Message {
	if queue.IsEmpty() {
		fmt.Println("Queue is empty")
		return nil
	}

	return queue.nodes[queue.front]
}

// select rear node
func (queue *Queue) Rear() *Message {
	if queue.IsEmpty() {
		fmt.Println("Queue is empty")
		return nil
	}

	return queue.nodes[queue.rear]
}
