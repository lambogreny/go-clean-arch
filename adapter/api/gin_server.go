package api

import (
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	godotenv.Load(".env")
	serverPort := os.Getenv("SERVER_PORT")
	r := NewRouter()
	r.Run(serverPort)
}
