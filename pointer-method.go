package main

import "fmt"

type Man struct {
	Name string
}

// create function/method with *
// add * operator right before struct data type
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	hakim := Man{"Hakim"}
	hakim.Married()

	fmt.Println(hakim.Name)
}

// Using pointer
