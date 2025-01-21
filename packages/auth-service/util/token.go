package util

import (
	config "auth-service/pkg"
	"auth-service/src/dto"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

func CreateAccessToken(UserId int, Role string) (token *string, err error) {
	loginExpDurStr := config.AppLoginExpirationDuration()
	loginExpDur, err := strconv.Atoi(loginExpDurStr)
	loginExpirationDuration := time.Now().Add(time.Duration(loginExpDur) * time.Minute)

	claims := dto.ClaimAuthData{
		UserId,
		Role,
		jwt.RegisteredClaims{
			Issuer:    config.AppName(),
			ExpiresAt: jwt.NewNumericDate(loginExpirationDuration),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
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

func ParseAccessToken(token string) (res *dto.ClaimAuthData, err error) {
	var jwtKey = []byte(config.AppJwtSignatureKey())
	tokenParsed, err := jwt.ParseWithClaims(token, &dto.ClaimAuthData{}, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, TOKENINVALID
		} else if method != jwt.SigningMethodHS256 {
			return nil, TOKENINVALID
		}

		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := tokenParsed.Claims.(*dto.ClaimAuthData)
	if !ok {
		return nil, TOKENINVALID
	}
	res = claims
	return res, nil
}
