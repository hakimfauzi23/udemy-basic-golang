package main

import "fmt"

// this funtion will be default run
func main() {

	for i := 0; i < 10; i++ {
		sayHello()
		sayHelloTo("Hanif", "Fauzi")
	}

	// a return with return value must be called inside a variable
	returnFunc := getHello("")
	fmt.Println(returnFunc)

	// for get data from function with  multiple value return.\
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
}

// function declaration without parameter
func sayHello() {
	fmt.Println("Hello this is a function!!")
}

// function declaration with parameter
func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello ", firstName, lastName)
}

// function declaration with parameter and return value
// function with return value have to assign to variable
// function that declared with string must be return string too!
func getHello(name string) string {
	if name == "" {
		return "Hello Bro!!"
	} else {
		return "Hello " + name
	}
}

// function also possible to return multiple value!
func getFullName() (string, string) {
	return "Hanif", "Fauzi"
}
