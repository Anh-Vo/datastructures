package main

import "fmt"

// An Element is an element of a linked list
type Element struct {
	// The value of this element
	Value interface{}

	// A pointer to the next element
	next *Element
}

// A LinkedList represents a singly linked list
type LinkedList struct {
	// A sentinel node that points to the first item in the list
	head *Element

	// A counter that keeps track of the total number of nodes in the list
	size int
}

// New returns an initialized list.
func New() *LinkedList {
	return &LinkedList{
		&Element{nil, nil},
		0,
	}
}

// Add a new node to the end of the list
func (l *LinkedList) insert(v int) {
	newElm := &Element{v, nil}
	if l.head.next == nil {
		l.head.next = newElm
		return
	}
	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = newElm
	l.size++
}

// Iterate and print each value in the list
func (l *LinkedList) print() {
	curr := l.head.next
	fmt.Print("head ->")
	for curr != nil {
		fmt.Printf(" %v", curr.Value)
		curr = curr.next
		if curr != nil {
			fmt.Print(" ->")
		}
	}
	fmt.Println()
}

func main() {
	list := New()
	list.insert(1)
	list.insert(2)
	list.insert(3)
	list.print()
}
