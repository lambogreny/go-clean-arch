package crmRepository

import (
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
	//TODO implement me
	panic("implement me")
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
