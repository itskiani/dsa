package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	val int
}
type Queue struct {
	nodes []*Node
	front int
	rear  int
	size  int
	cap   int
}

func isFull(queue *Queue) bool {
	return queue.size == queue.cap
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
				// push()
				fmt.Println("\n press 'Enter' to back to menu....")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				listMenu()
			case 2:
				// pop()
				fmt.Println("\n press 'Enter' to back to menu....")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				listMenu()
			case 3:
				// display()
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
