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

func (p *ProcessPrd) UseCaseSelect() ([]prd.Prd, error) {

	resp, err := p.Repository.Select()

	if err != nil {
		return []prd.Prd{}, err
	}

	return resp, nil

}

func (p *ProcessPrd) UseCaseUpdate(prd prd.Prd, owner string) error {

	err := p.Repository.Update(prd, owner)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProcessPrd) UseCaseDelete(codigoProduto string, tipo string) error {

	err := p.Repository.Delete(codigoProduto, tipo)

	if err != nil {
		return err
	}

	return nil
}

func (p *ProcessPrd) UseCaseInsert(prd prd.Prd, owner string) error {
	err := p.Repository.Insert(prd, owner)

	if err != nil {
		return err
	}
	return nil
}

func (p *ProcessPrd) UseCaseCheckUpdateCrm(codigoProduto string) (bool, error) {
	resp, err := p.Repository.CheckUpdateCrm(codigoProduto)

	if err != nil {
		return false, err
	}

	return resp, nil
}
