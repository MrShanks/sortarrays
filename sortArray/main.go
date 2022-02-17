package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"math"
	"net/http"
	"os"
	"sortarray/database"
	"sortarray/service"
	"time"
)

func restController() {

	histogram := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "create_array",
		Help:    "Time taken to sort and store an array",
		Buckets: []float64{0.5, 1, 2, 4, 8, 16, 32},
	}, []string{"code"})

	prometheus.Register(histogram)

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", service.HomePage)
	myRouter.HandleFunc("/api/v1/all", service.GetAllArrays).Methods("GET")
	myRouter.HandleFunc("/api/v1/array/default", service.CreateNewArray(histogram)).Methods("POST")
	myRouter.HandleFunc("/api/v1/array/{id}", service.GetArrayByID).Methods("GET")
	myRouter.HandleFunc("/health", Health())
	myRouter.Handle("/metrics", promhttp.HandlerFor(
		prometheus.DefaultGatherer,
		promhttp.HandlerOpts{
			EnableOpenMetrics: true,
		},
	))

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

// Health Endpoint
func Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("status: ok\n"))
	}
}

func retryBackOff(attempt float64) {
	wait := math.Pow(2, attempt)
	next := math.Pow(2, attempt+1)
	time.Sleep(time.Duration(wait) * time.Millisecond)
	log.Println(fmt.Sprintf("Next connection attempt will be performed in %v seconds", next/1000))
}

func main() {

	counter := 0.0
	config :=
		database.Config{
			Hostname: os.Getenv("DB_HOSTNAME"),
			Port:     "3306",
			User:     "root",
			Password: os.Getenv("DB_PASSWORD"),
			Database: "array",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)

	for err != nil {
		retryBackOff(counter)
		err = database.Connect(connectionString)
		counter++
	}

	restController()
}
