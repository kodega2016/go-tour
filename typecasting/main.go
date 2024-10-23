package main

import "fmt"

func main() {
	x, y := 3, 4
	var f float32 = float32(x * y)
	fmt.Printf("The type of %v is %T\n", f, f)
}
