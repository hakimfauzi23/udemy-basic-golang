package helper

import "fmt"

/*
	inside a package cannot have same function with same name
*/

/*
	Access modifier
	- if function/variable name with uppercase it's can be accessed outside package
	- if function/variable name with lowercase it's can't be accessed outside package
*/

// can be accessed outside package
func SayHello(name string) {
	fmt.Println("Hello", name)
}

// can't be accessed outside package
func sayGoodBye(name string) {
	fmt.Println("Byee!", name)
}
