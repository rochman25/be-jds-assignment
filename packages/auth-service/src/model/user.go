package model

import (
	"time"
)

type User struct {
	Id        int       `json:"id"`
	Nik       string    `json:"nik"`
	Role      string    `json:"role"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
