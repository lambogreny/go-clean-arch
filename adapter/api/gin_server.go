package api

import (
	"github.com/joho/godotenv"
)

func Init() {

	godotenv.Load(".env")
	//serverPort := os.Getenv("SERVER_PORT")
	serverPort := ":8080"
	r := NewRouter()
	r.Run(serverPort)
}
