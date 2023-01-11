package main

import "fmt"

func endApp() {
	// recover function is used for catch panic data
	// and with recover will be stopped panic phase
	// in other hands the program will still executed till the end
	message := recover()
	if message != nil {
		fmt.Println("Error with message", message)
	}
	fmt.Println("End of the application")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Application Goes Error and stopped here")
	}

	fmt.Println("Application is running")
}

func main() {
	runApp(true)
	fmt.Println("Hakimm")
}
