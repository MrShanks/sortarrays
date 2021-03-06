package jwtauth

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"sortarray/database"
	"sortarray/model"
	"time"
)

// Create the JWT key used to create the signature
var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var creds model.User
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = CheckUserPassword(&creds)
	if err != nil {
		log.Println(fmt.Sprintf("User: %s authentication failed:%v", creds.Username, err))
		return
	}

	// Declare the expiration time of the token
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time must be expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Create the token with the algorithm, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Finally, set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func CheckUserPassword(creds *model.User) error {
	var dbUser model.User
	err := database.Connector.Where("username = ?", creds.Username).First(&dbUser).Error
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(creds.Password))
	if err != nil {
		return err
	}
	log.Println(fmt.Sprintf("User: %s successfully authenticated", creds.Username))
	return nil
}
