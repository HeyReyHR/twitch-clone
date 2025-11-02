package model

import "time"

type User struct {
	UserId       string
	Username     string
	Email        string
	Role         Role
	PasswordHash string
	UpdatedAt    time.Time
	CreatedAt    time.Time
}

type Role string

const (
	UNKNOWN Role = "UNKNOWN"
	USER    Role = "USER"
	ADMIN   Role = "ADMIN"
)
