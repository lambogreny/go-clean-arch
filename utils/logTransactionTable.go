package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func LogDatabase(clientId string, tabela string, tipo string, pk string, errCase bool, message string) {
	absPath, _ := filepath.Abs("./") //Root do projeto
	filePath := absPath + "/data/crm/relationDatabases.json"

	file, err := ioutil.ReadFile(filePath)

	if err != nil {
		fmt.Printf("File error: %v\n", err)
	}

	myJson := string(file)

	//client := gjson.Get(myJson, "logs") //sqlite
	client := gjson.Get(myJson, "pgLogs")

	//Banco sqlite
	//dbLog, dbConnError := sql.Open(client.Get("database.Dialect").String(), client.Get("database.dataSourceName").String())

	var connString string = fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", client.Get("database.Username"), client.Get("database.Dbname"), client.Get("database.Passoword"), client.Get("database.Host"), client.Get("database.Port"))
	//dbLog, dbConnError := sql.Open(client.Get("database.Dialect").String(), connString)
	dbLog, dbConnError := sql.Open("postgres", connString)

	if dbConnError != nil {
		panic("Could not connect to log database")
	}
	defer dbLog.Close()

	//Retirando todas as aspas da mensagem de erro, para evitar error no PG
	message = strings.ReplaceAll(message, "'", "")

	queryString := fmt.Sprintf("INSERT INTO tb_logs (cliente,tabela,tipo,pk,error,message) VALUES ('%s','%s','%s','%s','%v','%s')", clientId, tabela, tipo, pk, errCase, message)

	//fmt.Println(queryString)

	r, queryError := dbLog.Exec(queryString)
	if queryError != nil {
		log.Println(queryError)
		panic("Could not execute log query")
	}
	fmt.Println(r.RowsAffected())

}
