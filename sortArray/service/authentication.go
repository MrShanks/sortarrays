package jwtauth

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"net/http"
)

func AuthenticateWithJWT(w http.ResponseWriter, r *http.Request) error {
	//Get the session token from the request cookie
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return err
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	// Get the JWT string from the cookie
	tknStr := cookie.Value

	// Initialize a new instance of `Claims`
	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	// This method will return an error
	// if the token is invalid (if it has expired according to the expiry time set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return err
		}
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return errors.New("unauthorized user")
	}
	return nil
}
