package main

import "fmt"

// variable argumennts use ... and only can use in the final of the argument
// it will make the function params is undeclarable or it's free to create the function
func sumAll(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}

	return total
}

func main() {
	// so the sumAll arguments is unlimited.
	total := sumAll(10, 10, 80, 80)
	total0 := sumAll()
	fmt.Println(total)
	fmt.Println(total0)

	// is it possible to make slice as variable arguments in variadic function
	// to make it happen use ... after numbers slice `numbers...`
	numbers := []int{10, 10, 10, 10}
	total2 := sumAll(numbers...)
	fmt.Println(total2)

}
