package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     int    `json:"phone"`
	Contacted bool   `json:"contacted"`
}

var customersList = map[int]Customer{
	1: Customer{
		Name:      "Andrew",
		Role:      "User",
		Email:     "andrew@test.com",
		Phone:     0000001,
		Contacted: false,
	},
	2: Customer{
		Name:      "Brian",
		Role:      "User",
		Email:     "brian@test.com",
		Phone:     0000002,
		Contacted: false,
	},
	3: Customer{
		Name:      "Carla",
		Role:      "User",
		Email:     "carla@test.com",
		Phone:     0000003,
		Contacted: true,
	},
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customersList)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	w.Header().Set("Content-Type", "application/json")
	if _, ok := customersList[id]; ok {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customersList[id])
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newCustomer Customer
	reqBody, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(reqBody, &newCustomer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	customersList[len(customersList)+1] = newCustomer
	w.WriteHeader(http.StatusCreated)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, ok := customersList[id]; ok {
		var updatedCustomer Customer
		reqBody, _ := ioutil.ReadAll(r.Body)

		err := json.Unmarshal(reqBody, &updatedCustomer)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		customersList[id] = updatedCustomer
		w.WriteHeader(http.StatusAccepted)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if _, ok := customersList[id]; ok {
		delete(customersList, id)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func main() {
	// Initialise Router
	router := mux.NewRouter()

	// Serve index.html for requests made to root
	fileServer := http.FileServer(http.Dir("./static"))
	router.Handle("/", fileServer).Methods("GET")

	// Get All Customers
	router.HandleFunc("/customers", getCustomers).Methods("GET")

	// Get Customer by ID
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")

	// Create new Customer
	router.HandleFunc("/customers", addCustomer).Methods("POST")

	// Update existing Customer
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")

	// Delete existing Customer
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")

	fmt.Println("Starting server...")
	http.ListenAndServe(":3000", router)
}
