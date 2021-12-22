package pedido

import "time"

type Quote struct {
	Id               string
	Filial           string
	Account_id       string
	ti9Codigo        string
	CodigoOperacao   string
	Finalidade       string
	Data_quoted      time.Time
	Moeda            int
	DataEntrada      time.Time
	DataEntrega      time.Time
	CondPagamento    string
	Indpres          string
	OrigemRegistro   string
	ObsSimples       string
	ObsFiscal        string
	OutrasInfoCom    string
	ObsPalm          string
	Amount           string
	DescFinanc       string
	TipoFrete        string
	Transp           string
	CodRepresentante string
	id               string
}

type QuoteItem struct {
	ItemPv             string
	ProductId          string
	Almoxarifado       string
	DataEntregaItem    string
	Finalidade         string
	QtdePedida         string
	NumeroPc           string
	ItemPc             string
	CodigoCentro       string
	Quantity           string
	UnitPrice          string
	Discount           string
	QuoteId            string
	Data_Entrega_Item  time.Time //Conferir aqui
	DataEntrega        time.Time
	Amount             string
	Order              string
	DescricaoInformada string
	FatorConv2         string
	UsuarioInclusao    string
	DescValor          float64
}
