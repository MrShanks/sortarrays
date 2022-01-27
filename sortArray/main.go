package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sortarray/database"
	"sortarray/service"
)

func restController() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", service.HomePage)
	myRouter.HandleFunc("/api/v1/all", service.GetAllArrays).Methods("GET")
	myRouter.HandleFunc("/api/v1/array/default", service.CreateNewArray).Methods("POST")
	myRouter.HandleFunc("/api/v1/array/{id}", service.GetArrayByID).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {

	config :=
		database.Config{
			Hostname: "database",
			Port:     "3306",
			User:     "root",
			Password: "password",
			Database: "array",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)

	if err != nil {
		panic(err.Error())
	}

	restController()
}
