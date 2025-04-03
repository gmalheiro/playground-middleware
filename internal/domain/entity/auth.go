package entity

import "github.com/golang-jwt/jwt/v5"

type AuthCredentials struct {
	Email    string
	Password string
}

type Claims struct {
	Email string
	jwt.RegisteredClaims
}
