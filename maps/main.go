package main

import "fmt"

func main() {
	users := make(map[string]User)

	users["kodega2016"] = User{
		"Khadga Shrestha",
		27,
	}

	users["rakas"] = User{
		Name: "Sakar Subedi",
		Age:  29,
	}

	fmt.Println(users)

	for key, value := range users {
		fmt.Println(key, value)
	}
}

type User struct {
	Name string
	Age  int
}
