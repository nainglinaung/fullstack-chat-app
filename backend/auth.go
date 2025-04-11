package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// var user = map[string]string{}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := &User{UserName: creds.Username, Password: creds.Password}

	if err := DB.Create(&user).Error; err != nil {
		http.Error(w, "User already exists or erro occur", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user User

	result := DB.Where("username = ?", creds.Username).First(&user)

	if result.Error != nil || user.Password != creds.Password {
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: creds.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
