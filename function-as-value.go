package main

import "fmt"

func getGoodBye(name string) string {
	return "Good By " + name
}

func main() {

	// declare sayGoodBye as getGoodBye container
	sayGoodBye := getGoodBye
	result := sayGoodBye("Hakim")

	fmt.Println(result)
}
