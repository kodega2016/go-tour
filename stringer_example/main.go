package main

import "fmt"

func main() {
	user := User{
		"Sakar Subedi",
		"Software Developer",
		25,
	}
	fmt.Println(user)
}

type User struct {
	Name string
	Role string
	Age  int
}

func (user User) String() string {
	return fmt.Sprintf("%s is %s and %v years old\n", user.Name, user.Role, user.Age)
}
