package main

import (
	"fmt"
	"strconv"
)

func main() {

	// string to bool
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	// string to int
	number, err := strconv.ParseInt("1000000", 10, 64)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	// int to string
	value := strconv.FormatInt(1000000, 10)
	fmt.Println(value)

}
