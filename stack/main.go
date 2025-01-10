/*
*	Implement stack data structure using an array
 */

package main

import (
	"fmt"
	"os"
)

// main
func main() {
	stack := InitStack()
	var option int

	for {
		fmt.Printf("\n\t\t\t*** THE STACK DATA STRUCTURE ***\n\n")
		fmt.Printf("Select operation:\n1. Push\n2. Pop\n3. Peek\n4. Print stack\n0. Exit\n\n")

		fmt.Print("option: ")
		fmt.Scanln(&option)

		switch option {
		case 1:
			var name string
			var age int

			fmt.Println("\n\t*** add node ***")
			fmt.Print("name: ")
			fmt.Scanln(&name)
			fmt.Print("age: ")
			fmt.Scanln(&age)

			node := &Node{name: name, age: age}
			stack.Push(node)
		case 2:
			stack.Pop()
		case 3:
			node, _ := stack.Peek()
			if node == nil {
				fmt.Println("cannot peek empty stack")
			} else {
				fmt.Println("Top node: ", *node)
			}
		case 4:
			stack.Print()
		case 0:
			fmt.Println("exiting...")
			os.Exit(0)
		default:
			fmt.Println("invalid operation. Please try again")
		}
	}
}
