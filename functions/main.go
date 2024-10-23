package main

import "fmt"

func main() {
	greet()
	logMessage("Welcome to Golang Tour")
	sum := getSum(10, 20)
	fmt.Println("The sum is:", sum)

	s, d := getResults(10, 40)
	fmt.Println(s, d)
}

// simple function
func greet() {
	fmt.Println("Greeting...")
}

// function with parameter
func logMessage(message string) {
	fmt.Println("INFO:", message)
}

// function with return type
func getSum(a, b int) int {
	return a + b
}

// function with multiple return types with named return values
func getResults(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}
