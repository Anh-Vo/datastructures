package main

import "fmt"

func main() {
	list := New()
	list.insert(1)
	list.insert(2)
	list.insert(3)
	list.print()
	fmt.Printf("It is [ %v ] that the list contains the value 2: \n", list.contains(2))
	fmt.Printf("Removing [ %v ]\n", list.remove(2))
	fmt.Printf("It is [ %v ] that the list contains the value 2: \n", list.contains(2))
	list.print()
}
