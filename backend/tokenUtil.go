package main

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


func GenerateToken(id int) ( string,  error) {
	
	

	claims := MyCustomClaim{
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	mySignedKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	ss, err := token.SignedString(mySignedKey)

	if err != nil {
		return "", err
	}

	return ss, nil
}

func ParseToken(tokenStr string) (*int, error) {
	
	token, err := jwt.ParseWithClaims(tokenStr, &MyCustomClaim{}, func(t *jwt.Token) (interface{}, error){
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}
	
	claim, ok := token.Claims.(*MyCustomClaim)
	if !ok {
		return nil, err
	}


	return &claim.UserId, err
}