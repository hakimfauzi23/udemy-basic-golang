package main

import "fmt"

// alias for func(string) string
type Filter func(string) string

// this function has parameter string and function
// filter function is one of parameter of this sayHelloWithFilter function
func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func harrasFilter(name string) string {
	if name == "Dog" {
		return "..."
	} else if name == "Bitch" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Hanif", harrasFilter)
	sayHelloWithFilter("Dog", harrasFilter)
	sayHelloWithFilter("Bitch", harrasFilter)
}
