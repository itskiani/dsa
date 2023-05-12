package main

import "fmt"

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
	size++
}

func listMenu() {
	fmt.Println("1. Insert in stack")
	fmt.Println("2. Remove in stack")
	fmt.Println("3. Display stack")
	fmt.Println("4. Exit")
	fmt.Println("Enter a number between 1 and 3:")
}

func main() {}
