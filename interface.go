package main

import "fmt"

/*
This is interface
in golang interface is used for contract
all function that has simliar as the interface
will automatically consider it as interface too.
*/
type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {

	var hanif Person
	hanif.Name = "Hanif"

	SayHello(hanif)

	cat := Animal{
		Name: "Meowt",
	}

	SayHello(cat)
}
