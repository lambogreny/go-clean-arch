package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
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
	//db, err := sql.Open("sqlite3", "test.db")

	switch clientId {
	case "baroneza":
		fmt.Println("Conectando ao postgresql da baroneza!")
		conexao := "user=postgres dbname=baroneza password=dtcinfpostgresqlqta1 host=177.126.104.48 port=5434 sslmode=disable"
		db, err := sql.Open("postgres", conexao)

		if err != nil {
			log.Fatal(err)
		}

		err = db.Ping()

		if err != nil {
			panic(err)
		}

		return db

	default:
		db, err := sql.Open("sqlite3", "test.db")
		if err != nil {
			log.Fatal(err)
		}
		return db
	}

}
