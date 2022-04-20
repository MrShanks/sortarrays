package handler

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
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

<<<<<<< HEAD:sortArray/service/handler.go
	myRouter.HandleFunc("/", HomePage)
	myRouter.HandleFunc("/signin", jwtauth.SignIn).Methods("POST")
	myRouter.HandleFunc("/refresh", jwtauth.Refresh).Methods("POST")
	myRouter.HandleFunc("/api/v1/all", GetAllArrays).Methods("GET")
	myRouter.HandleFunc("/api/v1/array/default", CreateNewArray(histogram)).Methods("POST")
	myRouter.HandleFunc("/api/v1/array/{id}", GetArrayByID).Methods("GET")
	myRouter.HandleFunc("/health", Health())
=======
	myRouter.HandleFunc("/", service.HomePage)
	myRouter.HandleFunc("/signup", service.SignUp)
	myRouter.HandleFunc("/signin", service.SignIn).Methods("POST")
	myRouter.HandleFunc("/api/v1/all", service.GetAllArrays).Methods("GET")
	myRouter.HandleFunc("/api/v1/array/default", service.CreateNewArray(histogram)).Methods("POST")
	myRouter.HandleFunc("/api/v1/array/{id}", service.GetArrayByID).Methods("GET")
	myRouter.HandleFunc("/health", service.Health())
>>>>>>> origin/DEV-28-implement-user-endpoint-to-create-a-new-user:sortArray/handler/handler.go
	myRouter.Handle("/metrics", promhttp.HandlerFor(
		prometheus.DefaultGatherer,
		promhttp.HandlerOpts{
			EnableOpenMetrics: true,
		},
	))

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
