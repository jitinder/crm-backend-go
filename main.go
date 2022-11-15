package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Contacted bool   `json:"contacted"`
}

var customersList = map[int]Customer{
	1: Customer{
		Name:      "Andrew",
		Role:      "User",
		Email:     "andrew@test.com",
		Phone:     "000000001",
		Contacted: false,
	},
	2: Customer{
		Name:      "Brian",
		Role:      "User",
		Email:     "brian@test.com",
		Phone:     "000000002",
		Contacted: false,
	},
	3: Customer{
		Name:      "Carla",
		Role:      "User",
		Email:     "carla@test.com",
		Phone:     "000000003",
		Contacted: true,
	},
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customersList)
}

func main() {
	// Initialise Router
	router := mux.NewRouter()

	// Serve index.html for requests made to root
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	// Get All Users
	router.HandleFunc("/customers", getAllCustomers).Methods("GET")

	fmt.Println("Starting server...")
	http.ListenAndServe(":3000", router)
}
