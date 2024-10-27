package vars

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	_ = godotenv.Load()

	PORT        = os.Getenv("PORT")
	APP_MODE    = os.Getenv("APP_MODE")
	SERVER_URL  = os.Getenv("SERVER_URL")
	APP_VERSION = os.Getenv("APP_VERSION")
	APP_NAME    = os.Getenv("APP_NAME")

	DATABASE_URL = os.Getenv("DATABASE_URL")

	JWT_SECRET = os.Getenv("JWT_SECRET")
)

func init() {
	if DATABASE_URL == "" {
		log.Fatal("DATABASE_URL не установлен в .env файле\nDATABASE_URL", DATABASE_URL)
	}
}
