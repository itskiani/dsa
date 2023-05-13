package main

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

func main() {

}
