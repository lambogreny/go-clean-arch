package crmRepository

import (
	"database/sql"
	"fmt"
	"github.com/augusto/imersao5-esquenta-go/entity/crm/cfr"
)

type CfrRepositoryDbErp struct {
	db *sql.DB
}

func NewCfrRepositoryDbErp(db *sql.DB) *CfrRepositoryDbErp {
	return &CfrRepositoryDbErp{db: db}
}

func (t CfrRepositoryDbErp) Select() ([]cfr.Cfr, error) {
	queryString := fmt.Sprintf(``)
	return []cfr.Cfr{}, nil
}
