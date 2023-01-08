package main

import "fmt"

func main() {

	var months = [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// if modify slice, data in array will also change.
	// months[5] = "ChangedMonth"
	// fmt.Println(slice1)

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Hanif")
	fmt.Println(slice3)
	slice3[1] = "Not Desember"
	fmt.Println(slice3)
	fmt.Println(slice2) // will not be modified because former slice already out of cap, it will make new slice.

	// create new slice without refference it to array
	newSlice := make([]string, 3, 5) // len:2 capacity 5
	var newSlice2 = append(newSlice, "Hakim")

	newSlice[0] = "Hanif"
	newSlice[1] = "Fauzi"

	fmt.Println(newSlice)
	fmt.Println(newSlice2)

	// copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	// array vs slice
	thisArray := [5]int{1, 2, 3, 4, 5}
	thisSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(thisArray)
	fmt.Println(thisSlice)
}
