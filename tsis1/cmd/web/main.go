package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	// "github.com/margulan-kalykul/Golang/pkg/types"
)

type Response struct{
	Persons []Person `json:"persons"`
}

type Person struct{
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main(){
	//create a new router
	router := mux.NewRouter()

	//specify endpoints, handler functions and HTTP method
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/persons", Persons).Methods("GET")
	http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":8080", router)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	//specify status code
	w.WriteHeader(http.StatusOK)
  
  //update response writer 
	fmt.Fprintf(w, "API is up and running")
}

func Persons(w http.ResponseWriter, r *http.Request) {
	//declare response variable
	var response Response
  
	//Retrieve person details
	persons := prepareResponse()
  
	//assign person details to response
	response.Persons = persons
  
	//update content type
	w.Header().Set("Content-Type", "application/json")
  
	//specify HTTP status code
	w.WriteHeader(http.StatusOK)
  
	//convert struct to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
	 return
	}
  
	//update response
	w.Write(jsonResponse)
  }

  func prepareResponse() []Person {
	var persons []Person
  
	var person Person
	person.Id = 1
	person.FirstName = "Issac"
	person.LastName = "Newton"
	persons = append(persons, person)
  
	person.Id = 2
	person.FirstName = "Albert"
	person.LastName = "Einstein"
	persons = append(persons, person)
  
	return persons
  }