package password

import (
	"github.com/HeyReyHR/twitch-clone/iam/internal/config"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	passBytes, err := bcrypt.GenerateFromPassword([]byte(passWithSalt(password)), bcrypt.MinCost)
	return string(passBytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(passWithSalt(password)))
	return err == nil
}

func passWithSalt(password string) string {
	return password + config.AppConfig().Password.PasswordSalt()
}
