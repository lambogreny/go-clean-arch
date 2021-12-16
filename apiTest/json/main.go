package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"os"
	"reflect"
)

func main() {
	fmt.Println("Testando a manipulaçãop de json")

	file, err := ioutil.ReadFile("./nestedJson.json")

	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	myJson := string(file)

	var client string
	client = "client2"

	teste := gjson.Get(myJson, client)
	fmt.Println(teste)
	fmt.Println(reflect.TypeOf(teste))
	if teste.Value() == nil {
		fmt.Println("Cliente nao encontrado!")
	}
	fmt.Println("O id é : ", teste.Get("id"))
	//fmt.Println(string(jsonBytes))

}
