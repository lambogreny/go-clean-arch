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

//func DatabaseConnection(clientId string) *sql.DB {
//	fmt.Println("Criando a conexão com o banco do cliente", clientId)
//
//	switch clientId {
//	case "baroneza":
//		fmt.Println("Conectando ao postgresql da baroneza!")
//		conexao := "user=postgres dbname=baroneza password=dtcinfpostgresqlqta1 host=177.126.104.48 port=5434 sslmode=disable"
//		db, err := sql.Open("postgres", conexao)
//
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		err = db.Ping()
//
//		if err != nil {
//			panic(err)
//		}
//
//		return db
//
//	default:
//		db, err := sql.Open("sqlite3", "test.db")
//		if err != nil {
//			log.Fatal(err)
//		}
//		return db
//	}
//
//}

func DatabaseConnection(clientId string) (*sql.DB, error) {
	fmt.Println("Criando a conexão com o banco do cliente", clientId)

	connDict, err := loadClientConnection(clientId)

	if err != nil {
		LogFile("ERROR", "loadDatabaseConn", "CRITICAL", err.Error(), "Erro ao carregar as informações do cliente!")
		return nil, err
	}

	if connDict.Value() == nil {
		return nil, fmt.Errorf("Erro de cliente não encontrado!")
	}

	dev := connDict.Get("database.DEV")

	//Montando a string de conexão
	var connString string = fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", dev.Get("Username"), dev.Get("Dbname"), dev.Get("Passoword"), dev.Get("Host"), dev.Get("Port"))

	switch dev.Get("Dialect").Value() {
	case "postgresql":

		db, err := sql.Open("postgres", connString)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		err = db.Ping()

		if err != nil {
			panic(err)
		}

		return db, nil

	case "sqlite3":
		db, err := sql.Open("sqlite3", "test.db")
		if err != nil {
			return nil, err
		}
		return db, nil

	default:
		db, err := sql.Open("sqlite3", "test.db")
		if err != nil {
			return nil, err
		}
		return db, nil
	}

}
