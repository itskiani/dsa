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

var head *Node

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

func insertBegin() {
	var (
		n, prev *Node
		val     int
	)

	fmt.Print("value: ")
	fmt.Scan(&val)

	n = createNode(val)

	if head == nil {
		head = n
		head.next = nil
	} else {
		prev = head
		head = n
		head.next = prev
	}

}

func insertLast() {
	var (
		n, h *Node
		val  int
	)

	fmt.Print("value: ")
	fmt.Scan(&val)

	n = createNode(val)
	h = head

	for h.next != nil {
		h = h.next
	}

	n.next = nil
	h.next = n
}

func insertInPosition() {
	var (
		n, s, p         *Node
		val, pos, count int
	)

	fmt.Print("data: ")
	fmt.Scan(&val)

	n = createNode(val)
	fmt.Print("position: ")
	fmt.Scan(&pos)

	var i int
	s = head

	for s != nil {
		s = s.next
		count++
	}

	if pos == 1 {
		if head == nil {
			head = n
			head.next = nil
		} else {
			p = head
			head = n
			head.next = p
		}
	} else if pos > 1 && pos <= count {
		s = head

		for i = 1; i < pos; i++ {
			p = s
			s = s.next
		}

		p.next = n
		n.next = s
	} else {
		fmt.Println("Position out of range.")
	}
}

func deleteInPosition() {
	var (
		s, prev       *Node
		pos, i, count int
	)

	if head == nil {
		return
	}

	fmt.Print("position: ")
	fmt.Scan(&pos)

	s = head
	if pos == 1 {
		head = s.next
	} else {
		for s != nil {
			s = s.next
			count++
		}

		if pos >= 2 && pos <= count {
			s = head
			for i = 1; i <= pos-1; i++ {
				prev = s
				s = s.next
			}
			prev.next = s.next
		} else {
			fmt.Println("out of range.")
		}
	}
}

func display(node *Node) {
	s := node

	fmt.Println("Elements: ")
	for {
		if s == nil {
			break
		}
		fmt.Printf("%d->", s.data)
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
				deleteInPosition()
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
