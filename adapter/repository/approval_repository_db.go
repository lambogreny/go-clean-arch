package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/augusto/imersao5-esquenta-go/entity"
	"github.com/augusto/imersao5-esquenta-go/utils"
)

type ApprovalRepositoryDb struct {
	db *sql.DB
}

func NewApprovalRepositoryDb(db *sql.DB) *ApprovalRepositoryDb {
	return &ApprovalRepositoryDb{db: db}
}

func (t ApprovalRepositoryDb) CheckPermission(usuario string) (bool, error) {
	queryString := fmt.Sprintf(`SELECT
										  permissao as v_tipo_aprovacao,
										  codigo_pessoa as v_prereq,
										  cencusto as v_devolve,
										  uso_geral as v_alcada,
										  login
										from
										  tb_permissao
										where
										  aplicacao = 'APCOT'
										  AND login = '%s' `, strings.ToUpper(usuario))
	var user string
	err := t.db.QueryRow(queryString).Scan(user)

	if err == sql.ErrNoRows {
		fmt.Println("Não encontrei resultados!!!")
		return false, nil
	}

	return true, nil
}

func (t ApprovalRepositoryDb) RecoverAllProviders(filial string, cotacao string) {

}

func (t ApprovalRepositoryDb) Interact(filial string, cotacao string, fornecedor string, tipoDeAprovacao string, usuario string, statusDeAprovacao string, justificativa string, seqConcatenada string) error {
	ctx := context.Background()
	tx, _ := t.db.BeginTx(ctx, nil)

	var treatBeforeInsertQuery string

	switch statusDeAprovacao {
	case "D":
		treatBeforeInsertQuery = fmt.Sprintf("UPDATE tb_aprov_cotacao_full set status_aprov = 'W' where filial = '%s' AND cotacao = '%s' AND  fornecedor = '%s' ", filial, cotacao, fornecedor)
	case "A", "R":
		treatBeforeInsertQuery = fmt.Sprintf("DELETE FROM tb_aprov_cotacao_full where  filial = '%s' AND cotacao = '%s' AND  fornecedor = '%s' AND status_aprov IN ('D','W')", filial, cotacao, fornecedor)
	}

	_, err := tx.ExecContext(ctx, treatBeforeInsertQuery)
	fmt.Println(treatBeforeInsertQuery)

	if err != nil {
		return err
	}

	insertString := fmt.Sprintf(`INSERT into
										  tb_aprov_cotacao_full
										values
										  (
											'%s', 
											'%s',
											'%s',
											'%s',
											'%s',
											'%s',
											CURRENT_TIMESTAMP,
											'%s',
											'%s' 
										  )`, filial, cotacao, fornecedor, tipoDeAprovacao, usuario, statusDeAprovacao, justificativa, seqConcatenada)

	_, err = tx.ExecContext(ctx, insertString)

	if err != nil {
		fmt.Println("Cai no rollback!!")
		tx.Rollback()
		return err
	}

	//Commitando a transação
	err = tx.Commit()

	if err != nil {
		fmt.Println("Erro ao commitar a transaction")
		return err
	}

	return nil
}

func (t ApprovalRepositoryDb) Select(user string) ([]entity.Approval, error) {
	queryString := fmt.Sprintf(`SELECT x.msg,
									x.tipo,
									x.validacao,
									Avg(x.valor),
									x.campo1,
									x.campo2,
									x.campo5
								FROM     (
										SELECT *
										FROM   fn_aprov_cotacao_full('%s', '4', '') AS ( msg text, tipo text, validacao text, valor DOUBLE PRECISION, campo1 text, campo2 text, campo5 text ) ) x
								GROUP BY x.msg,
									x.tipo,
									x.validacao,
									x.campo1,
									x.campo2,
									x.campo5
								ORDER BY x.msg limit 7`, user)
	rows, err := t.db.Query(queryString)

	if err != nil {
		utils.LogFile("ERROR", " approval", "CRITICAL ", err.Error(), queryString)
		log.Println("could not execute query: %v", err)

		return nil, err
	}

	approvals := []entity.Approval{}

	for rows.Next() {
		approval := entity.Approval{}

		if err := rows.Scan(&approval.Msg, &approval.Tipo, &approval.Validacao, &approval.Valor, &approval.Campo1, &approval.Campo2, &approval.Campo5); err != nil {

			log.Println(err.Error())
			log.Fatalf("could not scan row: %v", err)

			return nil, err
		}
		approvals = append(approvals, approval)
	}

	return approvals, nil
}
