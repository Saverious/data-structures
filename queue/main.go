package main

import (
	"fmt"
	"os"
)

// prompt user for input
func inputPrompt() *Message {
	var msg Message
	fmt.Print("enter message to queue: ")
	fmt.Scanln(&msg)

	return &msg
}

func main() {
	queue := InitQueue()

	for {
		fmt.Printf("\n\t\t*** QUEUE DATA STRUCTURE ***\n")
		fmt.Printf("Select operation:\n1. Enqueue\n2. Dequeue\n3. Peek\n4. Get rear node\n5. Get size of queue\n6. View queue\n0. Exit\n\n")

		var option int
		fmt.Print("option: ")
		fmt.Scanln(&option)

		switch option {
		case 1:
			msg := inputPrompt()
			queue.Enqueue(msg)
		case 2:
			queue.Dequeue()
		case 3:
			msg := queue.Peek()
			if msg != nil {
				fmt.Println("peeked message: ", *msg)
			} else {
				fmt.Println("first message not found")
			}
		case 4:
			msg := queue.Rear()
			if msg != nil {
				fmt.Println("last message: ", *msg)
			} else {
				fmt.Println("last message not found")
			}
		case 5:
			fmt.Printf("Queue size: %v\n", queue.size)
		case 6:
			queue.Print()
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice! Please try again")
		}
	}
}
