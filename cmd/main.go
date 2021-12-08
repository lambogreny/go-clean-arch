package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/augusto/imersao5-esquenta-go/adapter/repository"
	"github.com/augusto/imersao5-esquenta-go/usecase/process_transaction"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {

	var amount float64 = 0

	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewTransactionRepositoryDb(db)
	usecase := process_transaction.NewProcessTransaction(repo)

	input := process_transaction.TransactionDtoInput{
		ID:        "1",
		AccountID: "1",
		Amount:    amount,
	}

	output, err := usecase.Execute(input)

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(output)

	outputJson, _ := json.Marshal(output)

	fmt.Println(string(outputJson))
}
