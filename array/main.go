package main

import "fmt"

func main() {
	var members [2]string
	members[0] = "Shree Ram"
	members[1] = "Jack"
	// Out of bound execption will occur
	// members[2] = "John Shon"
	fmt.Println(members)

	primes := [3]int{
		2, 3, 5,
	}
	fmt.Println("primes:", primes)
	fmt.Println("total primes:", len(primes))
}
