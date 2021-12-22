package utils

import (
	"database/sql"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"path/filepath"
)

func LogDatabase(tabela string, tipo string, pk string, errCase bool, message string) {
	absPath, _ := filepath.Abs("./") //Root do projeto
	filePath := absPath + "/data/crm/relationDatabases.json"

	file, err := ioutil.ReadFile(filePath)

	if err != nil {
		fmt.Printf("File error: %v\n", err)
	}

	myJson := string(file)

	client := gjson.Get(myJson, "logs")

	dbLog, dbConnError := sql.Open(client.Get("database.Dialect").String(), client.Get("database.dataSourceName").String())
	if dbConnError != nil {
		panic("Could not connnect to log database")
	}
	defer dbLog.Close()

	queryString := fmt.Sprintf("INSERT INTO logs (tabela,tipo,pk,error,message) VALUES ('%s','%s','%s','%v','%s')", tabela, tipo, pk, errCase, message)

	r, queryError := dbLog.Exec(queryString)
	if queryError != nil {
		log.Println(queryError)
		panic("Could not execute log query")
	}
	fmt.Println(r)

}
