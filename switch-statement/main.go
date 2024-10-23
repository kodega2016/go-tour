package main

import (
	"fmt"
	"runtime"
)

func main() {
	role := "software developer"
	switch role {
	case "software developer":
		fmt.Println("You are software developer")
	default:
		fmt.Println("You are a guest")
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("You are running on macos")
	case "linux":
		fmt.Println("you are running on linux")
	default:
		fmt.Println("you are running on window")
	}
}
