package main

import (
	"fmt"

	singly "github.com/anhvoo/datastructures/linkedlist/singlylist"
)

func main() {
	list := singly.New()
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Print()
	fmt.Printf("It is [ %v ] that the list contains the value 2: \n", list.Contains(2))
	fmt.Printf("Removing [ %v ]\n", list.Remove(2))
	fmt.Printf("It is [ %v ] that the list contains the value 2: \n", list.Contains(2))
	list.Print()
}
