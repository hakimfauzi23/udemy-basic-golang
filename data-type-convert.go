package main

import "fmt"

func main() {
	var integer32 int32 = 10000
	var integer64 int64 = int64(integer32) // this will convert integer32 to int64
	var integer16 int16 = int16(integer32) // this will convert integer32 to int16

	fmt.Println(integer32)
	fmt.Println(integer64)
	fmt.Println(integer16)

	var name = "Hakim"          // this is string data type
	var charH = name[0]         // this is byte data type
	var hString = string(charH) // convert byte to string so it readable

	fmt.Println(charH)
	fmt.Println(hString)
}
