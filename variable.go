package main

import "fmt"

func main() {
	var name string // this variable must be filled with string

	name = "Hanif"
	fmt.Println(name)

	name = "Fauzi" // this will change name variable above
	fmt.Println(name)

	var friendName = "Joko" // data type declare is optional, go will detect the data type.
	fmt.Println(friendName)

	var age = 30
	fmt.Println(age)

	// another var declaration
	country := "Indonesia"
	fmt.Println(country)

	// multiple var declaration
	var (
		firstName = "Hakim"
		lastName  = "Hanif"
	)

	fmt.Println(firstName + " " + lastName)
}
