package main

import "fmt"

func main() {
	// Compare two variable and result bool value

	var name1 = "Hakim"
	var name2 = "Hakim"

	var result1 = name1 == name2
	fmt.Println(result1) // true

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
