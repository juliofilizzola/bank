package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	UrlDatabase = ""
	PORT        = ""
	SecretKey   = ""
	DbDriver    = ""
)

func Env() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	DbDriver = fmt.Sprint(os.Getenv("DBDRIVER"))
	PORT = fmt.Sprint(os.Getenv("API_PORT"))

	UrlDatabase = fmt.Sprint(os.Getenv("URL_DATABASE_ENV"))

	SecretKey = fmt.Sprint(os.Getenv("SECRET_KEY"))
}
