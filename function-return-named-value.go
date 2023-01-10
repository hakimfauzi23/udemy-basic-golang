package main

// this function named each variable in start and then it can be returned withoout
// returned the declared variable
getFullName2()(fistName string, middleName string , lastName string) {
	firstName = "Hanif"
	middleName = "Fauzi"
	lastName = "Hakim"
	return
}

func main() {
	a,b,c := getFullName2()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
