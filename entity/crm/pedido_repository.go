package crm

import "github.com/augusto/imersao5-esquenta-go/entity/crm/pedido"

/**
 * Interface que implementa os métodos relacionados ao pedido
 */
type PedidoRepository interface {
	SelectQuote(owner string) ([]pedido.Quote, error)
	SelectQuoteItem(owner string, id string) ([]pedido.QuoteItem, error)
	DeleteSincroniza(owner string, id string) error
}
