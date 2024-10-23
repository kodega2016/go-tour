package main

import "fmt"

const PI = 3.14

func main() {
	// PI = 12
	// Here we cannot re-assign the new value to the PI because it is declared
	// as a constant
	fmt.Println("Value of PI:", PI)
}
