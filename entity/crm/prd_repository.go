package crm

import "github.com/augusto/imersao5-esquenta-go/entity/crm/prd"

/**
 * Interface que implementa os métodos relacionados a prd
 */

type PrdRepository interface {
	Select() ([]prd.Prd, error)
	CheckUpdateCrm(codigoProduto string) (bool, error)
	//Update()
	//Insert()
}
