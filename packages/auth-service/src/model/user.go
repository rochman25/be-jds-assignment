package model

import "github.com/golang/protobuf/ptypes/timestamp"

type User struct {
	Id        int                 `json:"id"`
	Nik       string              `json:"nik"`
	Role      string              `json:"role"`
	Password  string              `json:"password"`
	CreatedAt timestamp.Timestamp `json:"created_at"`
	UpdatedAt timestamp.Timestamp `json:"updated_at"`
}
