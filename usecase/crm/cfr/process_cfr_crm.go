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

func (p *ProcessCfr) UserCaseCheckUpdateErp(id string) (bool, error) {
	resp, err := p.Repository.CheckUpdateErp(id)

	if err != nil {
		return false, err
	}

	return resp, nil
}

func (p *ProcessCfr) UseCaseSelect(owner string) ([]cfr.Account, error) {

	resp, err := p.Repository.SelectCrm(owner)

	if err != nil {
		return []cfr.Account{}, err
	}
	return resp, nil
}

func (p *ProcessCfr) UseCaseUpdate(account cfr.Account, owner string) error {
	err := p.Repository.UpdateErp(account, owner)

	if err != nil {
		return err
	}

	return nil
}

func (p *ProcessCfr) UseCaseInsert(account cfr.Account, owner string) error {
	err := p.Repository.InsertErp(account, owner)

	if err != nil {
		return err
	}
	return nil
}

func (p *ProcessCfr) UseCaseDelete(id string, owner string) error {

	err := p.Repository.DeleteCrm(id, owner)

	if err != nil {
		return err
	}
	return nil
}
