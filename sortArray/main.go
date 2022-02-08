package main

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log" "net/http"
	"sortarray/database"
	"sortarray/service"
)

func restController() {

	histogram := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "create_array",
		Help:    "Time taken to sort and store an array",
		Buckets: []float64{1, 2, 3, 4, 5},
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
