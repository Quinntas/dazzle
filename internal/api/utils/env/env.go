package env

import (
	"context"
	"os"

	"github.com/joho/godotenv"
)

func getEnv(key string, required bool) string {
	value := os.Getenv(key)
	if required && value == "" {
		panic("Missing required environment variable: " + key)
	}
	return value
}

type EnvVariables struct {
	Port        string
	DatabaseURL string
	RedisURL    string
	Pepper      string
	JwtSecret   string
}

func GetEnvVariablesFromContext(ctx context.Context) EnvVariables {
	return ctx.Value(ENV_CTX_KEY).(EnvVariables)
}

func NewEnvVariables() EnvVariables {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	return EnvVariables{
		Port:        getEnv("PORT", true),
		DatabaseURL: getEnv("DATABASE_URL", true),
		RedisURL:    getEnv("REDIS_URL", true),
		Pepper:      getEnv("PEPPER", true),
		JwtSecret:   getEnv("JWT_SECRET", true),
	}
}
