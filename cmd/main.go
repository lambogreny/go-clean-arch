package main

import (
	"github.com/augusto/imersao5-esquenta-go/adapter/api"
	_ "github.com/mattn/go-sqlite3"
)

//Examplo inserindo o dado na mão
//func main() {
//
//	var amount float64 = 5
//
//	db, err := sql.Open("sqlite3", "test.db")
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	repo := repository.NewTransactionRepositoryDb(db)
//	usecase := process_transaction.NewProcessTransaction(repo)
//
//	input := process_transaction.TransactionDtoInput{
//		ID:        "1",
//		AccountID: "1",
//		Amount:    amount,
//	}
//
//	output, err := usecase.Execute(input)
//
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	fmt.Println(output)
//
//	outputJson, _ := json.Marshal(output)
//
//	fmt.Println(string(outputJson))
//}

func main() {

	//db, err := sql.Open("sqlite3", "test.db")
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//repo := repository.NewTransactionRepositoryDb(db)
	//fmt.Println(repo)

	//Iniciar a api Echo
	//webserver := api.NewWebServer()
	//webserver.Repository = repo
	//webserver.Serve()

	//Inicia a api Gin
	api.Init()
}
