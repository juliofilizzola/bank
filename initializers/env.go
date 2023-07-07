package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	UrlDatabase     = ""
	PORT            = ""
	SecretKey       = ""
	DbDriver        = ""
	UrlDatabaseTest = ""
	DNS             = ""
)

func Env() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	DbDriver = fmt.Sprint(os.Getenv("DBDRIVER"))
	PORT = fmt.Sprint(os.Getenv("API_PORT"))
	DNS = fmt.Sprint(os.Getenv("DNS"))

	UrlDatabase = fmt.Sprint(os.Getenv("URL_DATABASE_ENV"))
	UrlDatabaseTest = fmt.Sprint(os.Getenv("URL_DATABASE_ENV_TEST"))

	SecretKey = fmt.Sprint(os.Getenv("SECRET_KEY"))
}
