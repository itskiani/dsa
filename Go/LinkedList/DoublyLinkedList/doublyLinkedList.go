package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
	prev *Node
}

var head *Node

func createNode(val int) *Node {
	var node, n *Node

	node = new(Node)
	node.data = val
	node.next = nil

	if head == nil {
		node.prev = nil
		head = node
	} else {
		n = head

		for n.next != nil {
			n = n.next
		}
		n.next = node
		node.prev = n
	}

	return node
}

func insertBegin(val int) {
	if head == nil {
		fmt.Println("First create the list.")
		return
	}

	var node *Node
	node = new(Node)
	node.data = val
	node.prev = nil
	node.next = head

	head.prev = node
	head = node
}

func listMenu() {
	fmt.Println("1. Insert in begin")
	fmt.Println("2. Insert in last")
	fmt.Println("3. Insert in position")
	fmt.Println("4. Delete in position")
	fmt.Println("5. Display linkedlist")
	fmt.Println("6. Exit")
	fmt.Println("Enter a number between 1 and 6:")
}

func main() {}
