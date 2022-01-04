package crm

import "github.com/augusto/imersao5-esquenta-go/entity/crm/pedido"

/**
 * Interface que implementa os métodos relacionados ao pedido
 */
type PedidoRepository interface {
	//#CRM -> ERP
	SelectQuote(owner string) ([]pedido.Quote, error)
	SelectQuoteItem(owner string, id string) ([]pedido.QuoteItem, error)
	DeleteSincroniza(owner string, id string) error
	//# ERP -> CRM
	SelectCpv() ([]pedido.Cpv, error)
	CheckUpdateCrm(id string, owner string) (bool, error)
	UpdateCrm(cpv pedido.Cpv, owner string) error
	InsertCrm(cpv pedido.Cpv, owner string) error
	DeleteErp(id string, tipo string) error
}
