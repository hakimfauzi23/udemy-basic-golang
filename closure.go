package main

import "fmt"

func main() {

	name := "InitName"
	counter := 0

	// this func cannot accessed by block code above
	increment := func() {

		name := "accidentName"
		fmt.Println("increment")

		// this counter increment is accidentally
		// but will cause counter variable will change
		counter++
		fmt.Println(name)
	}

	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)
}
