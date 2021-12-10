package utils

import (
	"database/sql"
	"fmt"
	"log"
)

type ClientDatabase struct {
	Id        string
	Host      string
	Username  string
	Passoword string
	Port      int
}

func DatabaseConnection(clientId string) *sql.DB {
	fmt.Println("Criando a conexão com o banco do cliente", clientId)
	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		log.Fatal(err)
	}

	return db

}
