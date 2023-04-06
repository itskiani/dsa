package main

import "fmt"

type Node struct {
	data string
	next *Node
}

func listMenu() {
	fmt.Println("1. Insert in stack")
	fmt.Println("2. Remove in stack")
	fmt.Println("3. Display linkedlist")
	fmt.Println("4. Exit")
	fmt.Println("Enter a number between 1 and 3:")
}

func main() {}
