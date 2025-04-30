package model

import "time"

type User struct {
	ID           int
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password""`
	CreateAt     time.Time
}
