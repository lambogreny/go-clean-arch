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

func (p *ProcessTransacion) GetAll() (entity.Transaction, error) {
	log.Println("Getting transaction...")

	// err := p.Repository.Select()

	return entity.Transaction{}, nil

}

func (p *ProcessTransacion) Execute(input TransactionDtoInput) (TransactionDtoOutput, error) {
	log.Println("Starting transaction...")
	transaction := entity.NewTransaction()
	transaction.ID = input.ID
	transaction.AccountID = input.AccountID
	transaction.Amount = input.Amount
	invalidTransaction := transaction.IsValid()

	//Chamando a função de teste
	transaction.Teste()

	log.Println("Validating transaction....")
	if invalidTransaction != nil {
		log.Println("Invalid transaction!")
		return p.rejectTransaction(transaction, invalidTransaction)
	}
	log.Println("Valid transaction!")
	return p.approveTransaction(transaction, invalidTransaction)

}

func (p *ProcessTransacion) approveTransaction(transaction *entity.Transaction, invalidTransaction error) (TransactionDtoOutput, error) {
	err := p.Repository.Insert(transaction.ID, transaction.AccountID, transaction.Amount, "approved", "")
	if err != nil {
		//return TransactionDtoOutput{}, err
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
	//Não insere no banco se nao for válido
	//err := p.Repository.Insert(transaction.ID, transaction.AccountID, transaction.Amount, "rejected", invalidTransaction.Error())
	//if err != nil {
	//	return TransactionDtoOutput{}, err
	//}
	output := TransactionDtoOutput{
		ID:           transaction.ID,
		Status:       "rejected",
		ErrorMessage: invalidTransaction.Error(),
	}
	return output, nil
}
