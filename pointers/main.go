package main

import "fmt"

func main() {
	name := "Khadga Bahadur Shrestha"
	ptr := &name
	fmt.Println("Address:", ptr)
	fmt.Println("Value:", *ptr)

	*ptr = "Jack Ma"
	fmt.Println(name)
}
