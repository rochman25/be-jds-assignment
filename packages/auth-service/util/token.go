package util

import (
	config "auth-service/pkg"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

func CreateAccessToken(UserId int) (token *string, err error) {
	loginExpDurStr := config.AppLoginExpirationDuration()
	loginExpDur, err := strconv.Atoi(loginExpDurStr)
	loginExpirationDuration := time.Now().Add(time.Duration(loginExpDur) * time.Minute)

	claims := jwt.RegisteredClaims{
		Issuer:    config.AppName(),
		ExpiresAt: jwt.NewNumericDate(loginExpirationDuration),
		ID:        strconv.Itoa(UserId),
	}

	tokenJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var jwtKey = []byte(config.AppJwtSignatureKey())
	tokenString, err := tokenJwt.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}
	token = &tokenString
	return token, nil
}
