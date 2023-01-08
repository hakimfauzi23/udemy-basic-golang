package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "Hakim",
		"address": "Cilacap",
	}

	// add data
	person["title"] = "Software Engineer"

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Learning Go-Lang"
	book["author"] = "PZN"
	book["deleted"] = "false data"

	// deleted map element
	delete(book, "deleted")
	fmt.Println(book)
}
