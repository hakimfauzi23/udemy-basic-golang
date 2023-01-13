package main

import (
	"fmt"
	"regexp"
)

func main() {

	var regex *regexp.Regexp = regexp.MustCompile("h([a-z])([a-z])([a-z])f")

	fmt.Println(regex.MatchString("hanif"))
	fmt.Println(regex.MatchString("honef"))
	fmt.Println(regex.MatchString("Honef"))
	regex.FindAllString("hanif hakim hoblii", 1)

}
