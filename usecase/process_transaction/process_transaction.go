package process_transaction

import (
	"log"

	"github.com/augusto/imersao5-esquenta-go/entity"
)

type ProcessTransacion struct {
	Repository entity.TransactionRepository
}

func NewProcessTransaction(repository entity.TransactionRepository) *ProcessTransacion {
	return &ProcessTransacion{Repository: repository}
}

func (p *ProcessTransacion) DeleteTransaction() error {
	log.Println("Starting delete transactions...")

	//TODO fazer um select aqui para recuperar uma transação específica
	var idFake string = "a7662337-d0b3-43f8-b815-ce45525f7eea"

	err := p.Repository.DeleteTransaction(idFake)

	if err != nil {
		return err
	}

	return nil

}

func (p *ProcessTransacion) GetAll(id string) ([]entity.Transaction, error) {
	log.Println("Getting transaction...")

	resp, err := p.Repository.Select(id)

	if err != nil {
		return nil, err
	}
	return resp, nil

}

func (p *ProcessTransacion) Execute(input TransactionDtoInput) (TransactionDtoOutput, error) {
	log.Println("Starting transaction...")
	transaction := entity.NewTransaction()
	transaction.AccountID = input.AccountID
	transaction.Amount = input.Amount
	invalidTransaction := transaction.IsValid()

	log.Println("Validating transaction....")
	if invalidTransaction != nil {
		log.Println("Invalid transaction!")
		return p.rejectTransaction(transaction, invalidTransaction)
	}
	log.Println("Valid transaction!")
	return p.approveTransaction(transaction, invalidTransaction)

}

func (p *ProcessTransacion) approveTransaction(transaction *entity.Transaction, invalidTransaction error) (TransactionDtoOutput, error) {
	err := p.Repository.Insert(transaction.AccountID, transaction.Amount, "approved", "")
	if err != nil {
		//return TransactionDtoOutput{}, err
		return TransactionDtoOutput{}, err
	}
	output := TransactionDtoOutput{
		Status:       "approved",
		ErrorMessage: "",
	}
	return output, nil
}

func (p *ProcessTransacion) rejectTransaction(transaction *entity.Transaction, invalidTransaction error) (TransactionDtoOutput, error) {
	//Não insere no banco se nao for válido

	output := TransactionDtoOutput{
		Status:       "rejected",
		ErrorMessage: invalidTransaction.Error(),
	}
	return output, nil
}
