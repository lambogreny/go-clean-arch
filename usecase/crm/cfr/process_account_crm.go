package cfr

import (
	"github.com/augusto/imersao5-esquenta-go/entity/crm"
	"github.com/augusto/imersao5-esquenta-go/entity/crm/cfr"
)

type ProcessAccount struct {
	Repository crm.AccountRepository
}

func NewProcessAccount(repository crm.AccountRepository) *ProcessAccount {
	return &ProcessAccount{Repository: repository}
}

func (p *ProcessAccount) UseCaseSelect() ([]cfr.Cfr, error) {

	resp, err := p.Repository.SelectErp()

	if err != nil {
		return nil, err
	}

	return resp, nil
}
