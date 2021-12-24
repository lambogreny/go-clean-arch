package crm

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"path/filepath"
)

/*
	Função que carrega as strings de conexão
	Devolve os objetos do De-Para
*/
func loadStringConnections(id string) (gjson.Result, error) {

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

func GetPedidoUrl(id string) (string, error) {

	absPath, _ := filepath.Abs("./") //Root do projeto
	filePath := absPath + "/data/crm/relationDatabases.json"

	file, err := ioutil.ReadFile(filePath)

	if err != nil {
		return "", err
	}

	myJson := string(file)

	client := gjson.Get(myJson, id)

	return client.Get("pedidosInfo.base_url").String(), nil

}

/*
	Função que devolve a conexão dos bancos: CRM , ti9, Owner do CRM
	Além do ownner da tabela do CRM
*/
func ServicesDatabases(clientId string) (*sql.DB, *sql.DB, string, error) {
	connDict, err := loadStringConnections(clientId)

	if err != nil {
		return nil, nil, "", err
	}

	if connDict.Value() == nil {
		return nil, nil, "", fmt.Errorf("Erro de cliente não encontrado!")
	}

	//Checando se cliente está ativo
	if connDict.Get("Active").Value() != true {
		return nil, nil, "", fmt.Errorf("Client is not active!")
	}

	dbCrm := connDict.Get("db_CRM")
	dbErp := connDict.Get("db_ERP")

	crmOwner := dbCrm.Get("Owner").String()

	dbCrmString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbCrm.Get("Username"), dbCrm.Get("Passoword"), dbCrm.Get("Host"), dbCrm.Get("Port"), dbCrm.Get("Dbname"))
	dbErpString := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", dbErp.Get("Username"), dbErp.Get("Dbname"), dbErp.Get("Passoword"), dbErp.Get("Host"), dbErp.Get("Port"))

	dbCrmConn, errCrm := sql.Open("mysql", dbCrmString)

	dbErpConn, errErp := sql.Open("postgres", dbErpString)

	if errCrm != nil || errErp != nil {
		log.Println(errErp, errCrm)
		return nil, nil, "", fmt.Errorf("Could not connect to integrations databases")
	}

	//Realizando o ping na conexão, para ver se os bancos estão acessíveis
	errPingCrm := dbCrmConn.Ping()
	errPingErp := dbErpConn.Ping()

	if errPingCrm != nil || errPingErp != nil {
		log.Println("erro ao realizar ping nos databases")
		return nil, nil, "", fmt.Errorf("Could not access integrations databases")
	}

	return dbCrmConn, dbErpConn, crmOwner, nil
}
