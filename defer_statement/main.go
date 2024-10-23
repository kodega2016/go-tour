package main

import "fmt"

func main() {
	defer fmt.Println("closing connection...")
	fmt.Println("processing payment request...")
}
