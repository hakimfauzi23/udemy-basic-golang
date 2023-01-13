package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	utc := time.Date(2020, 19, 10, 10, 10, 23, 10, time.UTC)
	fmt.Println(utc)

	// unique layout in golang
	layout := "2006-01-02"

	parse, _ := time.Parse(layout, "1970-03-01")
	fmt.Println(parse)

}
