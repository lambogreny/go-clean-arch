package crmRepository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/augusto/imersao5-esquenta-go/entity/crm/pedido"
	"github.com/augusto/imersao5-esquenta-go/utils"
	"github.com/augusto/imersao5-esquenta-go/utils/helpers"
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

//------------------------------------ ERP para CRM : CPV-------------------------------//

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
		return nil, err
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
									number_a  = '{{.number_a}}',
									pedidofil  = '{{.pedidofil}}',
									account_id  = '{{.account_id}}',
									codigooperacao  = '{{.codigooperacao}}',
									filial  = '{{.filial}}',
									date_quoted  = '{{.date_quoted}}',
									dataentrega  = '{{.dataentrega}}',
									finalidade  = '{{.finalidade}}',
									pedcliente  = '{{.pedcliente}}',
									condpagto  = '{{.condpagto}}',
									valormercadorias  = '{{.valormercadorias}}',
									valortotal  = '{{.valortotal}}',
									tipofrete  = '{{.tipofrete}}',
									pesoliquido  = '{{.pesoliquido}}',
									pesobruto  = '{{.pesobruto}}',
									shipping_address_street  = '{{.shipping_address_street}}',
									entregabairro  = '{{.entregabairro}}',
									shipping_address_postal_code  = '{{.shipping_address_postal_code}}',
									shipping_address_state  = '{{.shipping_address_state}}',
									transp  = '{{.transp}}',
									created_by_id  = '{{.created_by_id}}',
									modified_by_id  = '{{.modified_by_id}}',
									created_at  = '{{.created_at}}',
									modified_at  = '{{.modified_at}}',
									status  = '{{.status}}',
									tipoabordagem  = '{{.tipoabordagem}}',
									unidadenegocio  = '{{.unidadenegocio}}',
									probabfech  = '{{.probabfech}}',
									meioconhec  = '{{.meioconhec}}',
									emissao  = '{{.emissao}}',
									obssimples  = '{{.obssimples}}',
									codrepresentante  = '{{.codrepresentante}}'
									WHERE id = '{{.id}}'
									`, map[string]interface{}{
		"owner":                        owner,
		"number_a":                     helpers.String(cpv.Numero),
		"pedidofil":                    helpers.String(cpv.PedidoFil),
		"account_id":                   helpers.String(cpv.Cliente),
		"codigooperacao":               helpers.String(cpv.CodigoOperacao),
		"filial":                       helpers.String(cpv.Filial),
		"date_quoted":                  helpers.StringDatetime(cpv.DataEntrada),
		"dataentrega":                  helpers.StringDatetime(cpv.DataEntrega),
		"finalidade":                   helpers.String(cpv.Finalidade),
		"pedcliente":                   helpers.String(cpv.PedCliente),
		"condpagto":                    helpers.String(cpv.CondPagamento),
		"valormercadorias":             helpers.String(cpv.ValorMercadorias),
		"valortotal":                   helpers.String(cpv.ValorTotal),
		"tipofrete":                    helpers.String(cpv.TipoFrete),
		"pesoliquido":                  helpers.String(cpv.PesoLiquido),
		"pesobruto":                    helpers.String(cpv.PesoBruto),
		"shipping_address_street":      helpers.String(cpv.EntregaEnd),
		"entregabairro":                helpers.String(cpv.EntregaBairro),
		"shipping_address_postal_code": helpers.String(cpv.EntregaCep),
		"shipping_address_state":       helpers.String(cpv.EntregaUf),
		"transp":                       helpers.String(cpv.Transp),
		"created_by_id":                helpers.String(cpv.UsuarioInclusao),
		"modified_by_id":               helpers.String(cpv.UsuarioAlteracao),
		"created_at":                   helpers.StringDatetime(cpv.DataHoraInclusao),
		"modified_at":                  helpers.StringDatetime(cpv.DataHoraAlteracao),
		"status":                       helpers.String(cpv.Status),
		"tipoabordagem":                helpers.String(cpv.TipoAbordagem),
		"unidadenegocio":               helpers.String(cpv.UnidadeNegocio),
		"probabfech":                   helpers.String(cpv.ProbabFech),
		"meioconhec":                   helpers.String(cpv.MeioConhec),
		"emissao":                      helpers.StringDatetime(cpv.Emissao),
		"obssimples":                   helpers.String(cpv.ObsSimples),
		"codrepresentante":             helpers.String(cpv.CodRepresentante),
		"id":                           helpers.String(cpv.Numero),
	})

	// fmt.Println(queryString)

	//Iniciando uma transação
	ctx := context.Background()
	tx, _ := t.db.BeginTx(ctx, nil)

	_, err := tx.ExecContext(ctx, queryString)

	if err != nil {
		utils.LogDatabaseDetails("CPV", helpers.String(cpv.Numero), queryString, err.Error(), "")
		log.Println("Erro ao atualizar CPV: ", err)
		tx.Rollback()
		return err
	}

	log.Println("CPV atualizado com sucesso: ", cpv.Numero)

	commit := tx.Commit()

	if commit != nil {
		utils.LogDatabaseDetails("CFR", helpers.String(cpv.Numero), queryString, commit.Error(), "")
		return commit
	}

	return nil
}

func (t PedidoRepositoryDbErp) InsertCrm(cpv pedido.Cpv, owner string) error {

	queryString := utils.Msg(`INSERT into
									{{.owner}}.quote (id,
										number_a ,
										pedidofil,
										account_id,
										codigooperacao,
										filial,
										date_quoted,
										dataentrega,
										finalidade,
										pedcliente,
										condpagto,
										valormercadorias,
										valortotal,
										tipofrete ,
										pesoliquido,
										pesobruto,
										shipping_address_street,
										entregabairro,
										shipping_address_postal_code,
										shipping_address_state,
										transp,
										created_by_id,
										modified_by_id,
										created_at,
										modified_at,
										status,
										tipoabordagem,
										unidadenegocio,
										probabfech,
										meioconhec,
										emissao,
										obssimples,
										codrepresentante
									)VALUES (
									'{{.id}}',
									'{{.number_a}}',
									'{{.pedidofil}}',
									'{{.account_id}}',
									'{{.codigooperacao}}',
									'{{.filial}}',
									'{{.date_quoted}}',
									'{{.dataentrega}}',
									'{{.finalidade}}',
									'{{.pedcliente}}',
									'{{.condpagto}}',
									'{{.valormercadorias}}',
									'{{.valortotal}}',
									'{{.tipofrete}}',
									'{{.pesoliquido}}',
									'{{.pesobruto}}',
									'{{.shipping_address_street}}',
									'{{.entregabairro}}',
									'{{.shipping_address_postal_code}}',
									'{{.shipping_address_state}}',
									'{{.transp}}',
									'{{.created_by_id}}',
									'{{.modified_by_id}}',
									'{{.created_at}}',
									'{{.modified_at}}',
									'{{.status}}',
									'{{.tipoabordagem}}',
									'{{.unidadenegocio}}',
									'{{.probabfech}}',
									'{{.meioconhec}}',
									'{{.emissao}}',
									'{{.obssimples}}',
									'{{.codrepresentante}}'
									)`, map[string]interface{}{
		"owner":                        owner,
		"number_a":                     helpers.String(cpv.Numero),
		"pedidofil":                    helpers.String(cpv.PedidoFil),
		"account_id":                   helpers.String(cpv.Cliente),
		"codigooperacao":               helpers.String(cpv.CodigoOperacao),
		"filial":                       helpers.String(cpv.Filial),
		"date_quoted":                  helpers.StringDatetime(cpv.DataEntrada),
		"dataentrega":                  helpers.StringDatetime(cpv.DataEntrega),
		"finalidade":                   helpers.String(cpv.Finalidade),
		"pedcliente":                   helpers.String(cpv.PedCliente),
		"condpagto":                    helpers.String(cpv.CondPagamento),
		"valormercadorias":             helpers.String(cpv.ValorMercadorias),
		"valortotal":                   helpers.String(cpv.ValorTotal),
		"tipofrete":                    helpers.String(cpv.TipoFrete),
		"pesoliquido":                  helpers.String(cpv.PesoLiquido),
		"pesobruto":                    helpers.String(cpv.PesoBruto),
		"shipping_address_street":      helpers.String(cpv.EntregaEnd),
		"entregabairro":                helpers.String(cpv.EntregaBairro),
		"shipping_address_postal_code": helpers.String(cpv.EntregaCep),
		"shipping_address_state":       helpers.String(cpv.EntregaUf),
		"transp":                       helpers.String(cpv.Transp),
		"created_by_id":                helpers.String(cpv.UsuarioInclusao),
		"modified_by_id":               helpers.String(cpv.UsuarioAlteracao),
		"created_at":                   helpers.StringDatetime(cpv.DataHoraInclusao),
		"modified_at":                  helpers.StringDatetime(cpv.DataHoraAlteracao),
		"status":                       helpers.String(cpv.Status),
		"tipoabordagem":                helpers.String(cpv.TipoAbordagem),
		"unidadenegocio":               helpers.String(cpv.UnidadeNegocio),
		"probabfech":                   helpers.String(cpv.ProbabFech),
		"meioconhec":                   helpers.String(cpv.MeioConhec),
		"emissao":                      helpers.StringDatetime(cpv.Emissao),
		"obssimples":                   helpers.String(cpv.ObsSimples),
		"codrepresentante":             helpers.String(cpv.CodRepresentante),
		"id":                           helpers.String(cpv.Numero),
	})

	// fmt.Println(queryString)

	ctx := context.Background()
	tx, _ := t.db.BeginTx(ctx, nil)

	_, err := tx.ExecContext(ctx, queryString)

	if err != nil {
		utils.LogDatabaseDetails("CPV", helpers.String(cpv.Numero), queryString, err.Error(), "")
		tx.Rollback()
		return err
	}

	commit := tx.Commit()

	if commit != nil {
		utils.LogDatabaseDetails("CPV", helpers.String(cpv.Numero), queryString, commit.Error(), "")
		return commit
	}

	return nil
}

func (t PedidoRepositoryDbErp) DeleteErp(id string, tipo string) error {

	queryString := fmt.Sprintf(`DELETE FROM tb_crm_sincroniza WHERE pk = '%s' and tipo = '%s'`, id, tipo)

	fmt.Println(queryString)

	//Iniciando o contexto de transação
	ctx := context.Background()
	tx, _ := t.db.BeginTx(ctx, nil)

	_, err := tx.ExecContext(ctx, queryString)

	if err != nil {
		//utils.LogFile("CRM/CFR", " delete", "CRITICAL ", err.Error(), queryString)
		utils.LogDatabaseDetails("CPV", id, queryString, err.Error(), "")
		tx.Rollback()
		return err
	}

	//Para debug
	// tx.Rollback()

	commit := tx.Commit()

	if commit != nil {
		//utils.LogFile("CRM/CFR", " delete", "CRITICAL ", commit.Error(), queryString)
		utils.LogDatabaseDetails("CPV", id, queryString, commit.Error(), "")
		return err
	}

	return nil
}

//------------------------------------ ERP para CRM : IPV-------------------------------//

func (t PedidoRepositoryDbErp) SelectIpv() ([]pedido.Ipv, error) {

	var controlLimitQuery string = "LIMIT 10"

	queryString := utils.Msg(`SELECT
								tipo,
								pk,
								numero,
								item_pv,
								codigo_item,
								unidade,
								qtde_pedida,
								qtde_atendida,
								preco_unit,
								almoxarifado,
								descricao_informada,
								data_hora_inclusao,
								finalidade,
								data_entrega_item,
								usuario_inclusao
								from ipv
								inner join tb_crm_sincroniza ON numero||codigo_item||item_pv = pk and tabela = 'IPV'
								{{.controlLimitQuery}}
								`,
		map[string]interface{}{
			"controlLimitQuery": controlLimitQuery,
		})

	// fmt.Println(queryString)

	rows, err := t.db.Query(queryString)

	if err != nil {
		utils.LogDatabaseDetails("IPV", "SELECT", queryString, err.Error(), "")
		return nil, err
	}

	ipvs := []pedido.Ipv{}

	for rows.Next() {
		ipv := pedido.Ipv{}

		if err := rows.Scan(
			&ipv.Tipo,
			&ipv.Pk,
			&ipv.Numero,
			&ipv.ItemPv,
			&ipv.CodigoItem,
			&ipv.Unidade,
			&ipv.QtdePedida,
			&ipv.QtdeAtendida,
			&ipv.PrecoUnit,
			&ipv.Almoxarifado,
			&ipv.DescricaoInformada,
			&ipv.DataHoraInclusao,
			&ipv.Finalidade,
			&ipv.DataEntregaItem,
			&ipv.UsuarioInclusao,
		); err != nil {
			log.Println(err.Error())
			utils.LogDatabaseDetails("IPV", "SELECT", queryString, err.Error(), "")
			return nil, err
		}

		ipvs = append(ipvs, ipv)
	}

	return ipvs, nil
}

func (t PedidoRepositoryDbErp) CheckUpdateCrmIpv(id string, owner string) (bool, error) {

	queryString := fmt.Sprintf(`SELECT count(*) FROM %s.quote_item WHERE id = '%s' `, owner, id)

	fmt.Println(queryString)

	rows, err := t.db.Query(queryString)

	if err != nil {
		utils.LogDatabaseDetails("IPV", id, queryString, err.Error(), "")
		return false, err
	}

	var count int

	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			utils.LogDatabaseDetails("IPV", id, queryString, err.Error(), "")
			return false, err
		}
	}

	if count == 0 {
		return false, nil
	}

	return true, nil
}

//#TODO
func (t PedidoRepositoryDbErp) UpdateCrmIpv(ipv pedido.Ipv, owner string) error {
	return nil
}

//#TODO
func (t PedidoRepositoryDbErp) InsertCrmIpv(ipv pedido.Ipv, owner string) error {
	return nil
}

//#TODO
func (t PedidoRepositoryDbErp) DeleteCrmIpv(id string, owner string) error {
	return nil
}
