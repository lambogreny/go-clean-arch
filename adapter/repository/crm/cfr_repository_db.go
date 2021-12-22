package crmRepository

import (
	"database/sql"
	"github.com/augusto/imersao5-esquenta-go/entity/crm/cfr"
)

type CfrRepositoryDbErp struct {
	db *sql.DB
}

func NewCfrRepositoryDbErp(db *sql.DB) *CfrRepositoryDbErp {
	return &CfrRepositoryDbErp{db: db}
}

func (t CfrRepositoryDbErp) Select() ([]cfr.Cfr, error) {

	return []cfr.Cfr{}, nil
}
