package main

import (
	chatdb "chatapp/db"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var jwtKey []byte
var db chatdb.DBLayer

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	gorm, err := chatdb.InitDB()
	if err != nil {
		log.Fatal("DB init failed:", err)
	}

	db = gorm
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
