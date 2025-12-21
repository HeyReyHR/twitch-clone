package model

import "time"

type User struct {
	UserId       string    `json:"user_id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Role         Role      `json:"role"`
	PasswordHash string    `json:"password_hash"`
	AvatarUrl    string    `json:"avatar_url"`
	IsStreaming  bool      `json:"is_streaming"`
	StreamKey    string    `json:"stream_key"`
	UpdatedAt    time.Time `json:"updated_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type RefreshToken struct {
	Id           string    `json:"id"`
	UserId       string    `json:"user_id"`
	RefreshToken string    `json:"refresh_token"`
	CreatedAt    time.Time `json:"updated_at"`
	ExpiresAt    time.Time `json:"expires_at"`
}

type Role string

const (
	UNKNOWN Role = "UNKNOWN"
	USER    Role = "USER"
	ADMIN   Role = "ADMIN"
)
