package main

import "fmt"

func main() {

	// For Loop
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Loops :", counter)

	}

	// Print a slice

	// #1
	slice := [...]string{"Hanif", "Fauzi", "Hakim", "Angel", "Audri"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// #2
	// _ variable is use when you don't use index in for range
	for _, value := range slice {
		fmt.Println(value)
	}

	// Print a map
	person := make(map[string]string)
	person["name"] = "Hakim"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}

}
