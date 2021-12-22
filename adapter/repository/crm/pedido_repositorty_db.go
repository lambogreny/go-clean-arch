package crmRepository

import (
	"database/sql"
	"fmt"
	"github.com/augusto/imersao5-esquenta-go/entity/crm/pedido"
	"github.com/augusto/imersao5-esquenta-go/utils"
	"log"
)

type PedidoRepositoryDbErp struct {
	db *sql.DB
}

func NewPedidoRepositoryDbErp(db *sql.DB) *PedidoRepositoryDbErp {
	return &PedidoRepositoryDbErp{db: db}
}

func (t PedidoRepositoryDbErp) SelectAccount(owner string) ([]pedido.Quote, error) {
	queryString := utils.Msg(`SELECT id,
									   Upper(filial),
									   Upper(account_id),
									   (SELECT ti9codigo
										FROM   {{.owner}}.account
										WHERE  id = quote.account_id) AS ti9codigo,
									   Upper(codigooperacao),
									   Upper(finalidade),
									   date_quoted,
									   1                              AS moeda,
									   CURRENT_DATE()                 AS dataentrada,
									   dataentrega,
									   Upper(condpagto),
									   indpres,
									   origemregistro,
									   '9',
									   'CRM',
									   Upper(obssimples),
									   Upper(obsfiscal),
									   Upper(outrasinfcom),
									   id                             AS obspalm,
									   amount,
									   descfinanc,
									   codcobranca,
									   pedcliente,
									   Upper(descfinanc),
									   '',
									   Upper(LEFT(tipofrete, 1))      AS tipofrete,
									   Upper(transp),
									   0,
									   '',
									   '',
									   '',
									   Upper(codrepresentante),
									   id
								FROM   {{.owner}}.quote
									   INNER JOIN {{.owner}}.tb_crm_sincroniza
											   ON id = pk
												  AND tabela = 'CPV'
								WHERE  Ifnull(usuario, 'xx') <> 'CRM'
									   AND liberado = 1
									   AND integrado = 0 `, map[string]interface{}{
		"owner": owner,
	})
	fmt.Println(queryString)
	rows, err := t.db.Query(queryString)
	fmt.Println(rows)

	if err != nil {
		utils.LogFile("ERROR", " pedido", "CRITICAL ", err.Error(), queryString)
		log.Println("could not execute query:", err)

		return []pedido.Quote{}, err
	}

	return []pedido.Quote{}, nil
}
