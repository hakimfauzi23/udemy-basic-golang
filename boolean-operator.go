package main

import "fmt"

func main() {
	var finalValue = 85
	var presenceValue = 90

	var isPassedTest = finalValue > 80
	var isPassedPresence = presenceValue > 80

	var isPassed = isPassedTest && isPassedPresence
	fmt.Println(isPassed)

	// Real Case
	fmt.Println(finalValue >= 80 && presenceValue >= 80)
}
