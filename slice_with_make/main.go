package main

import "fmt"

func main() {
	users := make([]string, 4, 6)
	// adding values to the slices
	users = append(users, "Roshan", "Samir")
	fmt.Println(users)
	fmt.Println("length:", len(users))
	fmt.Println("Capacity:", cap(users))
}
