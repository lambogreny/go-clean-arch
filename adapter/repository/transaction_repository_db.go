package repository

import (
	"database/sql"
	"fmt"
	"github.com/augusto/imersao5-esquenta-go/entity"
	"github.com/augusto/imersao5-esquenta-go/utils"
	"github.com/satori/go.uuid"
	"log"
	"time"
)

type TransactionRepositoryDb struct {
	db *sql.DB
}

func NewTransactionRepositoryDb(db *sql.DB) *TransactionRepositoryDb {
	return &TransactionRepositoryDb{db: db}
}

func (t *TransactionRepositoryDb) Select() ([]entity.Transaction, error) {
	queryString := "SELECT id,account_id,amount,status,error_message FROM transactions"
	rows, err := t.db.Query(queryString)

	if err != nil {
		utils.LogFile("ERROR", " transaction", "CRITICAL ", "Problema de execução na query (select do GET ALL)", queryString)
		log.Fatalf("could not execute query: %v", err)

		return nil, err
	}

	//Lista com todas as transações
	transactions := []entity.Transaction{}

	for rows.Next() {

		transaction := entity.Transaction{}

		if err := rows.Scan(&transaction.ID, &transaction.AccountID, &transaction.Amount, &transaction.Status, &transaction.ErrorMessage); err != nil {
			log.Fatalf("could not scan row: %v", err)
			return nil, err
		}
		fmt.Println("O id do banco é : ", transaction.ID)
		transactions = append(transactions, transaction)
		fmt.Println("Aqui a transação : ", transaction)

	}

	return transactions, nil

}

func (t *TransactionRepositoryDb) Insert(account string, amount float64, status string, errorMessage string) error {
	u1 := uuid.NewV4()
	//u1 := "12312"

	stmt, err := t.db.Prepare(`
		insert into transactions (id, account_id, amount, status, error_message, created_at, updated_at)
		values($1,$2,$3,$4,$5,$6,$7)
		`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		u1,
		account,
		amount,
		status,
		errorMessage,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}
	return nil
}
