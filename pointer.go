package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	/*
		All variables in golang are pass by value not reference
		So golang automaticaly duplicate a variable1 into variable2
		And it caused variable2 and variable1 are separated variable
		So the operation on variable2 will not amend variable1
	*/

	address1 := Address{"Cilacap", "Middle Java", "Indonesia"}
	address2 := address1

	address2.City = "Surakarta"

	fmt.Println(address1) // Will not change
	fmt.Println(address2)

	/*
		So there's `Pointer` in golang
		this possible to make anything in golang passed by reference
		to make this pointer use & operator
	*/

	address3 := &address1
	address3.City = "Semarang"
	fmt.Println(address1) // Will change
	fmt.Println(address2) // Will Not change
	fmt.Println(address3)

	/*
		If use & only, other variables that has same reference but undirectly called will not changed
		if want to changed all variables that has same reference use * operator
	*/

	*address3 = Address{"Surabaya", "East Java", "Indonesia"}
	fmt.Println(address1) // Will change
	fmt.Println(address2) // Will change
	fmt.Println(address3)
}
