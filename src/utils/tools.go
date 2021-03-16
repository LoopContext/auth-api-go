package utils

import (
	"fmt"
	"os"
	"strconv"

	"github.com/loopcontext/checkmail"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

// ProjectContextKeys the project's context keys
var ProjectContextKeys = ContextKeys{
	GothicProviderCtxKey: "provider",
	ProviderCtxKey:       "gg-provider",
	UserCtxKey:           "gg-auth-user",
}

// ContextKey defines a type for context keys shared in the app
type ContextKey string

// ContextKeys holds the context keys throughout the project
type ContextKeys struct {
	GothicProviderCtxKey ContextKey // Provider for Gothic library
	ProviderCtxKey       ContextKey // Provider in Auth
	UserCtxKey           ContextKey // User db object in Auth
}

// GetEnv will return the env or empty string
func GetEnv(k string) string {
	return os.Getenv(k)
}

// MustGet will return the env or panic if it is not present
func MustGet(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Panic().Msgf("ENV missing, key: %s", k)
	}
	return v
}

// MustGetBool will return the env as boolean or panic if it is not present
func MustGetBool(k string) bool {
	v := os.Getenv(k)
	if v == "" {
		log.Panic().Msgf("ENV missing, key: %s", k)
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		log.Panic().Msgf("ENV key: %s - error: %q", k, err)
	}
	return b
}

// MustGetInt32 will return the env as int32 or panic if it is not present
func MustGetInt32(k string) int {
	v := os.Getenv(k)
	if v == "" {
		log.Panic().Msgf("ENV missing, key: %s", k)
	}
	i, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		log.Panic().Msgf("ENV key: %s - error: %q", k, err)
	}
	return int(i)
}

// MustGetInt64 will return the env as int64 or panic if it is not present
func MustGetInt64(k string) int64 {
	v := os.Getenv(k)
	if v == "" {
		log.Panic().Msgf("ENV missing, key: %s", k)
	}
	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		log.Panic().Msgf("ENV key: %s - error: %q", k, err)
	}
	return i
}

// EmailCheck checks the correct formatting of an email
func EmailCheck(email string) (err error) {
	if email == "" {
		return fmt.Errorf("email cannot be empty")
	}
	err = checkmail.ValidateFormat(email)
	if err != nil {
		return err
	}

	return nil
}

// HashPassword hash the password with bcrypt
func HashPassword(passw string) (string, error) {
	if passw != "" {
		if pw, err := bcrypt.GenerateFromPassword([]byte(passw), 11); err == nil {
			return string(pw), nil
		}
	}
	return "", fmt.Errorf("The password cannot be empty")
}

// PasswordCheck compares a password hash with what we have stored
func PasswordCheck(hashedPasswd string, passw string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(hashedPasswd), []byte(passw))
}

// StrToPtrStr return a pointer to the input string
func StrToPtrStr(str string) *string {
	return &str
}
