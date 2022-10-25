package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/adonism2k/golang-hactiv8/internal/model"
	"github.com/golang-jwt/jwt/v4"
)

// AuthToken struct
type AuthToken struct {
	Token string `json:"token"`
}

// AuthTokenClaim struct
type AuthTokenClaim struct {
	jwt.RegisteredClaims
	User model.User
}

var JWT_SECRET = []byte(os.Getenv("JWT_SECRET"))	

func GenerateToken(user model.User) (AuthToken, error) {
	var JWT_SIGNING_METHOD = jwt.SigningMethodHS256

	expired, err := strconv.Atoi(os.Getenv("JWT_EXPIRED"))
	if err != nil {
		fmt.Println(err)
		return AuthToken{}, err
	}

	newJWT := jwt.NewWithClaims(JWT_SIGNING_METHOD, AuthTokenClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expired))),
		},
		User: user,
	})

	token, err := newJWT.SignedString(JWT_SECRET)
	if err != nil {
		fmt.Println(err)
		return AuthToken{}, err
	}

	return AuthToken{token}, nil
}

func ValidateToken(tokenString string) (model.User, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AuthTokenClaim{}, func(token *jwt.Token) (interface{}, error) {
		return JWT_SECRET, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return model.User{}, errors.New("signature invalid")
		}
		return model.User{}, errors.New("could not parse the auth token")
	}

	if !token.Valid {
		return model.User{}, errors.New("invalid token")
	}

	fmt.Println("TOKEN is:", token.Valid)

	if claims, ok := token.Claims.(*AuthTokenClaim); ok && token.Valid {
		return claims.User, nil
	}

	return model.User{}, err
}
