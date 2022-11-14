package main

import (
	"fmt"
)

type Customer struct {
	id        int
	name      string
	role      string
	email     string
	phone     string
	contacted bool
}

func main() {
	fmt.Println("Hello World!")
}
