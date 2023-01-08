package main

import "fmt"

func main() {

	// array length after declare is cannot be changed!
	var names [3]string

	names[0] = "Hanif"
	names[1] = "Fauzi"
	names[2] = "Hakim"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// direct declaration
	var values = [3]int{
		90,
		95,
		80,
	}

	fmt.Println(values)

	// some function in array for access or modify array
	fmt.Println(len(names)) // for get array length even it empty
	fmt.Println(len(values))

}
