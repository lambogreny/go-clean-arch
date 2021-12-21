package crmRepository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/augusto/imersao5-esquenta-go/entity/crm/prd"
	"github.com/augusto/imersao5-esquenta-go/utils"
	"log"
)

type PrdRepositoryDbErp struct {
	db *sql.DB
}

func NewPrdRepositoryDbErp(db *sql.DB) *PrdRepositoryDbErp {
	return &PrdRepositoryDbErp{db: db}
}

func (t PrdRepositoryDbErp) CheckUpdateCrm(codigoProduto string) (bool, error) {
	queryString := fmt.Sprintf(`SELECT count(*) from epcrm_decorlit.product where codigoproduto = '%s' `, codigoProduto)

	rows, err := t.db.Query(queryString)

	if err != nil {
		return false, err
	}

	var count int

	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			log.Println("Erro ao buscar o produto!")
			return false, err
		}
	}

	if count == 0 {
		return false, nil
	}
	return true, nil

}

func (t PrdRepositoryDbErp) Update(prd prd.Prd, owner string) error {

	//Tratando a struct
	//prd.Descricao_produto = "Augusto"

	queryString := utils.Msg(`UPDATE {{.owner}}.product
										SET    NAME = '{{.NAME}}',
											   description = '{{.description}}',
											   codtipi = '{{.codtipi}}',
											   category_id = '{{.category_id}}',
											   category2 = '{{.category2}}',
											   category3 = '{{.category3}}',
											   category4 = '{{.category4}}',
											   unidade = '{{.unidade}}',
											   brand_id = '{{.brand_id}}',
											   created_at = '{{.created_at}}',
											   created_by_id = '{{.created_by_id}}',
											   cost_price = '{{.cost_price}}',
											   modified_at = '{{.modified_at}}',
											   modified_by_id = '{{.modified_by_id}}',
											   part_number = '{{.part_number}}',
											   status = '{{.status}}',
											   unit_price = '{{.unit_price}}',
											   weight = '{{.weight}}',
											   peso_bruto = '{{.peso_bruto}}'
										WHERE  codigoproduto = '{{.codigoproduto}}'`, map[string]interface{}{
		"owner":          owner,
		"NAME":           prd.Descricao_produto,
		"description":    prd.Descricao_produto,
		"codtipi":        prd.Cod_tipi,
		"category_id":    "",
		"category2":      "",
		"category3":      "",
		"category4":      "",
		"unidade":        prd.Unidade,
		"brand_id":       prd.Marca,
		"created_at":     prd.Data_cad.Format("2006-01-02 15:04:05"),
		"created_by_id":  prd.Usuario_inclusao,
		"cost_price":     prd.Ultimo_preco_liq,
		"modified_at":    prd.Data_hora_alteracao.Format("2006-01-02 15:04:05"),
		"modified_by_id": prd.Usuario_alteracao,
		"part_number":    prd.Partnumber,
		"status":         prd.Ativo,
		"unit_price":     prd.Ultimo_preco_liq,
		"weight":         prd.Peso_liquido,
		"peso_bruto":     prd.Peso_bruto,
		"codigoproduto":  prd.Codigo_produto,
	})

	//Limpando a queryString
	queryString = utils.CleanQueryString(queryString)

	//Iniciando o contexto de transação
	ctx := context.Background()
	tx, _ := t.db.BeginTx(ctx, nil)

	_, err := tx.ExecContext(ctx, queryString)

	//Log de controle
	//utils.LogFile("CRM/PRD", " update", "INFO ", "err.Error()", queryString)

	if err != nil {
		utils.LogFile("CRM/PRD", " update", "CRITICAL ", err.Error(), queryString)
		tx.Rollback()
		return err
	}

	commit := tx.Commit()

	if commit != nil {
		log.Println("Não consegui commitar!")
		return err
	}

	//Retirando esse check pois não considera se não houve nenhuma mudança (r do ExecContext)
	//rowsAffected, _ := r.RowsAffected()
	//if rowsAffected <= 0 {
	//	//utils.LogFile("CRM/PRD", "update", "CRITICAL ", err.Error(), queryString)
	//	tx.Rollback()
	//	return fmt.Errorf("Nenhuma linha afetada!!")
	//}

	return nil

}

func (t PrdRepositoryDbErp) Delete(codigoProduto string, tipo string) error {
	queryString := utils.Msg(`DELETE from tb_crm_sincroniza WHERE tabela = 'PRD' AND pk = '{{.pk}}' and tipo = '{{.tipo}}'`, map[string]interface{}{
		"pk":   codigoProduto,
		"tipo": tipo,
	})
	//Iniciando o contexto de transação
	ctx := context.Background()
	tx, _ := t.db.BeginTx(ctx, nil)

	_, err := tx.ExecContext(ctx, queryString)

	if err != nil {
		tx.Rollback()
		return err
	}
	fmt.Println(queryString)

	commit := tx.Commit()

	if commit != nil {
		log.Println("Não consegui commitar!")
		return err
	}

	return nil
}

