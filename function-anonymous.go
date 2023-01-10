package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("welcome!", name)

	}
}

func main() {

	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("joko", blacklist)

	// this function has a paramater that has anonymous function because
	// the function don't have a name
	registerUser("root", func(name string) bool {
		return name == "root"
	})
}
