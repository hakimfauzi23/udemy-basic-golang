package main

import "fmt"

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

// recursive is when calling function inside the current function it self
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}
func main() {
	loop := factorialLoop(5)
	loopRecursive := factorialRecursive(5)
	fmt.Println(loop)
	fmt.Println(loopRecursive)
}
