/*
* Implement linked list
 */

package main

import (
	"fmt"
	"os"
)

// prompt user for input
func inputPrompt() *Node {
	var country string
	fmt.Printf("\n\t*** add node data ***\n")
	fmt.Print("Country: ")
	fmt.Scanln(&country)
	node := &Node{country: country, next: nil}

	return node
}

func main() {
	linkedList := InitLinkedList()
	var option int

	for {
		fmt.Printf("\n\t\t\t*** LINKED LIST ***\n")
		fmt.Printf("Select operation:\n1. Insert in front\n2. Insert after specific node\n3. Insert at the end\n4. Print linked list\n5. Get linked list size\n0. Exit\n\n")

		fmt.Print("option: ")
		fmt.Scanln(&option)

		switch option {
		case 1:
			node := inputPrompt()
			linkedList.Prepend(&node)
		case 2:
			node := inputPrompt()

			var index int
			fmt.Print("insert after node number?: ")
			fmt.Scanln(&index)

			size := linkedList.size
			if size < index {
				fmt.Printf("Out of bounds! Nodes are less than %v\n", index)
				continue
			}

			linkedList.InsertAfter(index, &node)
		case 3:
			node := inputPrompt()
			linkedList.Append(&node)
		case 4:
			linkedList.Print()
		case 5:
			size := linkedList.Size()
			fmt.Printf("linked list size: %v\n", size)
		case 0:
			os.Exit(0)
		default:
			fmt.Println("invalid operation. Please try again")
		}
	}
}
