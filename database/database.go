package database

import "fmt"

var connection string

// will be executed in the very first time
func init() {
	fmt.Println("Will be executed")
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
