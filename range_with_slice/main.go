package main

import "fmt"

func main() {
	users := []string{
		"Khadga", "Sakar", "Sabin",
	}

	for i, user := range users {
		fmt.Println(i, user)
	}
}
