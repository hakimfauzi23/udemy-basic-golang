package main

import (
	"fmt"
	"strings"
)

func main() {

	words := "Hanif Fauzi Hakim"
	fmt.Println(strings.Contains(words, "Hanif"))
	fmt.Println(strings.Contains(words, "Joko"))

	fmt.Println(strings.Split(words, " "))
	fmt.Println(strings.ToLower(words))
	fmt.Println(strings.ToUpper(words))
	fmt.Println(strings.Trim("    "+words+"    ", " "))

	fmt.Println(strings.ReplaceAll(" Hanif Hanif Hanif", "Hanif", "Jhon"))

}
