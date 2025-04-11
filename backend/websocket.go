package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var (
	clients   = make(map[*websocket.Conn]string)
	broadcast = make(chan Message)
	lock      = sync.Mutex{}
)

func handleConnect(w http.ResponseWriter, r *http.Request) {
	tokenString := r.URL.Query().Get("token")

	if tokenString == "" {
		http.Error(w, "Missing Token", http.StatusUnauthorized)
		return
	}

	claims := &Claims{}

	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || claims.Username == "" {
		http.Error(w, "Invalid Token", http.StatusUnauthorized)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println("websocket upgrade error:", err)
	}

	defer conn.Close()

	lock.Lock()
	clients[conn] = claims.Username
	lock.Unlock()

	fmt.Println(claims.Username, "connect via websocket")

	for {
		var msg Message
		err := conn.ReadJSON(&msg)

		if err != nil {
			break
		}

		msg.Username = claims.Username
		broadcast <- msg
	}

}

func handleMessage() {
	for {
		msg := <-broadcast
		lock.Lock()
		for conn := range clients {
			err := conn.WriteJSON(msg)
			if err != nil {
				conn.Close()
				delete(clients, conn)
			}
		}

		lock.Unlock()
	}
}
