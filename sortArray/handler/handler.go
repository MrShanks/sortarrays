package handler

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"sortarray/jwtauth"
	"sortarray/service"
)

func RestController() {

	histogram := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "create_array",
		Help:    "Time taken to sort and store an array",
		Buckets: []float64{0.5, 1, 2, 4, 8, 16, 32},
	}, []string{"code"})

	prometheus.Register(histogram)

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", service.HomePage)
	myRouter.HandleFunc("/signin", jwtauth.SignIn).Methods("POST")
	myRouter.HandleFunc("/signup", jwtauth.SignUp)
	myRouter.HandleFunc("/refresh", jwtauth.Refresh).Methods("POST")
	myRouter.HandleFunc("/api/v1/all", service.GetAllArrays).Methods("GET")
	myRouter.HandleFunc("/api/v1/array/default", service.CreateNewArray(histogram)).Methods("POST")
	myRouter.HandleFunc("/api/v1/array/{id}", service.GetArrayByID).Methods("GET")
	myRouter.HandleFunc("/health", service.Health())
	myRouter.Handle("/metrics", promhttp.HandlerFor(
		prometheus.DefaultGatherer,
		promhttp.HandlerOpts{
			EnableOpenMetrics: true,
		},
	))

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
