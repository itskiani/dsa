package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	data int
	next *Node
}

var head *Node = nil

func push() {
	var val int
	fmt.Print("Value: ")
	fmt.Scan(&val)

	if head == nil {
		head = new(Node)
		head.data = val
	} else {
		var node *Node = new(Node)
		node.data = val
		node.next = head
		head = node
	}
}

func pop() {

	if head == nil {
		fmt.Println("Stack is empty :|")
	}

	val := head
	head = head.next
	val.next = nil

}

func display(node *Node) {
	s := node

	fmt.Println("\n Elements: ")
	for {
		if s == nil {
			break
		}
		fmt.Printf("%d \n", s.data)
		s = s.next
	}
}

func listMenu() {
	fmt.Println("1. Insert in stack")
	fmt.Println("2. Remove in stack")
	fmt.Println("3. Display stack")
	fmt.Println("4. Exit")
	fmt.Println("Enter a number between 1 and 4:")
}

func main() {
	var num int
	listMenu()

	for {
		fmt.Scan(&num)
		if num == 0 {
			fmt.Println("Thank You, exiting.....")
			break
		} else {
			switch num {
			case 1:
				push()
				fmt.Println("\n press 'Enter' to back to menu....")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				listMenu()
			case 2:
				pop()
				fmt.Println("\n press 'Enter' to back to menu....")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				listMenu()
			case 3:
				display(head)
				fmt.Println("\n press 'Enter' to back to menu....")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				listMenu()
			case 4:
				os.Exit(0)
			default:
				fmt.Println("Please enter a valid number between 1 to 4")
			}
		}
	}
}
