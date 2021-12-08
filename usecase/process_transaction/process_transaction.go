package process_transaction

import (
	"fmt"
	"github.com/augusto/imersao5-esquenta-go/entity"
)

type ProcessTransacion struct {
	Repository entity.TransactionRepository
}

func NewProcessTransaction(repository entity.TransactionRepository) *ProcessTransacion {
	return &ProcessTransacion{Repository: repository}
}

func (p *ProcessTransacion) Execute(input TransactionDtoInput) (TransactionDtoOutput, error) {
	fmt.Println("Starting transaction...")
	transaction := entity.NewTransaction()
	transaction.ID = input.ID
	transaction.AccountID = input.AccountID
	transaction.Amount = input.Amount
	invalidTransaction := transaction.IsValid()

	fmt.Println("Validating transaction....")
	if invalidTransaction != nil {
		fmt.Println("Invalid transaction!")
		return p.rejectTransaction(transaction, invalidTransaction)
	}
	fmt.Println("Valid transaction!")
	return p.approveTransaction(transaction, invalidTransaction)

}

func (p *ProcessTransacion) approveTransaction(transaction *entity.Transaction, invalidTransaction error) (TransactionDtoOutput, error) {
	err := p.Repository.Insert(transaction.ID, transaction.AccountID, transaction.Amount, "approved", "")
	if err != nil {
		return TransactionDtoOutput{}, err
	}
	output := TransactionDtoOutput{
		ID:           transaction.ID,
		Status:       "approved",
		ErrorMessage: "",
	}
	return output, nil
}

func (p *ProcessTransacion) rejectTransaction(transaction *entity.Transaction, invalidTransaction error) (TransactionDtoOutput, error) {
	err := p.Repository.Insert(transaction.ID, transaction.AccountID, transaction.Amount, "rejected", invalidTransaction.Error())
	if err != nil {
		return TransactionDtoOutput{}, err
	}
	output := TransactionDtoOutput{
		ID:           transaction.ID,
		Status:       "rejected",
		ErrorMessage: invalidTransaction.Error(),
	}
	return output, nil
}
