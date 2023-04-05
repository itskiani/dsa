package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	data int
	next *Node
	prev *Node
}

var head *Node

func createNode(val int) *Node {
	var node *Node

	node = new(Node)

	node.data = val
	node.next = nil
	node.prev = nil

	head = node

	return node
}

func insertBegin() {
	var val int
	fmt.Print("Value: ")
	fmt.Scan(&val)

	if head == nil {
		createNode(val)
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

func insertLast() {
	var (
		n, h *Node
		val  int
	)

	fmt.Print("value: ")
	fmt.Scan(&val)

	if head == nil {
		n = createNode(val)
		return
	} else {
		n = new(Node)
		h = head

		for h.next != nil {
			h = h.next
		}

		h.next = n
		n.next = nil
		n.data = val
		n.prev = h
	}

}

func insertInPosition() {
	var (
		n, p     *Node
		val, pos int
	)

	fmt.Print("Value: ")
	fmt.Scan(&val)
	fmt.Print("Position: ")
	fmt.Scan(&pos)

	if head == nil {
		return
	}

	p = head
	for i := 0; i < pos-1; i++ {
		p = p.next
		if p == nil {
			fmt.Println("Sorry, the position you entered is wrong.")
		}
	}

	n = new(Node)
	n.data = val
	if p.next == nil {
		p.next = n
		n.next = nil
		n.prev = p
	} else {
		n.next = p.next
		p.next.prev = n
		p.next = n
		n.prev = p
	}
}

func display(node *Node) {
	s := node

	fmt.Println("Elements: ")
	for {
		if s == nil {
			break
		}
		fmt.Printf("%d<->", s.data)
		s = s.next
	}
	fmt.Println("NULL")
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
			case 2:
				insertLast()
				fmt.Println("press 'Enter' to back to menu....")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				listMenu()
			case 3:
				insertInPosition()
				fmt.Println("press 'Enter' to back to menu....")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				listMenu()
			case 4:
				// deleteInPosition()
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
