package repository

import (
	"database/sql"
	"time"
)

type TransactionRepositoryDb struct {
	db *sql.DB
}

func NewTransactionRepositoryDb(db *sql.DB) *TransactionRepositoryDb {
	return &TransactionRepositoryDb{db: db}
}

//func (t *TransactionRepositoryDb) Select() error {
//	rows, err := t.db.Query("SELECT * FROM transactions")
//
//	if err != nil {
//		log.Fatalf("could not execute query: %v", err)
//		return err
//	}
//
//	//Lista com todas as transações
//	transactions := []entity.Transaction{}
//
//	for rows.Next() {
//		transaction := entity.Transaction{}
//
//		if err := rows.Scan(&transaction.ID, &transaction.AccountID, &transaction.Amount, &transaction.Status, &transaction.ErrorMessage, &transaction.CreatedAt, &transaction.UpdatedAt); err != nil {
//			log.Fatalf("could not scan row: %v", err)
//			return err
//		}
//		transactions = append(transactions, transaction)
//		fmt.Println(transaction)
//
//	}
//
//	return nil
//
//}

func (t *TransactionRepositoryDb) Insert(id string, account string, amount float64, status string, errorMessage string) error {
	stmt, err := t.db.Prepare(`
		insert into transactions (id, account_id, amount, status, error_message, created_at, updated_at)
		values($1,$2,$3,$4,$5,$6,$7)
		`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		id,
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
