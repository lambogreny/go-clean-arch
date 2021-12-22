package pedido

import (
	"github.com/augusto/imersao5-esquenta-go/entity/crm"
	"github.com/augusto/imersao5-esquenta-go/entity/crm/pedido"
)

type ProcessPedido struct {
	Repository crm.PedidoRepository
}

func NewProcessPedido(repository crm.PedidoRepository) *ProcessPedido {
	return &ProcessPedido{Repository: repository}
}

func (p *ProcessPedido) UseCaseSelect(owner string) ([]pedido.Quote, error) { //Mudar aqui

	resp, err := p.Repository.SelectAccount(owner)

	if err != nil {
		return []pedido.Quote{}, err
	}

	//Nessa caso de uso montar a struct do payload
	//Faver o loop e devolver o objeto
	return resp, nil
}
