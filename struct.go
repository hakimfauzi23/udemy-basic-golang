package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// Create stuct function
func (customer Customer) sayGoodDay(name string) {

	fmt.Println("Hello", name, "Have a good day!", "My name is", customer.Name)

}

func main() {

	// // #1
	var hakim Customer
	hakim.Name = "Hakim"
	hakim.Address = "Cilacap"
	hakim.Age = 22

	// call struct function
	hakim.sayGoodDay("Joni")
	fmt.Println(hakim)

	// // #2
	// hanif := Customer{
	// 	Name:    "Hanif",
	// 	Address: "Indonesia",
	// 	Age:     21,
	// }

	// fmt.Println(hanif)

	// // #3
	// fauzi := Customer{"Fauzi", "Indonesia", 18}
	// fmt.Println(fauzi)
}
