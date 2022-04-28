package jwtauth

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"net/http"
	"sortarray/database"
	"sortarray/model"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var user model.User

	json.Unmarshal(requestBody, &user)

	if user.Username == "" || user.Password == "" {
		log.Println("Username or Password is empty, that cannot be tolerated")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	status := CreateUser(&user)
	w.WriteHeader(status)
}

func CreateUser(user *model.User) int {

	var userCheck model.User
	database.Connector.Where("username = ?", user.Username).First(&userCheck)
	if userCheck.Username != "" {
		log.Println(fmt.Sprintf("User: %s already exists, please chose another one", user.Username))
		return http.StatusForbidden
	}

	hashedBytePassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		log.Println(fmt.Sprintf("Error occured while hashing: %v", err))
		return http.StatusInternalServerError
	}

	newUser := model.User{
		Username: user.Username,
		Password: string(hashedBytePassword),
	}
	database.Connector.Save(newUser)
	log.Println(fmt.Sprintf("User : %s has been created", user.Username))
	return http.StatusCreated
}
