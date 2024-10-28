package main

import "fmt"

func main() {
	user := User{Name: "John", Age: 25}
	user.logUser()
}

type User struct {
	Name string
	Age  int
}

func (user User) logUser() {
	fmt.Println("Name:", user.Name)
}
