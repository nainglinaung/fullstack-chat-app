package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var jwtKey []byte

func main() {
	err := godotenv.Load()

	initDB()

	if err != nil {
		log.Fatal("Error Loading .env file")
	}
	jwtKey = []byte(os.Getenv("JWT_SECRET"))

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}
	fmt.Println("server running at " + port)

	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/ws", handleConnect)

	go handleMessage()

	http.ListenAndServe(":"+port, nil)
}
