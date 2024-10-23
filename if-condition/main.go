package main

import "fmt"

func main() {
	age := 10
	if age < 18 {
		fmt.Println("You are under 18")
	} else {
		fmt.Println("You are welcome to this application")
	}

	if result := 10 + 2; result < 18 {
		fmt.Println("Result is less than 18")
	}
}
