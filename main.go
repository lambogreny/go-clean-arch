package main

import (
	"github.com/augusto/imersao5-esquenta-go/adapter/api"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	//Iniciar a api Echo
	//webserver := api.NewWebServer()
	//webserver.Repository = repo
	//webserver.Serve()

	//Inicia a api Gin
	api.Init()
}
