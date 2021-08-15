package token

import (
	"restful-template/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	ISSUER  = "aimerneige.com"
	SUBJECT = "user token"
)

var jwtKey = []byte("restful-template-secret")

// ReleaseToken generate jwt token
func ReleaseToken(user models.User, tokenExpireDuration time.Duration) (string, error) {
	currentTime := time.Now()
	expirationTime := currentTime.Add(tokenExpireDuration)

	claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  currentTime.Unix(),
			Issuer:    ISSUER,
			Subject:   SUBJECT,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}

// ParseToken parse jwt token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	// token, err = jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
	// 	return jwtKey, nil
	// })
	// return

	claims := &Claims{}
	var token *jwt.Token
	var err error
	// 解码
	token, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})

	return token, claims, err
}
