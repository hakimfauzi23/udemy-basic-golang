package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {

	// defer will assure execution of function even there's error between method called
	defer logging()
	fmt.Println("Run application")
	result := 10 / value
	fmt.Println("Result", result)

}

func main() {
	runApplication(10)
}