func (t PrdRepositoryDbErp) Select() ([]prd.Prd, error) {
	queryString := fmt.Sprintf(`SELECT tipo,
										   codigo_produto,
										   descricao_produto,
										   cod_tipi,
										   grupo_estoque,
										   grupo_estoque_n2,
										   grupo_estoque_n3,
										   grupo_estoque_n4,
										   unidade,
										   marca,
										   data_cad,
										   usuario_inclusao,
										   ultimo_preco_liq,
										   data_hora_alteracao,
										   usuario_alteracao,
										   partnumber,
										   ativo,
										   peso_liquido,
										   peso_bruto
									FROM   prd
										   INNER JOIN tb_crm_sincroniza
												   ON codigo_produto = pk
													  AND tabela = 'PRD' `)

	rows, err := t.db.Query(queryString)
	fmt.Println(rows)

	if err != nil {
		utils.LogFile("ERROR", " prd", "CRITICAL ", err.Error(), queryString)
		log.Println("could not execute query: %v", err)

		return []prd.Prd{}, err
	}

	products := []prd.Prd{}

	for rows.Next() {
		product := prd.Prd{}

		if err := rows.Scan(&product.Tipo,
			&product.Codigo_produto,
			&product.Descricao_produto,
			&product.Cod_tipi,
			&product.Grupo_estoque,
			&product.Grupo_estoque_n2,
			&product.Grupo_estoque_n3,
			&product.Grupo_estoque_n4,
			&product.Unidade,
			&product.Marca,
			&product.Data_cad,
			&product.Usuario_inclusao,
			&product.Ultimo_preco_liq,
			&product.Data_hora_alteracao,
			&product.Usuario_alteracao,
			&product.Partnumber,
			&product.Ativo,
			&product.Peso_liquido,
			&product.Peso_bruto); err != nil {
			log.Println(err.Error())
			return []prd.Prd{}, err
		}
		products = append(products, product)
	}
	return products, nil

}

func (t PrdRepositoryDbErp) Insert(prd prd.Prd, owner string) error {
	fmt.Println("Camada de banco para realizar o insert!")

	queryString := utils.Msg(`INSERT INTO {{.owner}}.product
            (
                        id,
                        NAME,
                        description,
                        codtipi,
                        category_id,
                        category2,
                        category3,
                        category4,
                        unidade,
                        brand_id,
                        created_at,
                        created_by_id,
                        cost_price,
                        modified_at,
                        modified_by_id,
                        part_number,
                        status,
                        unit_price,
                        weight,
                        peso_bruto,
                        codigoproduto,
                        deleted
            )
            VALUES
            (
                        '{{.codigo_produto}}',
                        '{{.descricao_produto}}',
                        '{{.description}}',
                        '{{.codtipi}}',
                        '{{.category_id}}',
                        '{{.category2}}',
                        '{{.category3}}',
                        '{{.category4}}',
                        '{{.unidade}}',
                        '{{.brand_id}}',
                        '{{.created_at}}',
                        '{{.created_by_id}}',
                        '{{.cost_price}}',
                        '{{.modified_at}}',
                        '{{.modified_by_id}}',
                        '{{.part_number}}',
                        '{{.status}}',
                        '{{.unit_price}}',
                        '{{.weight}}',
                        '{{.peso_bruto}}',
                        '{{.codigoproduto}}',
                        '{{.deleted}}'

            )`, map[string]interface{}{
		"owner":             owner,
		"codigo_produto":    prd.Codigo_produto,
		"descricao_produto": prd.Descricao_produto,
		"codtipi":           prd.Cod_tipi,
		"category_id":       "",
		"category2":         "",
		"category3":         "",
		"category4":         "",
		"unidade":           prd.Unidade,
		"brand_id":          prd.Marca,
		"created_at":        prd.Data_cad.Format("2006-01-02 15:04:05"),
		"created_by_id":     prd.Usuario_inclusao,
		"cost_price":        prd.Ultimo_preco_liq,
		"modified_at":       prd.Data_hora_alteracao.Format("2006-01-02 15:04:05"),
		"modified_by_id":    prd.Usuario_alteracao,
		"part_number":       prd.Partnumber,
		"status":            prd.Ativo,
		"unit_price":        prd.Ultimo_preco_liq,
		"weight":            prd.Peso_bruto,
		"codigoproduto":     prd.Codigo_produto,
		"deleted":           "0",
	})

	queryString = utils.CleanQueryString(queryString)
	fmt.Println(queryString)

	//Iniciando o contexto de transação
	ctx := context.Background()
	tx, _ := t.db.BeginTx(ctx, nil)

	_, err := tx.ExecContext(ctx, queryString)

	if err != nil {
		utils.LogFile("CRM/PRD", " insert", "CRITICAL ", err.Error(), queryString)
		tx.Rollback()
		return err
	}

	commit := tx.Commit()

	if commit != nil {
		log.Println("Não consegui commitar!")
		return err
	}

	fmt.Println(queryString)
	return nil
}
