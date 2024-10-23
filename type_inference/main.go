package main

import "fmt"

func main() {
	username := "kodega2016"
	// here we cannot assign the number type to the username because it has
	// string type inference from the previous statement
	// username = 10
	fmt.Println(username)
}
