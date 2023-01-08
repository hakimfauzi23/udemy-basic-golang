package main

import "fmt"

func main() {

	name := "Fauzi"

	switch name {
	case "Hakim":
		fmt.Println("Hello Hakim")
	case "Hanif":
		fmt.Println("Hello Hanif")
	case "Joni":
		fmt.Println("Hello Joni")
	default:
		fmt.Println("What's your name?")

	}

	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Name size is too long")
	// case false:
	// 	fmt.Println("Name is already correct")
	// }

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Name size is too long")
	case length > 5:
		fmt.Println("Name is still too long")
	default:
		fmt.Println("Name is already right!")
	}
}
