package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	UserID uint
	jwt.StandardClaims
}