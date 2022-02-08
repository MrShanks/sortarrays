package service

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"sort"
	"sortarray/database"
	"sortarray/model"
	"time"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func CreateNewArray(histogram *prometheus.HistogramVec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		code := 500

		requestBody, _ := ioutil.ReadAll(r.Body)

		var array model.Array
		var tmpArray model.TMPArray
		var unorderedArray string

		defer func() {
			httpDuration := time.Since(start)
			histogram.WithLabelValues(fmt.Sprintf("%d", code)).Observe(httpDuration.Seconds())
		}()

		json.Unmarshal(requestBody, &tmpArray)

		unorderedArray = fmt.Sprintf("%v", tmpArray.Elements)

		ShuffleSort(tmpArray.Elements)

		array = model.Array{
			Id:       tmpArray.Id,
			Elements: fmt.Sprintf("%v", tmpArray.Elements),
		}

		database.Connector.Create(&array)
		log.Println(fmt.Sprintf("new array has been created: %v and sorted: %s", unorderedArray, array.Elements))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(array)
	}
}

func ShuffleSort(array []int) {

	for !sort.SliceIsSorted(array, func(i, j int) bool {
		return array[i] < array[j]
	}) {
		rand.Seed(time.Now().Unix())

		rand.Shuffle(len(array), func(i, j int) {
			array[i], array[j] = array[j], array[i]
		})
	}
}

func GetArrayByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var array model.Array
	database.Connector.First(&array, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(array)
}

func GetAllArrays(w http.ResponseWriter, r *http.Request) {
	var arrays []model.Array
	database.Connector.Find(&arrays)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(arrays)
}
