package main

import "fmt"

// Continue is tu stop current loop, and go to the next loop.
func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Loops :", i)
	}
}
