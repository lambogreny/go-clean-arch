package pedido

import "github.com/augusto/imersao5-esquenta-go/entity/crm"

type ProcessPedido struct {
	Repository crm.PedidoRepository
}

func NewProcessPedido(repository crm.PedidoRepository) *ProcessPedido {
	return &ProcessPedido{Repository: repository}
}

func (p *ProcessPedido) UseCaseSelect() error {

	//Faver o loop e devolver o objeto
	return nil
}
