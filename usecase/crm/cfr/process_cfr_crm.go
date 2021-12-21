package cfr

import (
	"github.com/augusto/imersao5-esquenta-go/entity/crm"
	"github.com/augusto/imersao5-esquenta-go/entity/crm/cfr"
)

type ProcessCfr struct {
	Repository crm.CfrRepository
}

func NewProcessCfr(repository crm.CfrRepository) *ProcessCfr {
	return &ProcessCfr{Repository: repository}
}

func (p *ProcessCfr) UseCaseSelect() ([]cfr.Cfr, error) {

	resp, err := p.Repository.Select()

	if err != nil {
		return []cfr.Cfr{}, err
	}
	return []cfr.Cfr{}, nil
}
