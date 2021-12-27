package cfr

import (
	"database/sql"
)

/*
	Fazendo as duas entidades Iguais, para sempre deixar organizado os retornos de cada etapa
*/
type Account struct {
	Tipo                    sql.NullString
	Id                      sql.NullString
	Name                    sql.NullString
	Description             sql.NullString
	NomeFantasia            sql.NullString
	NaturezaPessoa          sql.NullString
	SicCode                 sql.NullString
	InscricaoEstadual       sql.NullString
	SitFed                  sql.NullString
	SitCms                  sql.NullString
	SitIpi                  sql.NullString
	BillingAddressCity      sql.NullString
	NumeroEnd               sql.NullString
	EndComplemento          sql.NullString
	Bairro                  sql.NullString
	BillingAdressState      sql.NullString
	BillingAdressPostalCode sql.NullString
	BillingAdressCountry    sql.NullString
	Telefone1               sql.NullString
	Telefone2               sql.NullString
	Contato                 sql.NullString
	Email                   sql.NullString
	DataNascimento          sql.NullString
	EmailNfe                sql.NullString
	WebSite                 sql.NullString
	NossoCliente            sql.NullString
	NossoFornecedor         sql.NullString
	NossoTransportador      sql.NullString
	Industry                sql.NullString
	CondPgPadraoClie        sql.NullString
	DocPadraoCliente        sql.NullString
	CobrPadraoCliente       sql.NullString
	PortPadraoCliente       sql.NullString
	CreatedAt               sql.NullString
	ZonaFranca              sql.NullString
	InscricaoSuframa        sql.NullString
	Vendedor1               sql.NullString
	Comissao1Fat            sql.NullFloat64
	TranspPadraoCliente     sql.NullString
	TipoFrete               sql.NullString
	Endereco                sql.NullString
	EnderecoCobranca        sql.NullString
	NumeroEndCobranca       sql.NullString
	BairroCobranca          sql.NullString
	CidadeCobranca          sql.NullString //subquery
	UfCobranca              sql.NullString //subquery
	Cep                     sql.NullString
	CepCobranca             sql.NullString
	CreatedById             sql.NullString
	ModifiedAt              sql.NullString
	ModifiedById            sql.NullString
	EndeLatitude            sql.NullFloat64
	EndeLongitude           sql.NullFloat64
	Ti9Codigo               sql.NullString
	CategoriaCliente        sql.NullString
	OptSimples              sql.NullString
	ContribIcms             sql.NullString
	ConsumidorFinal         sql.NullString
	CkVerVenctoLote         sql.NullString
	QtdeMinVenctoLote       sql.NullString
	InssRet                 sql.NullString
	IsentoIcms              sql.NullString
	RetPiscoFcsvll          sql.NullString
	RetIss                  sql.NullString
	RetIssFonte             sql.NullString
	SubstTributIcms         sql.NullString
	SubstTributPis          sql.NullString
	SubstTributConfis       sql.NullString
	Status                  sql.NullString
	Origem                  sql.NullString
	ContaContabil           sql.NullString
	ContaSintetica          sql.NullString
	Cidade                  sql.NullString
	Uf                      sql.NullString
	TelefoneCobranca        sql.NullString
}

type Cfr struct {
	Tipo                    sql.NullString
	Id                      sql.NullString
	Name                    sql.NullString
	Description             sql.NullString
	NomeFantasia            sql.NullString
	NaturezaPessoa          sql.NullString
	SicCode                 sql.NullString
	InscricaoEstadual       sql.NullString
	SitFed                  sql.NullString
	SitCms                  sql.NullString
	SitIpi                  sql.NullString
	BillingAddressCity      sql.NullString
	NumeroEnd               sql.NullString
	EndComplemento          sql.NullString
	Bairro                  sql.NullString
	BillingAdressState      sql.NullString
	BillingAdressPostalCode sql.NullString
	BillingAdressCountry    sql.NullString
	Telefone1               sql.NullString
	Telefone2               sql.NullString
	Contato                 sql.NullString
	Email                   sql.NullString
	DataNascimento          sql.NullString
	EmailNfe                sql.NullString
	WebSite                 sql.NullString
	NossoCliente            sql.NullString
	NossoFornecedor         sql.NullString
	NossoTransportador      sql.NullString
	Industry                sql.NullString
	CondPgPadraoClie        sql.NullString
	DocPadraoCliente        sql.NullString
	CobrPadraoCliente       sql.NullString
	PortPadraoCliente       sql.NullString
	CreatedAt               sql.NullString
	ZonaFranca              sql.NullString
	InscricaoSuframa        sql.NullString
	Vendedor1               sql.NullString
	Comissao1Fat            sql.NullFloat64
	TranspPadraoCliente     sql.NullString
	TipoFrete               sql.NullString
	Endereco                sql.NullString
	EnderecoCobranca        sql.NullString
	NumeroEndCobranca       sql.NullString
	BairroCobranca          sql.NullString
	CidadeCobranca          sql.NullString //subquery
	UfCobranca              sql.NullString //subquery
	Cep                     sql.NullString
	CepCobranca             sql.NullString
	CreatedById             sql.NullString
	ModifiedAt              sql.NullString
	ModifiedById            sql.NullString
	EndeLatitude            sql.NullFloat64
	EndeLongitude           sql.NullFloat64
	Ti9Codigo               sql.NullString
	CategoriaCliente        sql.NullString
	OptSimples              sql.NullString
	ContribIcms             sql.NullString
	ConsumidorFinal         sql.NullString
	CkVerVenctoLote         sql.NullString
	QtdeMinVenctoLote       sql.NullString
	InssRet                 sql.NullString
	IsentoIcms              sql.NullString
	RetPiscoFcsvll          sql.NullString
	RetIss                  sql.NullString
	RetIssFonte             sql.NullString
	SubstTributIcms         sql.NullString
	SubstTributPis          sql.NullString
	SubstTributConfis       sql.NullString
	Status                  sql.NullString
	Origem                  sql.NullString
	ContaContabil           sql.NullString
	ContaSintetica          sql.NullString
	Cidade                  sql.NullString
	Uf                      sql.NullString
	TelefoneCobranca        sql.NullString
}
