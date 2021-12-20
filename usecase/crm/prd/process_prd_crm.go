package prd

import (
	"github.com/augusto/imersao5-esquenta-go/entity/crm"
	"github.com/augusto/imersao5-esquenta-go/entity/crm/prd"
)

type ProcessPrd struct {
	Repository crm.PrdRepository
}

func NewProcessPrd(repository crm.PrdRepository) *ProcessPrd {
	return &ProcessPrd{Repository: repository}
}

func Select(p *ProcessPrd) ([]prd.Prd, error) {

	resp, err := p.Repository.Select()
	if err != nil {
		return []prd.Prd{}, err
	}

	return resp, nil
}
