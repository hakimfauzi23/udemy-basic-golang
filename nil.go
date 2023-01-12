package main

import "fmt"

/*
Nil is same as null or empty data
*/
func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}
func main() {

	var person map[string]string = NewMap("")

	// nil is really useful for validation
	if person == nil {
		fmt.Println("Data is empty")
	} else {
		fmt.Println(person)
	}

}
