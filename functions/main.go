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

func greet() {
	fmt.Println("Greeting...")
}

func logMessage(message string) {
	fmt.Println("INFO:", message)
}

func getSum(a, b int) int {
	return a + b
}

func getResults(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}
