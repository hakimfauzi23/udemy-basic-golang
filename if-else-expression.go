package main

import "fmt"

func main() {

	name := "Joko"

	if name == "Hakim" {
		fmt.Println("Hello Hakim")
	} else if name == "Joko" {
		fmt.Println("Hello Joko")
	} else {
		fmt.Println("Hi, what's is your name?")
	}

	// If with short statement
	if length := len(name); length > 5 {
		fmt.Println("Name is too long")
	} else {
		fmt.Println("The name is already correct!!")
	}

}
