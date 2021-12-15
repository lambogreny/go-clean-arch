package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/augusto/imersao5-esquenta-go/entity"
	"github.com/augusto/imersao5-esquenta-go/utils"
	"github.com/satori/go.uuid"
	"log"
	"strings"
	"time"
)

type TransactionRepositoryDb struct {
	db *sql.DB
}

func NewTransactionRepositoryDb(db *sql.DB) *TransactionRepositoryDb {
	return &TransactionRepositoryDb{db: db}
}

func (t *TransactionRepositoryDb) DeleteTransaction(id string) error {
	ctx := context.Background()
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		utils.LogFile("ERROR", " transaction", "CRITICAL ", err.Error(), "Erro ao iniciar a transação")
		return err
	}

	_, err = tx.ExecContext(ctx, "DELETE FROM transactions")
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, "insert into transactions values ('randomID',1,123,'approved','',current_timestamp,current_timestamp) ")
	//_, err = tx.ExecContext(ctx, "insert into transactions values ('randomID',1,123,'approved','',current_timestamp,) ") //Fornçando o erro

	if err != nil {
		tx.Rollback()
		return err
	}

	//Comitando a transação
	err = tx.Commit()

	if err != nil {
		utils.LogFile("ERROR", " transaction", "CRITICAL ", err.Error(), "Erro ao commitar a transação")
		return err
	}
	log.Println("Transaction commited!")

	return nil

}

func (t *TransactionRepositoryDb) Select(id string) ([]entity.Transaction, error) {

	var filterParam string

	if id != "" {
		fmt.Println("Chegou no repositório um id : ", id)
		filterParam = fmt.Sprintf("WHERE id = '%s' ", id)
	}

	queryString := fmt.Sprintf("SELECT id,account_id,amount,status,error_message FROM transactions %s", filterParam)
	rows, err := t.db.Query(queryString)

	if err != nil {
		utils.LogFile("ERROR", " transaction", "CRITICAL ", err.Error(), queryString)
		log.Println("could not execute query: %v", err) //Mata a aplicação

		return nil, err
		//return []entity.Transaction{}, err
	}

	//Lista com todas as transações
	transactions := []entity.Transaction{}

	for rows.Next() {

		transaction := entity.Transaction{}

		if err := rows.Scan(&transaction.ID, &transaction.AccountID, &transaction.Amount, &transaction.Status, &transaction.ErrorMessage); err != nil {
			if strings.Contains(err.Error(), "table") {
				fmt.Println("Tabela não encontrada!")
			}
			//log.Fatalf("could not scan row: %v", err)

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
