package env

import (
	"os"

	"github.com/Pedro-Cecilio/projeto-livro-go/models"
	_ "github.com/joho/godotenv/autoload"
)

var Env models.Env = models.Env{
	KEY_JWT: os.Getenv("KEY_JWT"),
	DB_NAME: os.Getenv("DB_NAME"),
	DB_USER: os.Getenv("DB_USER"),
	DB_PASSWORD: os.Getenv("DB_PASSWORD"),
}

