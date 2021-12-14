package process_approval

import (
	"fmt"
	"github.com/augusto/imersao5-esquenta-go/entity"
	"log"
)

type ProcessApproval struct {
	Repository entity.ApprovalRepository
}

func NewApprovalTransaction(repository entity.ApprovalRepository) *ProcessApproval {
	return &ProcessApproval{Repository: repository}
}

func (p *ProcessApproval) GetAll() ([]entity.Approval, error) {
	log.Println("Getting approvals...")

	resp, err := p.Repository.Select()

	if err != nil {
		fmt.Println("ERRO AQUI")
	}

	return resp, nil
}
