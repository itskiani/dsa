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

func listMenu() {
	fmt.Println("1. Insert in begin")
	fmt.Println("2. Insert in last")
	fmt.Println("3. Insert in position")
	fmt.Println("4. Delete in position")
	fmt.Println("5. Display linkedlist")
	fmt.Println("6. Exit")
	fmt.Println("Enter a number between 1 and 6:")
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
				insertBegin()
				fmt.Println("press 'Enter' to back to menu....")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				listMenu()
			case 5:
				display(head)
				fmt.Println("press 'Enter' to back to menu....")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				listMenu()
			case 6:
				os.Exit(0)
			default:
				fmt.Println("Please enter a valid number between 1 to 6")
			}
		}
	}

}
