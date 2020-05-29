package main

import "fmt"

// A ListNode represent a node in a singly linked list
type ListNode struct {
	data int       // value of the node
	next *ListNode // pointer to the next node in the list
}

// A LinkedList represents a singly linked list
type LinkedList struct {
	head *ListNode // A sentinel node that points to the first item in the list
	size int       // A counter that keeps track of the total number of nodes in the list
}

func (n *ListNode) add(int v) {

}

func main() {
	fmt.Println("vim-go")
}
