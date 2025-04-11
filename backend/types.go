package main

import "github.com/golang-jwt/jwt/v5"

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// // Valid implements jwt.Claims.
// func (c *Claims) Valid() error {
// 	return c.RegisteredClaims.Validate()
// }

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}
