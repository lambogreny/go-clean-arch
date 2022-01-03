package crmRepository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/augusto/imersao5-esquenta-go/entity/crm/pedido"
	"github.com/augusto/imersao5-esquenta-go/utils"
)

type PedidoRepositoryDbErp struct {
	db *sql.DB
}

func NewPedidoRepositoryDbErp(db *sql.DB) *PedidoRepositoryDbErp {
	return &PedidoRepositoryDbErp{db: db}
}

// ---------------------------------------- CRM -> ERP --------------------------//
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
		// utils.LogFile("ERROR", " pedido", "CRITICAL ", err.Error(), queryString)
		utils.LogDatabaseDetails("PEDIDO", "SELECT", queryString, err.Error(), "")
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
			utils.LogDatabaseDetails("PEDIDO", "SCAN", queryString, err.Error(), "")
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
		// log.Println("could not execute query:", err)
		// utils.LogFile("ERROR", " pedido", "CRITICAL ", err.Error(), queryString)
		utils.LogDatabaseDetails("PEDIDO", "SELECT", queryString, err.Error(), "")

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
			utils.LogDatabaseDetails("PEDIDO", "SCAN", queryString, err.Error(), "")
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
		// utils.LogFile("CRM/PEDIDO", " delete", "CRITICAL ", quoteErr.Error(), quoteQueryString)
		utils.LogDatabaseDetails("PEDIDO", "delete", quoteQueryString, quoteErr.Error(), "")
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
		// utils.LogFile("CRM/PEDIDO", " delete", "CRITICAL ", quoteItemErr.Error(), quoteQueryString)
		utils.LogDatabaseDetails("PEDIDO", "delete", quoteItemQueryString, quoteItemErr.Error(), "")
		tx.Rollback()
		return quoteItemErr
	}

	//commit := tx.Commit()

	//if commit != nil {
	//	log.Println("Erro ao realizar o commit de delete")
	//	utils.LogFile("CRM/PEDIDO", " delete", "CRITICAL ", commit.Error(), "Commit da transação")
	//	return commit
	//}

	return nil
}

//------------------------------------ ERP para CRM -------------------------------//

func (t PedidoRepositoryDbErp) SelectCpv() ([]pedido.Cpv, error) {

	var controlLimitQuery string = "LIMIT 10"

	queryString := utils.Msg(`SELECT 
								tipo,
								numero,
								pedido_fil,
								cliente,
								codigo_operacao,
								filial,
								data_entrada,
								data_entrega,
								finalidade,
								ped_cliente,
								cond_pagto,
								valor_mercadorias,
								valor_total,
								tipo_frete,
								peso_liquido,
								peso_bruto,
								entrega_end,
								entrega_bairro,
								entrega_cep,
								entrega_cidade,
								entrega_uf,
								transp,
								usuario_inclusao,
								usuario_alteracao,
								data_hora_inclusao,
								data_hora_alteracao,
								status,
								tipo_abordagem,
								unidade_negocio,
								probab_fech,
								meio_conhec,
								emissao,
								obs_simples,
								cod_representante
								from cpv
								inner join tb_crm_sincroniza ON numero = pk and tabela = 'CPV'
								{{.controlLimitQuery}}
							`, map[string]interface{}{"controlLimitQuery": controlLimitQuery})

	rows, err := t.db.Query(queryString)

	if err != nil {
		utils.LogDatabaseDetails("CPV", "SELECT", queryString, err.Error(), "")
	}

	cpvs := []pedido.Cpv{}

	for rows.Next() {
		cpv := pedido.Cpv{}

		if err := rows.Scan(&cpv.Tipo,
			&cpv.Numero,
			&cpv.PedidoFil,
			&cpv.Cliente,
			&cpv.CodigoOperacao,
			&cpv.Filial,
			&cpv.DataEntrada,
			&cpv.DataEntrega,
			&cpv.Finalidade,
			&cpv.PedCliente,
			&cpv.CondPagamento,
			&cpv.ValorMercadorias,
			&cpv.ValorTotal,
			&cpv.TipoFrete,
			&cpv.PesoLiquido,
			&cpv.PesoBruto,
			&cpv.EntregaEnd,
			&cpv.EntregaBairro,
			&cpv.EntregaCep,
			&cpv.EntregaCidade,
			&cpv.EntregaUf,
			&cpv.Transp,
			&cpv.UsuarioInclusao,
			&cpv.UsuarioAlteracao,
			&cpv.DataHoraInclusao,
			&cpv.DataHoraAlteracao,
			&cpv.Status,
			&cpv.TipoAbordagem,
			&cpv.UnidadeNegocio,
			&cpv.ProbabFech,
			&cpv.MeioConhec,
			&cpv.Emissao,
			&cpv.ObsSimples,
			&cpv.CodRepresentante,
		); err != nil {
			log.Println(err.Error())
			utils.LogDatabaseDetails("CPV", "SCAN", queryString, err.Error(), "")
			return nil, err
		}

		cpvs = append(cpvs, cpv)
	}

	return cpvs, nil
}

