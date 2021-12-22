package crmRepository

import "database/sql"

type PedidoRepositoryDbErp struct {
	db *sql.DB
}

func NewPedidoRepositoryDbErp(db *sql.DB) *PedidoRepositoryDbErp {
	return &PedidoRepositoryDbErp{db: db}
}

func (t PedidoRepositoryDbErp) Select() error {
	return nil
}
