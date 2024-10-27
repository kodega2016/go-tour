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
}
