package main

import "fmt"

type Node struct {
	data  int
	front *Node
	rear  *Node
	size  uint
	cap   uint
}

func isFull(queue *Node) bool {
	return (queue.size == queue.cap)
}

func isEmpty(queue *Node) bool {
	return (queue.size == 0)
}

func listMenu() {
	fmt.Println("1. Insert in queue")
	fmt.Println("2. Remove in queue")
	fmt.Println("3. Display queue")
	fmt.Println("4. Exit")
	fmt.Println("Enter a number between 1 and 4:")
}

func main() {

}
