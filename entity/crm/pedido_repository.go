package crm

import "github.com/augusto/imersao5-esquenta-go/entity/crm/pedido"

/**
 * Interface que implementa os métodos relacionados ao pedido
 */

type PedidoRepository interface {
	SelectAccount(owner string) ([]pedido.Quote, error)
}
