package crm

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"path/filepath"
)

func loadStringConnections(id string) (gjson.Result, error) {
	fmt.Println("A função de connectionString recebeu o id", id)
	absPath, _ := filepath.Abs("./") //Root do projeto
	filePath := absPath + "/data/crm/relationDatabases.json"

	file, err := ioutil.ReadFile(filePath)

	if err != nil {
		fmt.Printf("File error: %v\n", err)
		return gjson.Result{}, err
	}

	myJson := string(file)

	client := gjson.Get(myJson, id)

	return client, nil
}

//Função que devolve as duas conexões
func ServicesDatabases(clientId string) (*sql.DB, *sql.DB, error) {
	connDict, err := loadStringConnections(clientId)

	if err != nil {
		return nil, nil, err
	}

	if connDict.Value() == nil {
		return nil, nil, fmt.Errorf("Erro de cliente não encontrado!")
	}

	dbCrm := connDict.Get("db_CRM")
	dbErp := connDict.Get("db_ERP")

	dbCrmString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbCrm.Get("Username"), dbCrm.Get("Passoword"), dbCrm.Get("Host"), dbCrm.Get("Port"), dbCrm.Get("Dbname"))
	dbErpString := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", dbErp.Get("Username"), dbErp.Get("Dbname"), dbErp.Get("Passoword"), dbErp.Get("Host"), dbErp.Get("Port"))

	dbCrmConn, errCrm := sql.Open("mysql", dbCrmString)

	dbErpConn, errErp := sql.Open("postgresql", dbErpString)

	if errCrm != nil || errErp != nil {
		return nil, nil, errCrm
	}

	errPingCrm := dbCrmConn.Ping()
	errPingErp := dbErpConn.Ping()

	if errPingCrm != nil || errPingErp != nil {
		return nil, nil, errPingCrm
	}

	return dbCrmConn, dbErpConn, nil
}
