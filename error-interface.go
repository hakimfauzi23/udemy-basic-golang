package main

import (
	"errors"
	"fmt"
)

func Divider(value int, divider int) (int, error) {
	if divider == 0 {
		return 0, errors.New("Divider can't be zero")
	} else {
		result := value / divider
		return result, nil
	}
}

func main() {

	result, err := Divider(100, 0)
	if err == nil {
		fmt.Println("Result:", result)
	} else {
		fmt.Println("Error:", err.Error())
	}
}
