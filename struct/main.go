package main

import "fmt"

func main() {
	user := User{
		Name: "Khadga Bahadur Shrestha",
		Role: "Software Developer",
		Age:  25,
	}
	fmt.Println("Name:", user.Name, "Role:", user.Role, "Age:", user.Age)

	rakas := User{
		"Sakar Subedi",
		"React Developer",
		28,
	}

	fmt.Println(rakas)
}

type User struct {
	Name string
	Role string
	Age  int
}