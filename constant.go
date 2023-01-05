package main

// constant is immutable value
// should be direct assign the value
func main() {

	const firstName string = "Hanif"
	const lastName = "Hakim"
	const value = 1000

	// firstName = "Joko" // will generate error
	// lastName = "Joko"  // will generate error

	// multiple const declaration
	const (
		salary = 3100000
		job    = "Software engineer"
		skill  = "golang"
	)

}
