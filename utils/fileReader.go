package utils

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"path/filepath"
)

func loadSQLFile(sqlFile string) {
	//file, err := ioutil.ReadFile(sqlFile)
}

func readJson() {

}

type ClientData struct {
}

func loadClientConnection(id string) (gjson.Result, error) {

	absPath, _ := filepath.Abs("./") //Root do projeto
	filePath := absPath + "/data/databases.json"

	file, err := ioutil.ReadFile(filePath)

	if err != nil {
		fmt.Printf("File error: %v\n", err)
		return gjson.Result{}, err
	}

	myJson := string(file)

	client := gjson.Get(myJson, id)

	return client, nil

}
