package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

func LogDatabase(clientId string, tabela string, tipo string, pk string, errCase bool, message string) {
	// absPath, _ := filepath.Abs("./") //Root do projeto
	// filePath := absPath + "/data/crm/relationDatabases.json"

	// file, err := ioutil.ReadFile(filePath)

	// if err != nil {
	// 	fmt.Printf("File error: %v\n", err)
	// }

	// myJson := string(file)

	// client := gjson.Get(myJson, "pgLogs")
	// fmt.Println(client)

	// var connString string = fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", client.Get("database.Username"), client.Get("database.Dbname"), client.Get("database.Passoword"), client.Get("database.Host"), client.Get("database.Port"))
	var connString string = fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", os.Getenv("LOG_DATABASE_USERNAME"), os.Getenv("LOG_DATABASE_DBNAME"), os.Getenv("LOG_DATABASE_PASSWORD"), os.Getenv("LOG_DATABASE_HOST"), os.Getenv("LOG_DATABASE_PORT"))
	dbLog, dbConnError := sql.Open("postgres", connString)

	if dbConnError != nil {
		panic("Could not connect to log database")
	}
	defer dbLog.Close()

	//Retirando todas as aspas da mensagem de erro, para evitar error no PG
	message = strings.ReplaceAll(message, "'", "")

	// queryString := fmt.Sprintf("INSERT INTO tb_logs (cliente,tabela,tipo,pk,error,message) VALUES ('%s','%s','%s','%s','%v','%s')", clientId, tabela, tipo, pk, errCase, message)
	queryString := fmt.Sprintf("INSERT INTO %s (cliente,tabela,tipo,pk,error,message) VALUES ('%s','%s','%s','%s','%v','%s')", os.Getenv("LOG_DATABASE_TABLE"), clientId, tabela, tipo, pk, errCase, message)
	// fmt.Println(connString)
	// fmt.Println(queryString)

	_, queryError := dbLog.Exec(queryString)
	if queryError != nil {
		log.Println(queryError)
		panic("Could not execute log query")
	}

}

func LogDatabaseDetails(tabela string, pk string, queryString string, dbResponse string, responseType string) {

	// absPath, _ := filepath.Abs("./") //Root do projeto
	// filePath := absPath + "/data/crm/relationDatabases.json"

	// file, err := ioutil.ReadFile(filePath)

	// if err != nil {
	// 	fmt.Printf("File error: %v\n", err)
	// }

	// myJson := string(file)

	// client := gjson.Get(myJson, "pgLogs")

	// var connString string = fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", client.Get("database.Username"), client.Get("database.Dbname"), client.Get("database.Passoword"), client.Get("database.Host"), client.Get("database.Port"))
	var connString string = fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", os.Getenv("LOG_DATABASE_USERNAME"), os.Getenv("LOG_DATABASE_DBNAME"), os.Getenv("LOG_DATABASE_PASSWORD"), os.Getenv("LOG_DATABASE_HOST"), os.Getenv("LOG_DATABASE_PORT"))

	dbLog, dbConnError := sql.Open("postgres", connString)

	if dbConnError != nil {
		panic("Could not connect to log database")
	}
	defer dbLog.Close()

	//Retirando todas as aspas da mensagem de erro, para evitar error no PG
	// queryString = strings.ReplaceAll(queryString, "'", "''")
	queryString = strings.ReplaceAll(queryString, "'", "")
	dbResponse = strings.ReplaceAll(dbResponse, "'", "")

	insertString := fmt.Sprintf("INSERT INTO %s (tabela,pk,queryString,dbResponse,responseType) VALUES ('%s','%s','%s','%s','%s')", os.Getenv("LOG_DETAILS_TABLE"), tabela, pk, queryString, dbResponse, responseType)

	// fmt.Println(insertString)

	_, queryError := dbLog.Exec(insertString)
	if queryError != nil {
		log.Println(insertString)
		log.Println(queryError)
		panic("Could not execute log_details query")
	}

}
