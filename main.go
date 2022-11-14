package main

import (
	"fmt"
	"net/http"
)

type Customer struct {
	name      string
	role      string
	email     string
	phone     string
	contacted bool
}

var customersList = map[int]Customer{
	1: Customer{
		name:      "Andrew",
		role:      "User",
		email:     "andrew@test.com",
		phone:     "000000001",
		contacted: false,
	},
	2: Customer{
		name:      "Brian",
		role:      "User",
		email:     "brian@test.com",
		phone:     "000000002",
		contacted: false,
	},
	3: Customer{
		name:      "Carla",
		role:      "User",
		email:     "carla@test.com",
		phone:     "000000003",
		contacted: true,
	},
}

func main() {
	// Serve index.html for requests made to root
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	fmt.Println("Starting server...")
	http.ListenAndServe(":3000", nil)
}
