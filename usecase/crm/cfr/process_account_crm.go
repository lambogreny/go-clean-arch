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

func (p *ProcessAccount) UseCaseCheckUpdateCrm(id string, owner string) (bool, error) {

	resp, err := p.Repository.CheckUpdateCrm(id, owner)

	if err != nil {
		return false, err
	}
	return resp, nil
}

func (p *ProcessAccount) UseCaseUpdate(cfr cfr.Cfr, owner string) error {
	err := p.Repository.UpdateCrm(cfr, owner)

	if err != nil {
		return err
	}
	return nil
}

func (p *ProcessAccount) UseCaseSelect() ([]cfr.Cfr, error) {

	resp, err := p.Repository.SelectErp()

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *ProcessAccount) UseCaseDelete(id string, tipo string) error {
	err := p.Repository.DeleteErp(id, tipo)

	if err != nil {
		return err
	}
	return nil
}

func (p *ProcessAccount) UseCaseInsert(cfr cfr.Cfr, owner string) error {
	err := p.Repository.InsertCrm(cfr, owner)

	if err != nil {
		return err
	}
	return nil
}
