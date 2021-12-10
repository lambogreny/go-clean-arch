package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Book struct {
	Name      string `json:"name"`
	Publisher string `json:"publisher"`
	Author    string `json:"author"`
	Link      string `json:"link"`
}

func main() {
	jsonFile, err := os.Open(`book.json`)
	defer jsonFile.Close()

	if err != nil {
		fmt.Println(err)
	}

	byteValueJSON, _ := ioutil.ReadAll(jsonFile)

	objBook := Book{}
	json.Unmarshal(byteValueJSON, &objBook)

	fmt.Println(objBook.Name)

}
