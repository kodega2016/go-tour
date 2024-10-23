package main

import "fmt"

func main() {
	// simple variable declaration
	var name string = "Explore golang with Go Tour"
	fmt.Println(name)

	// short variable declaration
	username, role := "kodega2016", "superadmin"
	fmt.Println(username, role)

	// multiple variable declaration
	var (
		company_name    string = "Kodega Tech"
		company_address string = "Miklajung Morang"
		company_size    int    = 45
	)

	fmt.Println(company_name, company_address, company_size)

	// zero values
	var i int
	var f float64
	var s string
	var b bool
	fmt.Printf("%v %v %v %v\n", i, f, s, b)
}
