package crmRepository

import (
	"context"
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

func (t PedidoRepositoryDbErp) SelectQuote(owner string) ([]pedido.Quote, error) {
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
									   Upper(obssimples),
									   Upper(obsfiscal),
									   Upper(outrasinfcom),
									   id                             AS obspalm,
									   amount,
									   Upper(descfinanc),
									   codcobranca,
									   pedcliente,
									   Upper(LEFT(tipofrete, 1))      AS tipofrete,
									   Upper(transp),
									   Upper(codrepresentante)
								FROM   {{.owner}}.quote
									   INNER JOIN {{.owner}}.tb_crm_sincroniza
											   ON id = pk
												  AND tabela = 'CPV'
								WHERE  Ifnull(usuario, 'xx') <> 'CRM'
									   AND liberado = 1
									   AND integrado = 0 `, map[string]interface{}{
		"owner": owner,
	})
	//fmt.Println(queryString)
	rows, err := t.db.Query(queryString)

	if err != nil {
		utils.LogFile("ERROR", " pedido", "CRITICAL ", err.Error(), queryString)
		log.Println("could not execute query:", err)

		return []pedido.Quote{}, err
	}

	pedidos := []pedido.Quote{}

	for rows.Next() {
		pedido := pedido.Quote{}

		if err := rows.Scan(&pedido.Id,
			&pedido.Filial,
			&pedido.Account_id,
			&pedido.Ti9Codigo,
			&pedido.CodigoOperacao,
			&pedido.Finalidade,
			&pedido.Data_quoted,
			&pedido.Moeda,
			&pedido.DataEntrada,
			&pedido.DataEntrega,
			&pedido.CondPagamento,
			&pedido.Indpres,
			&pedido.OrigemRegistro,
			&pedido.ObsSimples,
			&pedido.ObsFiscal,
			&pedido.OutrasInfoCom,
			&pedido.ObsPalm,
			&pedido.Amount,
			&pedido.DescFinanc,
			&pedido.CodCobranca,
			&pedido.PedCliente,
			&pedido.TipoFrete,
			&pedido.Transp,
			&pedido.CodRepresentante,
		); err != nil {
			log.Println(err.Error())
			return nil, err
		}
		pedidos = append(pedidos, pedido)
	}

	return pedidos, nil
}

func (t PedidoRepositoryDbErp) SelectQuoteItem(owner string, id string) ([]pedido.QuoteItem, error) {
	queryString := utils.Msg(`SELECT itempv,
									   Upper(product_id),
									   Upper(almoxarifado),
									   dataentregaitem,
									   Upper(quote_item.finalidade),
									   qtdepedida,
									   numeropc,
									   itempc,
									   codigocentro,
									   quantity,
									   unit_price,
									   discount,
									   Upper(quote_id),
									   data_entrega_item,
									   data_entrega,
									   quote_item.amount,
									   quote_item.order,
									   quote_item.description                 AS descricao_informada,
									   '1'                                    AS fator_conv2,
									   'CRM'                                  AS usuario_inclusao,
									   ( quote_item.amount * discount ) / 100 AS desc_valor
								FROM   quote_item
									   INNER JOIN {{.owner}}.quote
											   ON ( {{.owner}}.quote_item.quote_id =
															   {{.owner}}.quote.id )
								WHERE  quote_id = '{{.id}}'
									   AND quote_item.deleted = 0  `, map[string]interface{}{
		"owner": owner,
		"id":    id,
	})
	//fmt.Println(queryString)
	rows, err := t.db.Query(queryString)

	if err != nil {
		log.Println("could not execute query:", err)
		utils.LogFile("ERROR", " pedido", "CRITICAL ", err.Error(), queryString)

		return []pedido.QuoteItem{}, err
	}

	items := []pedido.QuoteItem{}

	for rows.Next() {
		item := pedido.QuoteItem{}

		if err := rows.Scan(&item.ItemPv,
			&item.ProductId,
			&item.Almoxarifado,
			&item.DataEntregaItem,
			&item.Finalidade,
			&item.QtdePedida,
			&item.NumeroPc,
			&item.ItemPc,
			&item.CodigoCentro,
			&item.Quantity,
			&item.UnitPrice,
			&item.Discount,
			&item.QuoteId,
			&item.DataEntregaItem,
			&item.DataEntrega,
			&item.Amount,
			&item.Order,
			&item.DescricaoInformada,
			&item.FatorConv2,
			&item.UsuarioInclusao,
			&item.DescValor,
		); err != nil {
			log.Println(err.Error())
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (t PedidoRepositoryDbErp) DeleteSincroniza(owner string, id string) error {
	fmt.Println("Recebi o comando de delete no owner", owner, "e com o id", id)

	//Iniciando o contexto de transação
	ctx := context.Background()
	tx, _ := t.db.BeginTx(ctx, nil)

	quoteQueryString := utils.Msg(`DELETE FROM {{.owner}}.tb_crm_sincroniza
						WHERE pk = '{{.id}}' and tipo IN ('I','U') and tabela = 'CPV'`,
		map[string]interface{}{
			"owner": owner,
			"id":    id,
		},
	)

	_, quoteErr := tx.ExecContext(ctx, quoteQueryString)

	if quoteErr != nil {
		utils.LogFile("CRM/PEDIDO", " delete", "CRITICAL ", quoteErr.Error(), quoteQueryString)
		tx.Rollback()
		return quoteErr
	}

	quoteItemQueryString := utils.Msg(`DELETE FROM {{.owner}}.tb_crm_sincroniza
						WHERE pk = '{{.id}}' and tipo IN ('I','U') and tabela = 'IPV'`,
		map[string]interface{}{
			"owner": owner,
			"id":    id,
		},
	)

	_, quoteItemErr := tx.ExecContext(ctx, quoteItemQueryString)

	if quoteItemErr != nil {
		utils.LogFile("CRM/PEDIDO", " delete", "CRITICAL ", quoteItemErr.Error(), quoteQueryString)
		tx.Rollback()
		return quoteItemErr
	}

	//commit := tx.Commit()

	//if commit != nil {
	//	log.Println("Erro ao realizar o commit de delete")
	//	utils.LogFile("CRM/PEDIDO", " delete", "CRITICAL ", commit.Error(), "Commit da transação")
	//	return commit
	//}

	log.Println("Consegui executar tods os deletes")
	return nil
}