func (t PedidoRepositoryDbErp) CheckUpdateCrm(id string, owner string) (bool, error) {
	queryString := fmt.Sprintf(`SELECT count(*) from %s.quote where id= '%s'`, owner, id)
	fmt.Println(queryString)

	rows, err := t.db.Query(queryString)

	if err != nil {
		utils.LogDatabaseDetails("CPV", id, queryString, err.Error(), "")
		return false, err
	}

	var count int

	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			utils.LogDatabaseDetails("CPV", id, queryString, err.Error(), "")
			return false, err
		}
	}

	if count == 0 {
		return false, nil
	}
	return true, nil

}

func (t PedidoRepositoryDbErp) UpdateCrm(cpv pedido.Cpv, owner string) error {
	queryString := utils.Msg(`UPDATE {{.owner}}.quote SET
									number_a  = '{el['numero']}',
									pedidofil  = {el['pedido_fil']},
									account_id  = {el['cliente']},
									codigooperacao  = {el['codigo_operacao']},
									filial  = {el['filial']},
									date_quoted  = {el['data_entrada']},
									dataentrega  = {el['data_entrega']},
									finalidade  = {el['finalidade']},
									pedcliente  = {el['ped_cliente']},
									condpagto  = {el['cond_pagto']},
									valormercadorias  = {el['valor_mercadorias']},
									valortotal  = {el['valor_total']},
									tipofrete  = {el['tipo_frete']},
									pesoliquido  = {el['peso_liquido']},
									pesobruto  = {el['peso_bruto']},
									shipping_address_street  = {el['entrega_end']},
									entregabairro  = {el['entrega_bairro']},
									shipping_address_postal_code  = {el['entrega_cep']},
									shipping_address_state  = {el['entrega_uf']},
									transp  = {el['transp']},
									created_by_id  = {el['usuario_inclusao']},
									modified_by_id  = '{variables['usuario_alteracao']}',
									created_at  = {el['data_hora_inclusao']},
									modified_at  = {el['data_hora_alteracao']},
									status  = {el['status']},
									tipoabordagem  = {el['tipo_abordagem']},
									unidadenegocio  = {el['unidade_negocio']},
									probabfech  = {el['probab_fech']},
									meioconhec  = {el['meio_conhec']},
									emissao  = {el['emissao']}, 
									obssimples  = {el['obs_simples']},
									codrepresentante  = {el['cod_representante']}
									WHERE id = {el['numero']}`, map[string]interface{}{
		"owner": owner,
	})

	fmt.Println(queryString)
	return nil
}

func (t PedidoRepositoryDbErp) InsertCrm(pedido.Cpv) error {
	return nil
}

func (t PedidoRepositoryDbErp) DeleteErp(id string) error {
	return nil
}
