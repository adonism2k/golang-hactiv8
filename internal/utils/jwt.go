package utils

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/adonism2k/golang-hactiv8/internal/model"
	"github.com/golang-jwt/jwt/v4"
)

// AuthTokenClaim struct
type AuthTokenClaim struct {
	jwt.RegisteredClaims
	User model.User
}

var JWT_SIGNING_METHOD = jwt.SigningMethodHS256

func GenerateToken(user model.User, ttl time.Duration, secretKey string) (string, error) {
	newJWT := jwt.NewWithClaims(JWT_SIGNING_METHOD, &AuthTokenClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl * time.Hour)),
		},
		User: user,
	})

	token, err := newJWT.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(token), nil
}

func ValidateToken(token string, secretKey string) (model.User, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &AuthTokenClaim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			log.Println(err)
			return model.User{}, errors.New("signature invalid")
		}
		log.Println(err)
		return model.User{}, errors.New("could not parse the auth token")
	}

	if !parsedToken.Valid {
		log.Println(err)
		return model.User{}, errors.New("invalid token")
	}

	log.Println("TOKEN is valid:", parsedToken.Valid)

	if claims, ok := parsedToken.Claims.(*AuthTokenClaim); ok && parsedToken.Valid {
		return claims.User, nil
	}

	return model.User{}, err
}
