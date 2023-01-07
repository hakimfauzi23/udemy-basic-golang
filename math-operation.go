package main

import "fmt"

func main() {
	//Basically it's similiar to other programming language
	var result = 10 + 10
	fmt.Println(result)

	var a = 10
	var b = 10
	var c = a * b
	fmt.Println(c)

	// augmented assignments
	var i = 10
	i += 10 // i = i+10
	fmt.Println(i)

	// unary operator
	// is operation that modified one variable
	i++ // i = i+1
	var negative = -99

	fmt.Println(negative)
}
