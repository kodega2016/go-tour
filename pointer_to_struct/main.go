package main

import "fmt"

func main() {
	u := User{Name: "John", Age: 25}
	// pointer to the struct
	p := &u
	// chaning the value of the struct field
	p.Name = "Doe"
	fmt.Println(u)
}

type User struct {
	Name string
	Age  int
}
