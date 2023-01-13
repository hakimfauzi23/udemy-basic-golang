package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Hanif")
	data.PushBack("Fauzi")
	data.PushBack("Hakim")
	data.PushFront("John")

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)

	// Iterate double linked list
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	// Reversed iterate double linked list
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
}
