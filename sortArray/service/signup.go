package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sortarray/model"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var user model.User

	json.Unmarshal(requestBody, &user)

	if user.Username == "" || user.Password == "" {
		log.Println("Username or Password is empty, that cannot be tolerated")
	}

	CreateUser(&user)
}

