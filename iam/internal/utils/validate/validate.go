package validate

import (
	"net/mail"
	"regexp"

	"github.com/HeyReyHR/twitch-clone/iam/internal/model"
)

func LoginInput(input, password string) error {
	if !isStrongPassword(password) {
		return model.ErrInvalidCredentials
	}
	if !isValidEmail(input) {
		return model.ErrInvalidEmail
	}
	return nil
}

func RegistrationInput(email, username, password string) error {
	if !isValidEmail(email) {
		return model.ErrInvalidEmail
	}

	if !isValidUsername(username) {
		return model.ErrInvalidUsername
	}

	if !isStrongPassword(password) {
		return model.ErrWeakPassword
	}

	return nil
}

func isValidEmail(email string) bool {
	if len(email) > 255 {
		return false
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return false
	}

	return true
}

func isValidUsername(username string) bool {
	if len(username) < 3 || len(username) > 30 {
		return false
	}
	matched, err := regexp.MatchString(`^[a-zA-Z0-9_]+$`, username)
	if err != nil {
		return false
	}

	return matched
}

func isStrongPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)

	return hasUpper && hasLower && hasNumber
}
