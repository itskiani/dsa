package main

import (
	"fmt"
	"os"
)

type Node struct {
	data int
	next *Node
}

var head *Node = new(Node)

func createNode(val int) *Node {
	var node *Node
	node = new(Node)

	if node == nil {
		fmt.Println("not allocated")
		os.Exit(0)
	}

	node.data = val
	node.next = nil
	return node
}

func main() {
	for {
		fmt.Println("1. Insert in begin")
		fmt.Println("2. Insert in last")
		fmt.Println("3. Insert in position")
		fmt.Println("4. Delete in position")
		fmt.Println("5. Display linkedlist")
		fmt.Println("6. Exit")
		fmt.Println("Enter a number between 1 and 6:")

		var i int
		num, _ := fmt.Scanf("%d", &i)

		switch {
		case num == 1:
			os.Exit(0)
		case num == 6:
			os.Exit(0)
		}
	}

}
