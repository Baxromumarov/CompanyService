package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	Environment             string // develop, staging, production
	RPCPort                 string

	PostgresHost            string
	PostgresPort            int
	PostgresDatabase        string
	PostgresUser            string
	PostgresPassword        string

	AuthServiceHost         string
	AuthServicePort         int

}

func Load() *Config {
	c := &Config{}
	path, ok := os.LookupEnv("ENV_FILE_PATH")
	if ok && path != "" {
		if err := godotenv.Load(path); err != nil {
			log.Print("No .env file found")
		}
	}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	c.RPCPort = cast.ToString(getOrReturnDefault("RPC_PORT", "8002"))

	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "eld_go_company_service"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "postgres"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "password"))

	c.AuthServiceHost = cast.ToString(getOrReturnDefault("AUTH_SERVICE_HOST", "localhost"))
	c.AuthServicePort = cast.ToInt(getOrReturnDefault("AUTH_SERVICE_PORT", 8080))
	return c
}

func getOrReturnDefault(key string, defaultValue any) any {
	v, exists := os.LookupEnv(key)
	if exists {
		return v
	}

	return defaultValue
}
