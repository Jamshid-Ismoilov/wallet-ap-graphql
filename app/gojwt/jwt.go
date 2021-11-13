package gojwt

import (
	"app/graph/model"
	"app/myerrors"
	"log"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret_key")

type Payload struct {
	ID    int
	Email string
	jwt.StandardClaims
}

func GenerateJWT(input model.JWTUser) string {
	payload := Payload{
		ID:             input.ID,
		Email:          input.Email,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		log.Fatalf("failed to take string token: %v", err)
	}
	return tokenString
}

func ParseJWT(token string) (model.JWTUser, error) {

	payload := &Payload{}

	tkn, err := jwt.ParseWithClaims(token, payload, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return model.JWTUser{ID: 0, Email: ""}, &myerrors.ErrSignatureInvalid{}
		}
		return model.JWTUser{ID: 0, Email: ""}, &myerrors.Unauthorized{}
	}
	if !tkn.Valid {
		return model.JWTUser{ID: 0, Email: ""}, &myerrors.Notvalid{}
	}
	return model.JWTUser{ID: payload.ID, Email: payload.Email}, nil

}
