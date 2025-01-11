package dto

import "github.com/golang-jwt/jwt/v5"

type (
	RegisterRequest struct {
		Nik  string `json:"nik" binding:"required,min=16,max=16"`
		Role string `json:"role" binding:"required"`
	}

	RegisterResponse struct {
		Nik      string `json:"nik"`
		Role     string `json:"role"`
		Password string `json:"password"`
	}

	LoginRequest struct {
		Nik      string `json:"nik" binding:"required,min=16,max=16"`
		Password string `json:"password" binding:"required,min=6,max=6"`
	}

	LoginResponse struct {
		ID          int    `json:"id"`
		Nik         string `json:"nik"`
		Role        string `json:"role"`
		AccessToken string `json:"access_token"`
	}

	ClaimAuthData struct {
		UserId int `json:"user_id"`
		jwt.RegisteredClaims
	}
)
