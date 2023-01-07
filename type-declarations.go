package main

import "fmt"

func main() {

	// this is how to make type declaration
	// it's useful for making it aliases
	type idNumber string
	type maritalStatus bool

	var myIdNumber idNumber = "M3118035"
	var myMaritalStatus maritalStatus = false

	fmt.Println(myIdNumber)
	fmt.Println(myMaritalStatus)

}
