package service

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/crypto/bcrypt"
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
		err := AuthenticateWithJWT(w, r)
		if err != nil {
			log.Fatal(err.Error())
		}
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

// Health Endpoint
func Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("status: ok\n"))
	}
}

func CreateUser(user *model.User) {
	hashedBytePassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		log.Println(fmt.Sprintf("Error occured while hashing: %v", err))
	}

	newUser := model.User{
		Username: user.Username,
		Password: string(hashedBytePassword),
	}
	database.Connector.Save(newUser)
	log.Println(fmt.Sprintf("User : %s has been created", user.Username))
}

func CheckUserPassword(creds *model.User) {
	var dbUser model.User
	database.Connector.Where("username = ?", creds.Username).First(&dbUser)
	err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(creds.Password))
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(fmt.Sprintf("User: %s successfully authenticated", creds.Username))
}
