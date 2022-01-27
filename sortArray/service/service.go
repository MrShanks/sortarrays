package service

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"sortarray/database"
	"sortarray/model"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func CreateNewArray(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var array model.Array
	var tmpArray model.TMPArray
	var unorderedArray string

	json.Unmarshal(requestBody, &tmpArray)

	unorderedArray = fmt.Sprintf("%v", tmpArray.Elements)

	sort.Ints(tmpArray.Elements)

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

func GetArrayByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var array model.Array
	database.Connector.First(&array, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(array)
}

func GetAllArrays(w http.ResponseWriter, r *http.Request) {
	var array []model.Array
	database.Connector.Find(&array)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(array)
}
