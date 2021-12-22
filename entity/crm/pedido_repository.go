package crm

/**
 * Interface que implementa os métodos relacionados ao pedido
 */

type PedidoRepository interface {
	Select() error
}
