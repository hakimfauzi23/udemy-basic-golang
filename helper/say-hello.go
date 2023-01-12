package helper

import "fmt"

/*
	inside a package cannot have same function with same name
*/
func SayHello(name string) {
	fmt.Println("Hello", name)
}
