package main

import "fmt"

// break is use for stops the loop
func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Loops ", i)
	}
}
