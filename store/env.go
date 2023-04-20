package store

import "os"

var (
	EnvVars *Env
)

type Env struct {
	DbHost string
	DbPort string
	DbName string
	DbUser string
	DbPass string
}

func loadEnvVars() *Env {
	return &Env{
		DbHost: os.Getenv("DB_HOST"),
		DbPort: os.Getenv("DB_PORT"),
		DbName: os.Getenv("DB_NAME"),
		DbUser: os.Getenv("DB_USER"),
		DbPass: os.Getenv("DB_PASS"),
	}
}
