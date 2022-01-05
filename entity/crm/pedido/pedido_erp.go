package pedido

import "database/sql"

type Cpv struct {
	Tipo              sql.NullString
	Numero            sql.NullString
	PedidoFil         sql.NullString
	Cliente           sql.NullString
	CodigoOperacao    sql.NullString
	Filial            sql.NullString
	DataEntrada       sql.NullString
	DataEntrega       sql.NullString
	Finalidade        sql.NullString
	PedCliente        sql.NullString
	CondPagamento     sql.NullString
	ValorMercadorias  sql.NullString
	ValorTotal        sql.NullString
	TipoFrete         sql.NullString
	PesoLiquido       sql.NullString
	PesoBruto         sql.NullString
	EntregaEnd        sql.NullString
	EntregaBairro     sql.NullString
	EntregaCep        sql.NullString
	EntregaCidade     sql.NullString
	EntregaUf         sql.NullString
	Transp            sql.NullString
	UsuarioInclusao   sql.NullString
	UsuarioAlteracao  sql.NullString
	DataHoraInclusao  sql.NullString
	DataHoraAlteracao sql.NullString
	Status            sql.NullString
	TipoAbordagem     sql.NullString
	UnidadeNegocio    sql.NullString
	ProbabFech        sql.NullString
	MeioConhec        sql.NullString
	Emissao           sql.NullString
	ObsSimples        sql.NullString
	CodRepresentante  sql.NullString
}

type Ipv struct {
	Tipo               sql.NullString
	Pk                 sql.NullString
	Numero             sql.NullString
	ItemPv             sql.NullString
	CodigoItem         sql.NullString
	Unidade            sql.NullString
	QtdePedida         sql.NullString
	QtdeAtendida       sql.NullString
	PrecoUnit          sql.NullString
	Almoxarifado       sql.NullString
	DescricaoInformada sql.NullString
	DataHoraInclusao   sql.NullString
	Finalidade         sql.NullString
	DataEntregaItem    sql.NullString
	UsuarioInclusao    sql.NullString
}
