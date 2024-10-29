package main

import "fmt"

func main() {
	var auth BaseAuth = Auth{
		Email:    "admin@admin.com",
		Password: "admin",
	}
	auth.login("admin@admin.com", "password")
}

type BaseAuth interface {
	login(email, password string) bool
}

type Auth struct {
	Email    string
	Password string
}

// login implements BaseAuth.
func (a Auth) login(email string, password string) bool {
	fmt.Println("Login with:", email, password)
	return true
}
