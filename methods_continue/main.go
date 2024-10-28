package main

import "fmt"

func main() {
	f := MyFloat(10)
	f.printMyValue()
}

type MyFloat float32

func (myFloat MyFloat) printMyValue() {
	fmt.Println("My value is:", myFloat)
}
