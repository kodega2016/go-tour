package main

import "fmt"

func main() {
	c := Counter{
		count: 2,
	}
	c.increment()
	fmt.Println("count:", c.count)
}

type Counter struct {
	count int
}

func (c *Counter) increment() {
	c.count++
}
