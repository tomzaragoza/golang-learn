package main
import "fmt"

type Node struct {
	value float64
	next *Node
}

func all_values(head *Node) {
	for head != nil {
		fmt.Println(head.value)
		head = head.next
	}
}

func main() {
	head := new(Node)
	node1 := new(Node)
	node2 := new(Node)

	head.next = node1
	node1.next = node2

	head.value = 1
	node1.value = 2
	node2.value = 3

	all_values(head)
}