package utils

import (
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
)

// ContextKey defines a type for context keys shared in the app
type ContextKey string

// ContextKeys holds the context keys throughout the project
type ContextKeys struct {
	GothicProviderCtxKey ContextKey // Provider for Gothic library
	ProviderCtxKey       ContextKey // Provider in Auth
	UserCtxKey           ContextKey // User db object in Auth
}

var (
	// ProjectContextKeys the project's context keys
	ProjectContextKeys = ContextKeys{
		GothicProviderCtxKey: "provider",
		ProviderCtxKey:       "gg-provider",
		UserCtxKey:           "gg-auth-user",
	}
)

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
