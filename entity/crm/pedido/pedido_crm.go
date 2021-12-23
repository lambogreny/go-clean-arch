package pedido

import (
	"database/sql"
)

type Quote struct {
	Id               string
	Filial           sql.NullString
	Account_id       sql.NullString
	Ti9Codigo        sql.NullString
	CodigoOperacao   sql.NullString
	Finalidade       sql.NullString
	Data_quoted      sql.NullString
	Moeda            sql.NullInt64
	DataEntrada      sql.NullString //Tratando data como string mesmo
	DataEntrega      sql.NullString //Tratando data como string mesmo
	CondPagamento    sql.NullString
	Indpres          sql.NullString
	OrigemRegistro   sql.NullString
	ObsSimples       sql.NullString
	ObsFiscal        sql.NullString
	OutrasInfoCom    sql.NullString
	ObsPalm          sql.NullString
	Amount           sql.NullString
	DescFinanc       sql.NullString
	CodCobranca      sql.NullInt64
	PedCliente       sql.NullString
	TipoFrete        sql.NullString
	Transp           sql.NullString
	CodRepresentante sql.NullString
	id               sql.NullString
}

type QuoteItem struct {
	ItemPv             sql.NullString
	ProductId          sql.NullString
	Almoxarifado       sql.NullString
	DataEntregaItem    sql.NullString
	Finalidade         sql.NullString
	QtdePedida         sql.NullString
	NumeroPc           sql.NullString
	ItemPc             sql.NullString
	CodigoCentro       sql.NullString
	Quantity           sql.NullString
	UnitPrice          sql.NullString
	Discount           sql.NullString
	QuoteId            sql.NullString
	Data_Entrega_Item  sql.NullTime
	DataEntrega        sql.NullTime
	Amount             sql.NullString
	Order              sql.NullString
	DescricaoInformada sql.NullString
	FatorConv2         sql.NullString
	UsuarioInclusao    sql.NullString
	DescValor          sql.NullInt64
}

type ItemsPayload struct {
	ItemPv            string `json:"item_pv"`
	ProductId         string `json:"codigo_item"`
	Almoxarifado      string `json:"almoxarifado"`
	DataEntregaItem   string `json:"data_entrega_item"`
	Finalidade        string `json:"finalidade"`
	QtdePedida        string `json:"qtde_pedida"`
	CodigoCentro      string `json:"codigo_centro"`
	UnitPrice         string `json:"preco_unit"`
	Discount          string `json:"desc_perc"`
	DescValor         string `json:"desc_valor"`
	Data_Entrega_Item int64  `json:"data_entrega_item"`
}

type PedidoPayload struct {
	Filial             string         `json:"filial"`
	Cliente            string         `json:"cliente"`
	CodOperacao        string         `json:"codigo_operacao"`
	DataEntrada        string         `json:"data_entrada"`
	DataEntrega        string         `json:"data_entrega"`
	CondPagamento      string         `json:"cond_pagto"`
	IndPres            string         `json:"ind_pres"`
	OrigemRegistro     string         `json:"origem_registro"`
	ObsSimples         string         `json:"obs_simples"`
	ObsFiscal          string         `json:"obs_fiscal"`
	OutrasInfoCom      string         `json:"outras_inf_com"`
	Moeda              int64          `json:"tipo_moeda"`
	TipoFrete          string         `json:"tipo_frete"`
	Transp             string         `json:"transp"`
	DescFinanc         string         `json:"desc_financ"`
	CodUsuarioInclusao string         `json:"cod_usuario_inclusao"` //Ver um chapado
	Finalidade         string         `json:"finalidade"`
	NumeroPedPalm      string         `json:"numeropedpalm"` //id
	Items              []ItemsPayload `json:"items"`
}
