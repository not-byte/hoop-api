package config

import (
	"os"
	"strconv"
	"tournament_api/server/types"

	"github.com/joho/godotenv"
)

func LoadConfig() (*types.AppConfig, error) {
	godotenv.Load()

	config := &types.AppConfig{
		PRODUCTION:                        getEnvAsBool("PRODUCTION", false),
		PUBLIC_HOST:                       getEnv("PUBLIC_HOST", "http://localhost"),
		PORT:                              getEnv("PORT", ":8080"),
		DB_USER:                           getEnv("DB_USER", "root"),
		DB_PASSWORD:                       getEnv("DB_PASSWORD", ""),
		DB_HOST:                           getEnv("DB_HOST", "eu2.notbyte.com"),
		DB_PORT:                           getEnv("DB_PORT", "60009"),
		DB_NAME:                           getEnv("DB_NAME", "tournament_dev"),
		JWT_ACCESS_SECRET:                 getEnv("JWT_ACCESS_SECRET", "not-so-secret-now-is-it?"),
		JWT_REFRESH_SECRET:                getEnv("JWT_REFRESH_SECRET", "not-so-secret-now-is-it?"),
		JWT_ACCESS_EXPIRATION_IN_SECONDS:  getEnvAsInt("JWT_ACCESS_EXPIRATION_IN_SECONDS", 600),         //12 minutes
		JWT_REFRESH_EXPIRATION_IN_SECONDS: getEnvAsInt("JWT_REFRESH_EXPIRATION_IN_SECONDS", 3600*24*14), //14 days
	}

	return config, nil
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}

func getEnvAsBool(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		b, err := strconv.ParseBool(value)
		if err != nil {
			return fallback
		}
		return b
	}
	return fallback
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
