package main

import "fmt"

func random() interface{} {
	return 12
}

func main() {
	var result interface{} = random()

	// // must sure if blank interface return string data
	// var resultString string = result.(string)

	// // var resultBoolean bool = result.(bool) // will go panic
	// fmt.Println(resultString)

	// handle type assertsion with switch
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is integer")
	default:
		fmt.Println("Unknown type")
	}

}
