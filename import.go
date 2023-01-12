package main

// if want to access function from another package
import "basic-golang-udemy/helper"

func main() {
	helper.SayHello("Hanif")
	// helper.sayGoodBye("Hanif") // cannot be used because of access modifier

}
