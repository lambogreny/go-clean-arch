package process_approval

import (
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

	var user string = "ELISA"

	resp, err := p.Repository.Select(user)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
