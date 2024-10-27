package main

import "fmt"

func main() {
	var skills []string
	skills = append(skills, "Flutter", "Golang", "Docker", "Kubernetes")
	fmt.Println(skills)

	users := []string{
		"Sakar", "Sabin", "Ashok",
	}

	users = append(users, "Jack")
	fmt.Println(users)
	fmt.Println(users[:2])
	new_users := users[:2]
	new_users[0] = "Kumari"
	fmt.Println("users:", users)

	fmt.Println("first two elements", users[:2])
	fmt.Println("last two elements", users[2:])
}
