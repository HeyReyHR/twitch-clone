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

type TokenPair struct {
	AccessToken           string
	RefreshToken          string
	AccessTokenExpiresAt  time.Time
	RefreshTokenExpiresAt time.Time
}

type Claims struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	Role     Role   `json:"role"`
}
