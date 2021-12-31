package crmRepository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/augusto/imersao5-esquenta-go/entity/crm/cfr"
	"github.com/augusto/imersao5-esquenta-go/utils"
	"github.com/augusto/imersao5-esquenta-go/utils/helpers"
)

type CfrRepositoryDbErp struct {
	db *sql.DB
}

func NewCfrRepositoryDbErp(db *sql.DB) *CfrRepositoryDbErp {
	return &CfrRepositoryDbErp{db: db}
}

func (t CfrRepositoryDbErp) CheckUpdateErp(id string) (bool, error) {
	queryString := fmt.Sprintf(`SELECT count(*) from cfr  where codigo_pessoa= '%s'`, id)

	rows, err := t.db.Query(queryString)

	if err != nil {
		utils.LogDatabaseDetails("CFR", id, queryString, err.Error(), "")
		return false, err
	}

	var count int

	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			utils.LogFile("CRM/CFR", " CheckUpdateErp", "CRITICAL ", err.Error(), queryString)
			return false, err
		}
	}

	if count == 0 {
		return false, nil
	}
	return true, nil
}

func (t CfrRepositoryDbErp) SelectCrm(owner string) ([]cfr.Account, error) {

	queryString := utils.Msg(`SELECT
                                    tipo,
                                    id,
                                    UPPER(name),
                                    UPPER(description),
                                    UPPER(nomefantasia),
                                    UPPER(naturezapessoa),
                                    UPPER(sic_code),
                                    UPPER(inscricaoestadual),
                                    'COM' AS sitfed,
                                    'CI' AS siticms,
                                    '1' AS sitipi,
                                    UPPER(billing_address_street),
                                    UPPER(numeroend),
                                    UPPER(endcomplemento),
                                    UPPER(bairro),
                                    UPPER(billing_address_state),
                                    UPPER(billing_address_postal_code),
                                    'BRA' AS billing_address_country,
                                    telefone1,
                                    telefone2,
                                    UPPER(contato),
                                    UPPER(email),
                                    '' AS data_nascimento,
                                    UPPER(emailnfe),
                                    UPPER(website),
                                    't' AS nosso_cliente,
                                    'f' AS nosso_fornecedor,
                                    'f' AS nosso_transportador,
                                    UPPER(industry),
                                    '' AS cond_pg_padrao_clie,
                                    '' AS docpadrao_cliente,
                                    '' AS cobrpadrao_cliente,
                                    '' AS portpadrao_cliente,
                                    current_timestamp AS created_at,
                                    'f' AS zonafranca,
                                    '' AS inscricao_suframa,
                                    '' AS vendedor1,
                                    0.0 AS comissao1_fat,
                                    '' AS transp_padr_clie,
                                    '' AS tipo_frete,
                                    endereco,
                                    enderecocobranca AS endereco_cobranca,
                                    numeroendcobranca as numeroend_cobranca,
                                    bairroCobranca as bairro_cobranca,
                                    (SELECT codigomunicipio FROM cadastro_municpios WHERE cadastro_municpios.id = account.cadastro_municpios12_id) AS cidade_cobranca,
                                    (SELECT uf FROM cadastro_municpios WHERE cadastro_municpios.id = account.cadastro_municpios12_id) AS uf_cobranca,
                                    cep,
                                    cepCobranca as cep_cobranca,
                                    'CRM' AS created_by_id,
                                    current_timestamp AS modified_at,
                                    'CRM' AS modified_by_id,
                                    0.0 AS ende_latitude,
                                    0.0 AS ende_longitude,
                                    TI9CODIGO,
                                    '001' AS categoria_cliente,
                                    'f' AS opt_simples,
                                    't' AS contrib_icms,
                                    'f' AS consumidor_final,
                                    'f' AS ck_ver_vencto_lote,
                                    0 AS qtde_min_vencto_lote,
                                    'f' AS inss_ret,
                                    'f' AS isento_icms,
                                    'f' AS ret_piscofcsll,
                                    'f' AS ret_iss,
                                    'f' AS ret_iss_fonte,
                                    'f' AS subst_tribut_icms,
                                    'f' AS subst_tribut_pis,
                                    'f' AS subst_tribut_cofins,
                                    status,
                                    origem,
                                    conta_contabil,
                                    conta_sintetica,
                                    (SELECT codigomunicipio FROM cadastro_municpios WHERE cadastro_municpios.id = account.cadastro_municpios_id) AS cidade,
									(SELECT uf FROM cadastro_municpios WHERE cadastro_municpios.id = account.cadastro_municpios_id) AS uf,
									telefonecobranca
                                    from
                                    {{.owner}}.account
                                    inner join {{.owner}}.tb_crm_sincroniza ON id = pk
                                    AND tabela = 'CFR'
                                    WHERE
                                    IFNULL(usuario, 'xx') <> 'CRM'
`, map[string]interface{}{
		"owner": owner,
	})

	rows, err := t.db.Query(queryString)

	if err != nil {
		utils.LogDatabaseDetails("CFR", "SELECT", queryString, err.Error(), "")
		return nil, err
	}

	accounts := []cfr.Account{}

	for rows.Next() {
		account := cfr.Account{}

		if err := rows.Scan(&account.Tipo,
			&account.Id,
			&account.Name,
			&account.Description,
			&account.NomeFantasia,
			&account.NaturezaPessoa,
			&account.SicCode,
			&account.InscricaoEstadual,
			&account.SitFed,
			&account.SitCms,
			&account.SitIpi,
			&account.BillingAddressCity,
			&account.NumeroEnd,
			&account.EndComplemento,
			&account.Bairro,
			&account.BillingAdressState,
			&account.BillingAdressPostalCode,
			&account.BillingAdressCountry,
			&account.Telefone1,
			&account.Telefone2,
			&account.Contato,
			&account.Email,
			&account.DataNascimento,
			&account.EmailNfe,
			&account.WebSite,
			&account.NossoCliente,
			&account.NossoFornecedor,
			&account.NossoTransportador,
			&account.Industry,
			&account.CondPgPadraoClie,
			&account.DocPadraoCliente,
			&account.CobrPadraoCliente,
			&account.PortPadraoCliente,
			&account.CreatedAt,
			&account.ZonaFranca,
			&account.InscricaoSuframa,
			&account.Vendedor1,
			&account.Comissao1Fat,
			&account.TranspPadraoCliente,
			&account.TipoFrete,
			&account.Endereco,
			&account.EnderecoCobranca,
			&account.NumeroEndCobranca,
			&account.BairroCobranca,
			&account.CidadeCobranca,
			&account.UfCobranca,
			&account.Cep,
			&account.CepCobranca,
			&account.CreatedById,
			&account.ModifiedAt,
			&account.ModifiedById,
			&account.EndeLatitude,
			&account.EndeLongitude,
			&account.Ti9Codigo,
			&account.CategoriaCliente,
			&account.OptSimples,
			&account.ContribIcms,
			&account.ConsumidorFinal,
			&account.CkVerVenctoLote,
			&account.QtdeMinVenctoLote,
			&account.InssRet,
			&account.IsentoIcms,
			&account.RetPiscoFcsvll,
			&account.RetIss,
			&account.RetIssFonte,
			&account.SubstTributIcms,
			&account.SubstTributPis,
			&account.SubstTributConfis,
			&account.Status,
			&account.Origem,
			&account.ContaContabil,
			&account.ContaSintetica,
			&account.Cidade,
			&account.Uf,
			&account.TelefoneCobranca,
		); err != nil {
			log.Println(err.Error())
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (t CfrRepositoryDbErp) UpdateErp(account cfr.Account, owner string) error {

	queryString := utils.Msg(`UPDATE cfr SET
                                        nome_pessoa = '{{.nome_pessoa}}',
                                        nome_fantasia = '{{.nome_fantasia}}',
                                        natureza_pessoa = '{{.natureza_pessoa}}',
                                        cnpj_cpf = '{{.cnpj_cpf}}',
                                        inscricao_estadual= '{{.inscricao_estadual}}',
                                        sitfed= '{{.sitfed}}',
                                        siticms = '{{.siticms}}',
                                        sitipi = '{{.sitipi}}',
                                        endereco = '{{.endereco}}',
                                        numeroend = '{{.numeroend}}',
                                        end_complemento= '{{.end_complemento}}',
                                        bairro = '{{.bairro}}',
                                        cidade = '{{.cidade}}',
                                        uf = '{{.uf}}',
                                        cep = '{{.cep}}',
                                        pais = '{{.pais}}',
                                        telefone_1 = '{{.telefone_1}}',
                                        telefone_2= '{{.telefone_2}}',
                                        contato = '{{.contato}}',
                                        email = '{{.email}}',
                                        emailnfe = '{{.emailnfe}}',
                                        home_page = '{{.home_page}}',
                                        nosso_cliente = '{{.nosso_cliente}}',
                                        cond_pg_padrao_clie = '{{.cond_pg_padrao_clie}}',
                                        docpadrao_cliente = '{{.docpadrao_cliente}}',
                                        cobrpadrao_cliente = '{{.cobrpadrao_cliente}}',
                                        portpadrao_cliente = '{{.portpadrao_cliente}}',
                                        status = '{{.status}}',
                                        data_cad = '{{.data_cad}}',
                                        zonafranca = '{{.zonafranca}}',
                                        inscricao_suframa = '{{.inscricao_suframa}}',
                                        vendedor1 = '{{.vendedor1}}',
                                        comissao1_fat = '{{.comissao1_fat}}',
                                        transp_padr_clie = '{{.transp_padr_clie}}',
                                        tipo_frete = '{{.tipo_frete}}',
                                        endereco_cobranca = '{{.endereco_cobranca}}',
                                        bairro_cobranca= '{{.bairro_cobranca}}',
                                        cidade_cobranca = '{{.cidade_cobranca}}',
                                        uf_cobranca = '{{.uf_cobranca}}',
                                        cep_cobranca = '{{.cep_cobranca}}',
                                        data_hora_inclusao = '{{.data_hora_inclusao}}',
                                        usuario_inclusao = '{{.usuario_inclusao}}',
                                        data_hora_alteracao = '{{.data_hora_alteracao}}',
                                        usuario_alteracao = '{{.usuario_alteracao}}',
                                        ende_latitude = '{{.ende_latitude}}',
                                        ende_longitude = '{{.ende_longitude}}',
                                        categoria_cliente = '{{.categoria_cliente}}',
                                        opt_simples = '{{.opt_simples}}',
                                        contrib_icms = '{{.contrib_icms}}',
                                        consumidor_final = '{{.consumidor_final}}',
                                        ck_ver_vencto_lote = '{{.ck_ver_vencto_lote}}',
                                        qtde_min_vencto_lote = '{{.qtde_min_vencto_lote}}',
                                        inss_ret = '{{.inss_ret}}',
                                        isento_icms = '{{.isento_icms}}',
                                        ret_piscofcsll = '{{.ret_piscofcsll}}',
                                        ret_iss = '{{.ret_iss}}',
                                        ret_iss_fonte = '{{.ret_iss_fonte}}',
                                        subst_tribut_icms = '{{.subst_tribut_icms}}',
                                        subst_tribut_pis = '{{.subst_tribut_pis}}',
                                        subst_tribut_cofins = '{{.subst_tribut_cofins}}',
                                        origem = '{{.origem}}',
                                        conta_contabil_clie_1 = '{{.conta_contabil_clie_1}}',
                                        conta_contabil_s_clie_1 = '{{.conta_contabil_s_clie_1}}'
									WHERE codigo_pessoa = '{{.codigo_pessoa}}' `, map[string]interface{}{
		"nome_pessoa":             helpers.String(account.Name),
		"nome_fantasia":           helpers.String(account.NomeFantasia),
		"natureza_pessoa":         helpers.String(account.NaturezaPessoa),
		"cnpj_cpf":                helpers.String(account.SicCode),
		"inscricao_estadual":      helpers.String(account.InscricaoEstadual),
		"sitfed":                  helpers.String(account.SitFed),
		"siticms":                 helpers.String(account.SitCms),
		"sitipi":                  helpers.String(account.SitIpi),
		"endereco":                helpers.String(account.Endereco),
		"numeroend":               helpers.String(account.NumeroEnd),
		"end_complemento":         helpers.String(account.EndComplemento),
		"bairro":                  helpers.String(account.Bairro),
		"cidade":                  helpers.String(account.Cidade),
		"uf":                      helpers.String(account.Uf),
		"cep":                     helpers.String(account.Cep),
		"pais":                    helpers.String(account.BillingAdressCountry),
		"telefone_1":              helpers.String(account.Telefone1),
		"telefone_2":              helpers.String(account.Telefone2),
		"contato":                 helpers.String(account.Contato),
		"email":                   helpers.String(account.Email),
		"emailnfe":                helpers.String(account.EmailNfe),
		"home_page":               helpers.String(account.WebSite),
		"nosso_cliente":           helpers.StringBoolean(account.NossoCliente),
		"cond_pg_padrao_clie":     helpers.String(account.CondPgPadraoClie),
		"docpadrao_cliente":       helpers.String(account.DocPadraoCliente),
		"cobrpadrao_cliente":      helpers.String(account.CobrPadraoCliente),
		"portpadrao_cliente":      helpers.String(account.PortPadraoCliente),
		"status":                  helpers.String(account.Status),
		"data_cad":                helpers.String(account.CreatedAt),
		"zonafranca":              helpers.StringBoolean(account.ZonaFranca),
		"inscricao_suframa":       helpers.String(account.InscricaoSuframa),
		"vendedor1":               helpers.String(account.Vendedor1),
		"comissao1_fat":           helpers.Float64(account.Comissao1Fat),
		"transp_padr_clie":        helpers.String(account.TranspPadraoCliente),
		"tipo_frete":              helpers.String(account.TipoFrete),
		"endereco_cobranca":       helpers.String(account.EnderecoCobranca),
		"bairro_cobranca":         helpers.String(account.BairroCobranca),
		"cidade_cobranca":         helpers.String(account.CidadeCobranca),
		"uf_cobranca":             helpers.String(account.UfCobranca),
		"cep_cobranca":            helpers.String(account.CepCobranca),
		"data_hora_inclusao":      helpers.String(account.ModifiedAt),
		"usuario_inclusao":        helpers.String(account.CreatedById),
		"data_hora_alteracao":     helpers.String(account.ModifiedAt),
		"usuario_alteracao":       helpers.String(account.ModifiedById),
		"ende_latitude":           helpers.Float64(account.EndeLatitude),
		"ende_longitude":          helpers.Float64(account.EndeLongitude),
		"categoria_cliente":       helpers.String(account.CategoriaCliente),
		"opt_simples":             helpers.StringBoolean(account.OptSimples),
		"contrib_icms":            helpers.StringBoolean(account.ContribIcms),
		"consumidor_final":        helpers.StringBoolean(account.ConsumidorFinal),
		"ck_ver_vencto_lote":      helpers.StringBoolean(account.CkVerVenctoLote),
		"qtde_min_vencto_lote":    helpers.String(account.QtdeMinVenctoLote),
		"inss_ret":                helpers.StringBoolean(account.InssRet),
		"isento_icms":             helpers.StringBoolean(account.IsentoIcms),
		"ret_piscofcsll":          helpers.StringBoolean(account.RetPiscoFcsvll),
		"ret_iss":                 helpers.StringBoolean(account.RetIss),
		"ret_iss_fonte":           helpers.StringBoolean(account.RetIssFonte),
		"subst_tribut_icms":       helpers.StringBoolean(account.SubstTributIcms),
		"subst_tribut_pis":        helpers.StringBoolean(account.SubstTributPis),
		"subst_tribut_cofins":     helpers.StringBoolean(account.SubstTributConfis),
		"origem":                  helpers.String(account.Origem),
		"conta_contabil_clie_1":   helpers.String(account.ContaContabil),
		"conta_contabil_s_clie_1": helpers.String(account.ContaSintetica),
		"codigo_pessoa":           helpers.String(account.Ti9Codigo),
	})

	//Iniciando uma transação
	ctx := context.Background()
	tx, _ := t.db.BeginTx(ctx, nil)

	_, err := tx.ExecContext(ctx, queryString)

	//Debug
	//utils.LogFile("CRM/DEBUG", " updateCfr", "DEBUG ", helpers.String(account.Id), queryString)

	if err != nil {
		//utils.LogFile("CRM/CFR", " update", "CRITICAL ", err.Error(), queryString)
		utils.LogDatabaseDetails("CFR", helpers.String(account.Id), queryString, err.Error(), "")
		tx.Rollback()
		return err
	}

	commit := tx.Commit()

	if commit != nil {
		//utils.LogFile("CRM/CFR", " update", "CRITICAL ", err.Error(), "erro no commmit")
		utils.LogDatabaseDetails("CFR", helpers.String(account.Id), queryString, commit.Error(), "")
		return commit
	}

	return nil

}

func (t CfrRepositoryDbErp) DeleteCrm(owner string, id string, tipo string) error {
	queryString := utils.Msg(`DELETE from {{.owner}}.tb_crm_sincroniza WHERE tabela = 'CFR' AND pk = '{{.pk}}' and tipo = '{{.tipo}}'`, map[string]interface{}{
		"owner": owner,
		"pk":    id,
		"tipo":  tipo,
	})
	fmt.Println(queryString)

	//Iniciando o contexto de transação
	ctx := context.Background()
	tx, _ := t.db.BeginTx(ctx, nil)

	_, err := tx.ExecContext(ctx, queryString)

	if err != nil {
		//utils.LogFile("CRM/CFR", " delete", "CRITICAL ", err.Error(), queryString)
		utils.LogDatabaseDetails("CFR", id, queryString, err.Error(), "")
		tx.Rollback()
		return err
	}

	commit := tx.Commit()

	if commit != nil {
		//utils.LogFile("CRM/PRD", " delete", "CRITICAL ", commit.Error(), queryString)
		utils.LogDatabaseDetails("CFR", id, queryString, err.Error(), "")
		return err
	}

	return nil
}

func (t CfrRepositoryDbErp) InsertErp(account cfr.Account, owner string) error {

	//Iniciando uma transação
	ctx := context.Background()
	tx, _ := t.db.BeginTx(ctx, nil)

	// ----------------------------------------------------------- Selecionando os valores da sys sequencial ------------------------------------------------------------//

	var v_valor int
	var valor_final int

	sqlGetSysSequencial := `SELECT
                                        valor AS v_valor,
                                        LPAD((valor::integer + 1) :: text, 6, '0') AS valor_final
                                    FROM
                                        sys_sequencial
                                    WHERE
                                        tabela = 'cfr'`
	err := tx.QueryRowContext(ctx, sqlGetSysSequencial).Scan(&v_valor, &valor_final)

	switch {
	case err == sql.ErrNoRows:
		fmt.Errorf("Valores da sys sequencial não encontrados!")
		utils.LogDatabaseDetails("ACCOUNT", helpers.String(account.Id), sqlGetSysSequencial, err.Error(), "")
		return err
	case err != nil:
		utils.LogDatabaseDetails("ACCOUNT", helpers.String(account.Id), sqlGetSysSequencial, err.Error(), "")
		return err
	}

	// ----------------------------------------------------------- Realizando update da sys sequencial ------------------------------------------------------------//

	sqlUpdateSysSequencial := fmt.Sprintf(`UPDATE
                                            sys_sequencial
                                        set
                                            valor_anterior = %d,
                                            valor = %d
                                        where
                                            tabela = 'cfr'`, v_valor, valor_final)
	_, updateErr := tx.ExecContext(ctx, sqlUpdateSysSequencial)

	if updateErr != nil {
		//utils.LogFile("CRM/CFR", " insert", "CRITICAL ", updateErr.Error(), sqlUpdateSysSequencial)
		utils.LogDatabaseDetails("ACCOUNT", helpers.String(account.Id), sqlUpdateSysSequencial, updateErr.Error(), "")
		tx.Rollback()
		return updateErr
	}

	// ----------------------------------------------------------- Realizando o insert na CFR ------------------------------------------------------------//

	sqlInsert := utils.Msg(`INSERT INTO cfr (
                                        nome_pessoa ,
                                        nome_fantasia ,
                                        natureza_pessoa ,
                                        cnpj_cpf ,
                                        inscricao_estadual ,
                                        sitfed ,
                                        siticms ,
                                        sitipi ,
                                        endereco ,
                                        numeroend ,
                                        end_complemento ,
                                        bairro ,
                                        cidade ,
                                        uf ,
                                        cep ,
                                        pais ,
                                        telefone_1 ,
                                        telefone_2 ,
                                        contato ,
                                        email ,
                                        emailnfe ,
                                        home_page ,
                                        nosso_cliente ,
                                        cond_pg_padrao_clie ,
                                        docpadrao_cliente ,
                                        cobrpadrao_cliente ,
                                        portpadrao_cliente ,
                                        status ,
                                        data_cad ,
                                        zonafranca ,
                                        inscricao_suframa ,
                                        vendedor1 ,
                                        comissao1_fat ,
                                        transp_padr_clie ,
                                        tipo_frete ,
                                        endereco_cobranca ,
                                        bairro_cobranca ,
                                        cidade_cobranca ,
                                        uf_cobranca ,
                                        cep_cobranca ,
                                        data_hora_inclusao ,
                                        usuario_inclusao ,
                                        data_hora_alteracao ,
                                        usuario_alteracao ,
                                        ende_latitude ,
                                        ende_longitude ,
                                        categoria_cliente ,
                                        opt_simples ,
                                        contrib_icms ,
                                        consumidor_final ,
                                        ck_ver_vencto_lote ,
                                        qtde_min_vencto_lote ,
                                        inss_ret ,
                                        isento_icms ,
                                        ret_piscofcsll ,
                                        ret_iss ,
                                        ret_iss_fonte ,
                                        subst_tribut_icms ,
                                        subst_tribut_pis ,
                                        subst_tribut_cofins ,
                                        origem ,
                                        conta_contabil_clie_1 ,
                                        conta_contabil_s_clie_1 ,
										telefone_cobranca
										)
										VALUES(
										'{{.nome_pessoa}}',
										'{{.nome_fantasia}}',
										'{{.natureza_pessoa}}',
										'{{.cnpj_cpf}}',
										'{{.inscricao_estadual}}',
										'{{.sitfed}}',
										'{{.siticms}}',
										'{{.sitipi}}',
										'{{.endereco}}',
										'{{.numeroend}}',
										'{{.end_complemento}}',
										'{{.bairro}}',
										'{{.cidade}}',
										'{{.uf}}',
										'{{.cep}}',
										'{{.pais}}',
										'{{.telefone_1}}',
										'{{.telefone_2}}',
										'{{.contato}}',
										'{{.email}}',
										'{{.emailnfe}}',
										'{{.home_page}}',
										'{{.nosso_cliente}}',
										'{{.cond_pg_padrao_clie}}',
										'{{.docpadrao_cliente}}',
										'{{.cobrpadrao_cliente}}',
										'{{.portpadrao_cliente}}',
										'{{.status}}',
										'{{.data_cad}}',
										'{{.zonafranca}}',
										'{{.inscricao_suframa}}',
										'{{.vendedor1}}',
										'{{.comissao1_fat}}',
										'{{.transp_padr_clie}}',
										'{{.tipo_frete}}',
										'{{.endereco_cobranca}}',
										'{{.bairro_cobranca}}',
										'{{.cidade_cobranca}}',
										'{{.uf_cobranca}}',
										'{{.cep_cobranca}}',
										'{{.data_hora_inclusao}}',
										'{{.usuario_inclusao}}',
										'{{.data_hora_alteracao}}',
										'{{.usuario_alteracao}}',
										'{{.ende_latitude}}',
										'{{.ende_longitude}}',
										'{{.categoria_cliente}}',
										'{{.opt_simples}}',
										'{{.contrib_icms}}',
										'{{.consumidor_final}}',
										'{{.ck_ver_vencto_lote}}',
										'{{.qtde_min_vencto_lote}}',
										'{{.inss_ret}}',
										'{{.isento_icms}}',
										'{{.ret_piscofcsll}}',
										'{{.ret_iss}}',
										'{{.ret_iss_fonte}}',
										'{{.subst_tribut_icms}}',
										'{{.subst_tribut_pis}}',
										'{{.subst_tribut_cofins}}',
										'{{.origem}}',
										'{{.conta_contabil_clie_1}}',
										'{{.conta_contabil_s_clie_1}}',
										'{{.telefone_cobranca}}'
										)
									`, map[string]interface{}{
		"nome_pessoa":             helpers.String(account.Name),
		"nome_fantasia":           helpers.String(account.NomeFantasia),
		"natureza_pessoa":         helpers.String(account.NaturezaPessoa),
		"cnpj_cpf":                helpers.String(account.SicCode),
		"inscricao_estadual":      helpers.String(account.InscricaoEstadual),
		"sitfed":                  helpers.String(account.SitFed),
		"siticms":                 helpers.String(account.SitCms),
		"sitipi":                  helpers.String(account.SitIpi),
		"endereco":                helpers.String(account.Endereco),
		"numeroend":               helpers.String(account.NumeroEnd),
		"end_complemento":         helpers.String(account.EndComplemento),
		"bairro":                  helpers.String(account.Bairro),
		"cidade":                  helpers.String(account.Cidade),
		"uf":                      helpers.String(account.Uf),
		"cep":                     helpers.String(account.Cep),
		"pais":                    helpers.String(account.BillingAdressCountry),
		"telefone_1":              helpers.String(account.Telefone1),
		"telefone_2":              helpers.String(account.Telefone2),
		"contato":                 helpers.String(account.Contato),
		"email":                   helpers.String(account.Email),
		"emailnfe":                helpers.String(account.EmailNfe),
		"home_page":               helpers.String(account.WebSite),
		"nosso_cliente":           helpers.StringBoolean(account.NossoCliente),
		"cond_pg_padrao_clie":     helpers.String(account.CondPgPadraoClie),
		"docpadrao_cliente":       helpers.String(account.DocPadraoCliente),
		"cobrpadrao_cliente":      helpers.String(account.CobrPadraoCliente),
		"portpadrao_cliente":      helpers.String(account.PortPadraoCliente),
		"status":                  helpers.String(account.Status),
		"data_cad":                helpers.String(account.CreatedAt),
		"zonafranca":              helpers.StringBoolean(account.ZonaFranca),
		"inscricao_suframa":       helpers.String(account.InscricaoSuframa),
		"vendedor1":               helpers.String(account.Vendedor1),
		"comissao1_fat":           helpers.Float64(account.Comissao1Fat),
		"transp_padr_clie":        helpers.String(account.TranspPadraoCliente),
		"tipo_frete":              helpers.String(account.TipoFrete),
		"endereco_cobranca":       helpers.String(account.EnderecoCobranca),
		"bairro_cobranca":         helpers.String(account.BairroCobranca),
		"cidade_cobranca":         helpers.String(account.CidadeCobranca),
		"uf_cobranca":             helpers.String(account.UfCobranca),
		"cep_cobranca":            helpers.String(account.CepCobranca),
		"data_hora_inclusao":      helpers.String(account.ModifiedAt),
		"usuario_inclusao":        helpers.String(account.CreatedById),
		"data_hora_alteracao":     helpers.String(account.ModifiedAt),
		"usuario_alteracao":       helpers.String(account.ModifiedById),
		"ende_latitude":           helpers.Float64(account.EndeLatitude),
		"ende_longitude":          helpers.Float64(account.EndeLongitude),
		"categoria_cliente":       helpers.String(account.CategoriaCliente),
		"opt_simples":             helpers.StringBoolean(account.OptSimples),
		"contrib_icms":            helpers.StringBoolean(account.ContribIcms),
		"consumidor_final":        helpers.StringBoolean(account.ConsumidorFinal),
		"ck_ver_vencto_lote":      helpers.StringBoolean(account.CkVerVenctoLote),
		"qtde_min_vencto_lote":    helpers.String(account.QtdeMinVenctoLote),
		"inss_ret":                helpers.StringBoolean(account.InssRet),
		"isento_icms":             helpers.StringBoolean(account.IsentoIcms),
		"ret_piscofcsll":          helpers.StringBoolean(account.RetPiscoFcsvll),
		"ret_iss":                 helpers.StringBoolean(account.RetIss),
		"ret_iss_fonte":           helpers.StringBoolean(account.RetIssFonte),
		"subst_tribut_icms":       helpers.StringBoolean(account.SubstTributIcms),
		"subst_tribut_pis":        helpers.StringBoolean(account.SubstTributPis),
		"subst_tribut_cofins":     helpers.StringBoolean(account.SubstTributConfis),
		"origem":                  helpers.String(account.Origem),
		"conta_contabil_clie_1":   helpers.String(account.ContaContabil),
		"conta_contabil_s_clie_1": helpers.String(account.ContaSintetica),
		"codigo_pessoa":           helpers.String(account.Ti9Codigo),
		"telefone_cobranca":       helpers.String(account.TelefoneCobranca),
	})

	_, insertErr := tx.ExecContext(ctx, sqlInsert)

	//Debug
	//utils.LogFile("CRM/DEBUG", " insertCfr", "CRITICAL ", helpers.String(account.Id), sqlInsert)

	if insertErr != nil {
		//utils.LogFile("CRM/CFR", " insert", "CRITICAL ", insertErr.Error(), sqlInsert)
		utils.LogDatabaseDetails("ACCOUNT", helpers.String(account.Id), sqlInsert, insertErr.Error(), "")
		tx.Rollback()
		return insertErr
	}

	commit := tx.Commit()

	if commit != nil {
		//utils.LogFile("CRM/CFR", " update", "CRITICAL ", commit.Error(), "erro de commit")
		utils.LogDatabaseDetails("ACCOUNT", helpers.String(account.Id), "COMMIT", commit.Error(), "")
		return commit
	}

	return nil
}

// ---------------------------------

func (t CfrRepositoryDbErp) SelectErp() ([]cfr.Cfr, error) {

	queryString := utils.Msg(`SELECT          
							tipo,  
							codigo_pessoa as id,                                   
							nome_pessoa as name,
							nome_pessoa as description,
							nome_fantasia as nomefantasia,
							natureza_pessoa as naturezapessoa,
							cnpj_cpf as sic_code,
							inscricao_estadual as inscricaoestadual,
							sitfed AS sitfed,
							siticms AS siticms,
							sitipi AS sitipi,
							endereco AS billing_address_street,
							numeroend AS numeroend,
							end_complemento AS endcomplemento,
							bairro AS BAIRRO,
							uf AS billing_address_state,
							cep AS billing_address_postal_code,
							pais AS billing_address_country,
							telefone_1 AS telefone1,
							telefone_2 AS telefone2,
							contato,
							email,
							data_nascimento AS datanascimento,
							emailnfe,
							home_page AS website,
							nosso_cliente AS nossocliente, -- não tem
							'f' AS nosso_fornecedor,
							'f' AS nosso_transportador,  -- não tem
							categoria_cliente AS industry,
							cond_pg_padrao_clie AS condpgpadraoclie,
							docpadrao_cliente   AS docpadraocliente,
							cobrpadrao_cliente  AS cobrpadraocliente,
							portpadrao_cliente  AS portpadraocliente,
							data_cad AS created_at,
							zonafranca AS zonafranca,
							inscricao_suframa,  -- não tem
							vendedor1, -- não tem
							comissao1_fat,-- não tem
							transp_padr_clie, -- não tem
							tipo_frete, -- não tem
							endereco,
							endereco_cobranca AS enderecocobranca,
							numeroend_cobranca AS numeroendcobranca,
							bairro_cobranca AS bairroCobranca,
							cidade_cobranca AS cidade_cobranca,
							uf_cobranca,
							cep,
							cep_cobranca,
							usuario_inclusao,
							data_hora_alteracao,
							usuario_alteracao,
							ende_latitude,
							ende_longitude,
							'' as TI9CODIGO,
							'001' AS categoria_cliente,
							'f' AS opt_simples,
							't' AS contrib_icms,
							'f' AS consumidor_final,
							'f' AS ck_ver_vencto_lote,
							0 AS qtde_min_vencto_lote,
							'f' AS inss_ret,
							'f' AS isento_icms,
							'f' AS ret_piscofcsll,
							'f' AS ret_iss,
							'f' AS ret_iss_fonte,
							'f' AS subst_tribut_icms,
							'f' AS subst_tribut_pis,
							'f' AS subst_tribut_cofins, 
							status,
							origem,
							'' as conta_contabil,
							'' as conta_sintetica,
							cidade,
							uf,
							telefone_cobranca
						from cfr
						inner join tb_crm_sincroniza ON codigo_pessoa = pk AND tabela = 'CFR'`, map[string]interface{}{})

	rows, err := t.db.Query(queryString)
	//fmt.Println(queryString)

	if err != nil {
		return nil, err
	}

	cfrs := []cfr.Cfr{}

	for rows.Next() {
		cfr := cfr.Cfr{}

		if err := rows.Scan(&cfr.Tipo,
			&cfr.Id,
			&cfr.Name,
			&cfr.Description,
			&cfr.NomeFantasia,
			&cfr.NaturezaPessoa,
			&cfr.SicCode,
			&cfr.InscricaoEstadual,
			&cfr.SitFed,
			&cfr.SitCms,
			&cfr.SitIpi,
			&cfr.BillingAddressCity,
			&cfr.NumeroEnd,
			&cfr.EndComplemento,
			&cfr.Bairro,
			&cfr.BillingAdressState,
			&cfr.BillingAdressPostalCode,
			&cfr.BillingAdressCountry,
			&cfr.Telefone1,
			&cfr.Telefone2,
			&cfr.Contato,
			&cfr.Email,
			&cfr.DataNascimento,
			&cfr.EmailNfe,
			&cfr.WebSite,
			&cfr.NossoCliente,
			&cfr.NossoFornecedor,
			&cfr.NossoTransportador,
			&cfr.Industry,
			&cfr.CondPgPadraoClie,
			&cfr.DocPadraoCliente,
			&cfr.CobrPadraoCliente,
			&cfr.PortPadraoCliente,
			&cfr.CreatedAt,
			&cfr.ZonaFranca,
			&cfr.InscricaoSuframa,
			&cfr.Vendedor1,
			&cfr.Comissao1Fat,
			&cfr.TranspPadraoCliente,
			&cfr.TipoFrete,
			&cfr.Endereco,
			&cfr.EnderecoCobranca,
			&cfr.NumeroEndCobranca,
			&cfr.BairroCobranca,
			&cfr.CidadeCobranca,
			&cfr.UfCobranca,
			&cfr.Cep,
			&cfr.CepCobranca,
			&cfr.CreatedById,
			&cfr.ModifiedAt,
			&cfr.ModifiedById,
			&cfr.EndeLatitude,
			&cfr.EndeLongitude,
			&cfr.Ti9Codigo,
			&cfr.CategoriaCliente,
			&cfr.OptSimples,
			&cfr.ContribIcms,
			&cfr.ConsumidorFinal,
			&cfr.CkVerVenctoLote,
			&cfr.QtdeMinVenctoLote,
			&cfr.InssRet,
			&cfr.IsentoIcms,
			&cfr.RetPiscoFcsvll,
			&cfr.RetIss,
			&cfr.RetIssFonte,
			&cfr.SubstTributIcms,
			&cfr.SubstTributPis,
			&cfr.SubstTributConfis,
			&cfr.Status,
			&cfr.Origem,
			&cfr.ContaContabil,
			&cfr.ContaSintetica,
			&cfr.Cidade,
			&cfr.Uf,
			&cfr.TelefoneCobranca,
		); err != nil {
			log.Println(err.Error())
			utils.LogDatabaseDetails("ACCOUNT", helpers.String(cfr.Id), queryString, err.Error(), "")
			return nil, err
		}
		cfrs = append(cfrs, cfr)
	}
	return cfrs, nil
}

func (t CfrRepositoryDbErp) CheckUpdateCrm(id string, owner string) (bool, error) {
	queryString := fmt.Sprintf(`SELECT count(*) from %s.account where id= '%s'`, owner, id)

	rows, err := t.db.Query(queryString)

	if err != nil {
		utils.LogDatabaseDetails("ACCOUNT", id, queryString, err.Error(), "")
		return false, err
	}

	var count int

	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			utils.LogFile("CRM/CFR", " check", "CRITICAL ", err.Error(), queryString)
			return false, err
		}
	}

	if count == 0 {
		return false, nil
	}
	return true, nil
}

func (t CfrRepositoryDbErp) UpdateCrm(cfr cfr.Cfr, owner string) error {
	queryString := utils.Msg(`UPDATE {{.owner}}.account SET
											name = '{{.name}}',
											description = '{{.description}}',
											nomefantasia = '{{.nomefantasia}}',
											naturezapessoa = '{{.naturezapessoa}}',
											sic_code = '{{.sic_code}}',
											inscricaoestadual= '{{.inscricaoestadual}}',
											sitfed= '{{.sitfed}}',
											siticms = '{{.siticms}}',
											sitipi = '{{.sitipi}}',
											billing_address_street = '{{.billing_address_street}}',
											numeroend = '{{.numeroend}}',
											endcomplemento= '{{.endcomplemento}}',
											bairro = '{{.bairro}}',
											billing_address_city = '{{.billing_address_city}}',
											billing_address_state = '{{.billing_address_state}}',
											billing_address_postal_code = '{{.billing_address_postal_code}}',
											billing_address_country = '{{.billing_address_country}}',
											telefone1 = '{{.telefone1}}',
											telefone2= '{{.telefone2}}',
											contato = '{{.contato}}',
											email = '{{.email}}',
											emailnfe = '{{.emailnfe}}',
											website = '{{.website}}',
											nossocliente = '{{.nossocliente}}',
											industry = '{{.industry}}',
											condpgpadraoclie = '{{.condpgpadraoclie}}',
											docpadraocliente = '{{.docpadraocliente}}',
											cobrpadraocliente = '{{.cobrpadraocliente}}',
											portpadraocliente = '{{.portpadraocliente}}',
											status = '{{.status}}',
											created_at = '{{.created_at}}',
											zonafranca = '{{.zonafranca}}',
											inscricaosuframa = '{{.inscricaosuframa}}',
											vendedor1 = '{{.vendedor1}}',
											comissao1fat = '{{.comissao1fat}}',
											transppadrclie = '{{.transppadrclie}}',
											tipofrete = '{{.tipofrete}}',
											enderecocobranca = '{{.enderecocobranca}}',
											bairrocobranca= '{{.bairrocobranca}}',
											cidadecobranca = '{{.cidadecobranca}}',
											ufcobranca = '{{.ufcobranca}}',
											cepcobranca = '{{.cepcobranca}}',
											created_by_id = '{{.created_by_id}}',
											modified_at = '{{.modified_at}}',
											modified_by_id = '{{.modified_by_id}}',
											endelatitude = '{{.endelatitude}}',
											endelongitude = '{{.endelongitude}}',
											endereco = '{{.endereco}}',
											bairro = '{{.bairro}}',
											cep = '{{.cep}}'
                                        WHERE TI9CODIGO = '{{.TI9CODIGO}}' `, map[string]interface{}{
		"owner":                       owner,
		"name":                        helpers.String(cfr.Name),
		"description":                 helpers.String(cfr.Description),
		"nomefantasia":                helpers.String(cfr.NomeFantasia),
		"naturezapessoa":              helpers.String(cfr.NaturezaPessoa),
		"cnpj_cpf":                    helpers.String(cfr.SicCode),
		"inscricaoestadual":           helpers.String(cfr.InscricaoEstadual),
		"sitfed":                      helpers.String(cfr.SitFed),
		"siticms":                     helpers.String(cfr.SitCms),
		"sitipi":                      helpers.String(cfr.SitIpi),
		"billing_address_street":      helpers.String(cfr.Endereco),
		"numeroend":                   helpers.String(cfr.NumeroEnd),
		"endcomplemento":              helpers.String(cfr.EndComplemento),
		"bairro":                      helpers.String(cfr.Bairro),
		"billing_address_city":        helpers.String(cfr.Cidade),
		"billing_address_state":       helpers.String(cfr.Uf),
		"billing_address_postal_code": helpers.String(cfr.Cep),
		"billing_address_country":     helpers.String(cfr.BillingAdressCountry),
		"telefone1":                   helpers.String(cfr.Telefone1),
		"telefone2":                   helpers.String(cfr.Telefone2),
		"contato":                     helpers.String(cfr.Contato),
		"email":                       helpers.String(cfr.Email),
		"emailnfe":                    helpers.String(cfr.EmailNfe),
		"website":                     helpers.String(cfr.WebSite),
		"nossocliente":                helpers.StringBoolean(cfr.NossoCliente),
		"industry":                    helpers.StringBoolean(cfr.Industry),
		"condpgpadraoclie":            helpers.String(cfr.CondPgPadraoClie),
		"docpadraocliente":            helpers.String(cfr.DocPadraoCliente),
		"cobrpadraocliente":           helpers.String(cfr.CobrPadraoCliente),
		"portpadraocliente":           helpers.String(cfr.PortPadraoCliente),
		"status":                      helpers.String(cfr.Status),
		"created_at":                  helpers.String(cfr.CreatedAt),
		"zonafranca":                  helpers.StringBoolean(cfr.ZonaFranca),
		"inscricaosuframa":            helpers.String(cfr.InscricaoSuframa),
		"vendedor1":                   helpers.String(cfr.Vendedor1),
		"comissao1fat":                helpers.Float64(cfr.Comissao1Fat),
		"transppadrclie":              helpers.String(cfr.TranspPadraoCliente),
		"tipofrete":                   helpers.String(cfr.TipoFrete),
		"enderecocobranca":            helpers.String(cfr.EnderecoCobranca),
		"bairrocobranca":              helpers.String(cfr.BairroCobranca),
		"cidadecobranca":              helpers.String(cfr.CidadeCobranca),
		"ufcobranca":                  helpers.String(cfr.UfCobranca),
		"cepcobranca":                 helpers.String(cfr.CepCobranca),
		"created_by_id":               helpers.String(cfr.CreatedById),
		"modified_at":                 helpers.String(cfr.ModifiedAt),
		"modified_by_id":              helpers.String(cfr.ModifiedById),
		"endelatitude":                helpers.Float64(cfr.EndeLatitude),
		"endelongitude":               helpers.Float64(cfr.EndeLongitude),
		//"categoria_cliente":           helpers.String(cfr.CategoriaCliente),
		//"opt_simples":                 helpers.StringBoolean(cfr.OptSimples),
		//"contrib_icms":                helpers.StringBoolean(cfr.ContribIcms),
		//"consumidor_final":            helpers.StringBoolean(cfr.ConsumidorFinal),
		//"ck_ver_vencto_lote":          helpers.StringBoolean(cfr.CkVerVenctoLote),
		//"qtde_min_vencto_lote":        helpers.String(cfr.QtdeMinVenctoLote),
		//"inss_ret":                    helpers.StringBoolean(cfr.InssRet),
		//"isento_icms":                 helpers.StringBoolean(cfr.IsentoIcms),
		//"ret_piscofcsll":              helpers.StringBoolean(cfr.RetPiscoFcsvll),
		//"ret_iss":                     helpers.StringBoolean(cfr.RetIss),
		//"ret_iss_fonte":               helpers.StringBoolean(cfr.RetIssFonte),
		//"subst_tribut_icms":           helpers.StringBoolean(cfr.SubstTributIcms),
		//"subst_tribut_pis":            helpers.StringBoolean(cfr.SubstTributPis),
		//"subst_tribut_cofins":         helpers.StringBoolean(cfr.SubstTributConfis),
		//"origem":                      helpers.String(cfr.Origem),
		//"conta_contabil_clie_1":   helpers.String(cfr.ContaContabil),
		//"conta_contabil_s_clie_1": helpers.String(cfr.ContaSintetica),
		"TI9CODIGO": helpers.String(cfr.Id),
		"endereco":  helpers.String(cfr.Endereco),
		"cidade":    helpers.String(cfr.Cidade),
		"cep":       helpers.String(cfr.Cep),
	})

	//Iniciando uma transação
	ctx := context.Background()
	tx, _ := t.db.BeginTx(ctx, nil)

	_, err := tx.ExecContext(ctx, queryString)

	////Debug
	//utils.LogFile("CRM/DEBUG", " updateCfr", "DEBUG ", helpers.String(account.Id), queryString)

	if err != nil {
		//utils.LogFile("CRM/CFR", " update", "CRITICAL ", err.Error(), queryString)
		utils.LogDatabaseDetails("ACCOUNT", helpers.String(cfr.Id), queryString, err.Error(), "")
		tx.Rollback()
		return err
	}

	commit := tx.Commit()

	if commit != nil {
		utils.LogFile("CRM/CFR", " update", "CRITICAL ", err.Error(), "erro no commmit")
		return commit
	}

	return nil
}

func (t CfrRepositoryDbErp) DeleteErp(id string, tipo string) error {
	queryString := utils.Msg(`DELETE from tb_crm_sincroniza WHERE tabela = 'CFR' AND pk = '{{.pk}}' and tipo = '{{.tipo}}'`, map[string]interface{}{
		"pk":   id,
		"tipo": tipo,
	})
	fmt.Println(queryString)

	//Iniciando o contexto de transação
	ctx := context.Background()
	tx, _ := t.db.BeginTx(ctx, nil)

	_, err := tx.ExecContext(ctx, queryString)

	if err != nil {
		//utils.LogFile("CRM/CFR", " delete", "CRITICAL ", err.Error(), queryString)
		utils.LogDatabaseDetails("ACCOUNT", id, queryString, err.Error(), "")
		tx.Rollback()
		return err
	}

	commit := tx.Commit()

	if commit != nil {
		//utils.LogFile("CRM/CFR", " delete", "CRITICAL ", commit.Error(), queryString)
		utils.LogDatabaseDetails("ACCOUNT", id, queryString, commit.Error(), "")
		return err
	}

	return nil
}

func (t CfrRepositoryDbErp) InsertCrm(cfr cfr.Cfr, owner string) error {
	queryString := utils.Msg(`INSERT INTO {{.owner}}.account (
										id,
                                        name ,
                                        description ,
                                        nomefantasia ,
                                        naturezapessoa ,
                                        sic_code ,
                                        inscricaoestadual ,
                                        sitfed ,
                                        siticms ,
                                        sitipi ,
                                        billing_address_street ,
                                        numeroend ,
                                        endcomplemento ,
                                        bairro ,
                                        billing_address_city ,
                                        billing_address_state ,
                                        billing_address_postal_code ,
                                        billing_address_country ,
                                        telefone1 ,
                                        telefone2 ,
                                        contato ,
                                        email ,
                                        emailnfe ,
                                        website ,
                                        nossocliente ,
                                        industry ,
                                        condpgpadraoclie ,
                                        docpadraocliente ,
                                        cobrpadraocliente ,
                                        portpadraocliente ,
                                        status ,
                                        created_at ,
                                        zonafranca ,
                                        inscricaosuframa ,
                                        vendedor1 ,
                                        comissao1fat ,
                                        transppadrclie ,
                                        tipofrete ,
                                        enderecocobranca ,
                                        bairrocobranca ,
                                        cidadecobranca ,
                                        ufcobranca ,
                                        cepcobranca ,
                                        created_by_id ,
                                        modified_at ,
                                        modified_by_id ,
                                        endelatitude ,
                                        endelongitude ,
                                        endereco ,
                                        cep 
										)
										VALUES(
										'{{.id}}',
										'{{.name}}',
										'{{.description}}',
										'{{.nomefantasia}}',
										'{{.naturezapessoa}}',
										'{{.sic_code}}',
										'{{.inscricaoestadual}}',
										'{{.sitfed}}',
										'{{.siticms}}',
										'{{.sitipi}}',
										'{{.billing_address_street}}',
										'{{.numeroend}}',
										'{{.endcomplemento}}',
										'{{.bairro}}',
										'{{.billing_address_city}}',
										'{{.billing_address_state}}',
										'{{.billing_address_postal_code}}',
										'{{.billing_address_country}}',
										'{{.telefone1}}',
										'{{.telefone2}}',
										'{{.contato}}',
										'{{.email}}',
										'{{.emailnfe}}',
										'{{.website}}',
										{{.nossocliente}},
										'{{.industry}}',
										'{{.condpgpadraoclie}}',
										'{{.docpadraocliente}}',
										'{{.cobrpadraocliente}}',
										'{{.portpadraocliente}}',
										'{{.status}}',
										'{{.created_at}}',
										'{{.zonafranca}}',
										'{{.inscricaosuframa}}',
										'{{.vendedor1}}',
										'{{.comissao1fat}}',
										'{{.transppadrclie}}',
										'{{.tipofrete}}',
										'{{.enderecocobranca}}',
										'{{.bairrocobranca}}',
										'{{.cidadecobranca}}',
										'{{.ufcobranca}}',
										'{{.cepcobranca}}',
										'{{.created_by_id}}',
										'{{.modified_at}}',
										'{{.modified_by_id}}',
										'{{.endelatitude}}',
										'{{.endelongitude}}',
										'{{.endereco}}',
										'{{.cep}}'
										)`, map[string]interface{}{
		"owner":                       owner,
		"id":                          helpers.String(cfr.Id),
		"name":                        helpers.String(cfr.Name),
		"description":                 helpers.String(cfr.Description),
		"nomefantasia":                helpers.String(cfr.NomeFantasia),
		"naturezapessoa":              helpers.String(cfr.NaturezaPessoa),
		"cnpj_cpf":                    helpers.String(cfr.SicCode),
		"inscricaoestadual":           helpers.String(cfr.InscricaoEstadual),
		"sitfed":                      helpers.String(cfr.SitFed),
		"siticms":                     helpers.String(cfr.SitCms),
		"sitipi":                      helpers.String(cfr.SitIpi),
		"billing_address_street":      helpers.String(cfr.Endereco),
		"numeroend":                   helpers.String(cfr.NumeroEnd),
		"endcomplemento":              helpers.String(cfr.EndComplemento),
		"bairro":                      helpers.String(cfr.Bairro),
		"billing_address_city":        helpers.String(cfr.Cidade),
		"billing_address_state":       helpers.String(cfr.Uf),
		"billing_address_postal_code": helpers.String(cfr.Cep),
		"billing_address_country":     helpers.String(cfr.BillingAdressCountry),
		"telefone1":                   helpers.String(cfr.Telefone1),
		"telefone2":                   helpers.String(cfr.Telefone2),
		"contato":                     helpers.String(cfr.Contato),
		"email":                       helpers.String(cfr.Email),
		"emailnfe":                    helpers.String(cfr.EmailNfe),
		"website":                     helpers.String(cfr.WebSite),
		"nossocliente":                helpers.StringBooleanInt(cfr.NossoCliente),
		"industry":                    helpers.StringBoolean(cfr.Industry),
		"condpgpadraoclie":            helpers.String(cfr.CondPgPadraoClie),
		"docpadraocliente":            helpers.String(cfr.DocPadraoCliente),
		"cobrpadraocliente":           helpers.String(cfr.CobrPadraoCliente),
		"portpadraocliente":           helpers.String(cfr.PortPadraoCliente),
		"status":                      helpers.String(cfr.Status),
		"created_at":                  helpers.StringDatetime(cfr.CreatedAt),
		"zonafranca":                  helpers.StringBoolean(cfr.ZonaFranca),
		"inscricaosuframa":            helpers.String(cfr.InscricaoSuframa),
		"vendedor1":                   helpers.String(cfr.Vendedor1),
		"comissao1fat":                helpers.Float64(cfr.Comissao1Fat),
		"transppadrclie":              helpers.String(cfr.TranspPadraoCliente),
		"tipofrete":                   helpers.String(cfr.TipoFrete),
		"enderecocobranca":            helpers.String(cfr.EnderecoCobranca),
		"bairrocobranca":              helpers.String(cfr.BairroCobranca),
		"cidadecobranca":              helpers.String(cfr.CidadeCobranca),
		"ufcobranca":                  helpers.String(cfr.UfCobranca),
		"cepcobranca":                 helpers.StringInt(cfr.CepCobranca),
		"created_by_id":               helpers.String(cfr.CreatedById),
		"modified_at":                 helpers.StringDatetime(cfr.ModifiedAt),
		"modified_by_id":              helpers.String(cfr.ModifiedById),
		"endelatitude":                helpers.Float64(cfr.EndeLatitude),
		"endelongitude":               helpers.Float64(cfr.EndeLongitude),
		//"categoria_cliente":           helpers.String(cfr.CategoriaCliente),
		//"opt_simples":                 helpers.StringBoolean(cfr.OptSimples),
		//"contrib_icms":                helpers.StringBoolean(cfr.ContribIcms),
		//"consumidor_final":            helpers.StringBoolean(cfr.ConsumidorFinal),
		//"ck_ver_vencto_lote":          helpers.StringBoolean(cfr.CkVerVenctoLote),
		//"qtde_min_vencto_lote":        helpers.String(cfr.QtdeMinVenctoLote),
		//"inss_ret":                    helpers.StringBoolean(cfr.InssRet),
		//"isento_icms":                 helpers.StringBoolean(cfr.IsentoIcms),
		//"ret_piscofcsll":              helpers.StringBoolean(cfr.RetPiscoFcsvll),
		//"ret_iss":                     helpers.StringBoolean(cfr.RetIss),
		//"ret_iss_fonte":               helpers.StringBoolean(cfr.RetIssFonte),
		//"subst_tribut_icms":           helpers.StringBoolean(cfr.SubstTributIcms),
		//"subst_tribut_pis":            helpers.StringBoolean(cfr.SubstTributPis),
		//"subst_tribut_cofins":         helpers.StringBoolean(cfr.SubstTributConfis),
		//"origem":                      helpers.String(cfr.Origem),
		//"conta_contabil_clie_1":   helpers.String(cfr.ContaContabil),
		//"conta_contabil_s_clie_1": helpers.String(cfr.ContaSintetica),
		"TI9CODIGO": helpers.String(cfr.Id),
		"endereco":  helpers.String(cfr.Endereco),
		"cidade":    helpers.String(cfr.Cidade),
		"cep":       helpers.String(cfr.Cep),
	})
	fmt.Println(queryString)

	//Debug
	//fmt.Println("AQUI é :", helpers.StringBoolean(cfr.NossoCliente))
	//fmt.Println("AQUI é :", cfr.NossoCliente)

	// Iniciando uma transação
	ctx := context.Background()
	tx, _ := t.db.BeginTx(ctx, nil)

	_, err := tx.ExecContext(ctx, queryString)

	if err != nil {
		//utils.LogFile("CRM/CFR", " insert", "CRITICAL ", err.Error(), queryString)
		utils.LogDatabaseDetails("ACCOUNT", helpers.String(cfr.Id), queryString, err.Error(), "")
		tx.Rollback()
		return err
	}

	//Checando as linhas afetadas
	//rowsAffected, _ := r.RowsAffected()
	//fmt.Println("Linhas afetadas : ", rowsAffected)

	commit := tx.Commit()

	if commit != nil {
		//utils.LogFile("CRM/CFR", " insert", "CRITICAL ", err.Error(), "erro no commmit")
		utils.LogDatabaseDetails("ACCOUNT", helpers.String(cfr.Id), "COMMIT", err.Error(), "")
		return commit
	}

	return nil
}
