package crmRepository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/augusto/imersao5-esquenta-go/entity/crm/cfr"
	"github.com/augusto/imersao5-esquenta-go/utils"
	"github.com/augusto/imersao5-esquenta-go/utils/helpers"
	"log"
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
		utils.LogFile("CRM/CFR", " update", "CRITICAL ", err.Error(), queryString)
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
		utils.LogFile("CRM/CFR", " delete", "CRITICAL ", err.Error(), queryString)
		tx.Rollback()
		return err
	}

	commit := tx.Commit()

	if commit != nil {
		utils.LogFile("CRM/PRD", " delete", "CRITICAL ", commit.Error(), queryString)
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
		return err
	case err != nil:
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
		utils.LogFile("CRM/CFR", " insert", "CRITICAL ", updateErr.Error(), sqlUpdateSysSequencial)
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
		utils.LogFile("CRM/CFR", " insert", "CRITICAL ", insertErr.Error(), sqlInsert)
		tx.Rollback()
		return insertErr
	}

	commit := tx.Commit()

	if commit != nil {
		utils.LogFile("CRM/CFR", " update", "CRITICAL ", commit.Error(), "erro de commit")
		return commit
	}

	return nil
}
