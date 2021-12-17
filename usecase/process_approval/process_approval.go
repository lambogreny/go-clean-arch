package process_approval

import (
	"fmt"
	"log"

	"github.com/augusto/imersao5-esquenta-go/entity"
)

type ProcessApproval struct {
	Repository entity.ApprovalRepository
}

func NewApprovalTransaction(repository entity.ApprovalRepository) *ProcessApproval {
	return &ProcessApproval{Repository: repository}
}

func (p *ProcessApproval) Interact(input ApprovalDtoInteractionInput) error {

	//Validando a permissão
	hasPermission, err := p.Repository.CheckPermission(input.Usuario)

	if err != nil {
		return err
	}
	if hasPermission == false {
		return fmt.Errorf("Usuário não autorizado!")
	}

	//Chamando a transação de insert
	insertTransaction := p.Repository.Interact(input.Filial, input.Cotacao, input.Fornecedor, input.TipoDeAprovacao, input.Usuario, input.StatusDeAprovacao, input.Justificativa, input.SeqConcatenada)

	if insertTransaction != nil {
		return insertTransaction
	}

	return nil
}

func (p *ProcessApproval) GetAll() ([]entity.Approval, error) {
	log.Println("Getting approvals...")

	var user string = "DONATTI"

	resp, err := p.Repository.Select(user)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
