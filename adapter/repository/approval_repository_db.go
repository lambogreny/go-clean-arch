package repository

import (
	"database/sql"
	"github.com/augusto/imersao5-esquenta-go/entity"
	"github.com/augusto/imersao5-esquenta-go/utils"
	"log"
)

type ApprovalRepositoryDb struct {
	db *sql.DB
}

func NewApprovalRepositoryDb(db *sql.DB) *ApprovalRepositoryDb {
	return &ApprovalRepositoryDb{db: db}
}

func (t ApprovalRepositoryDb) Select() ([]entity.Approval, error) {
	queryString := "SELECT\n  *\nfrom\n  \tfn_aprov_cotacao_full ('ELISA') AS (\n    msg text,\n    tipo text,\n    responsavel text,\n    validacao text,\n    valor double precision,\n    campo1 text,\n    campo2 text,\n    campo4 text,\n    campo5 text\n  )"
	rows, err := t.db.Query(queryString)

	if err != nil {
		utils.LogFile("ERROR", " approval", "CRITICAL ", err.Error(), queryString)
		log.Println("could not execute query: %v", err)

		return nil, err
	}

	approvals := []entity.Approval{}

	for rows.Next() {
		approval := entity.Approval{}

		if err := rows.Scan(&approval.Msg, &approval.Tipo, &approval.Responsavel, &approval.Validacao, &approval.Valor, &approval.Campo1, &approval.Campo2, &approval.Campo4, &approval.Campo5); err != nil {
			//if err == sql.ErrTxDone {
			//	fmt.Println("Erro de ErrTxDone")
			//}
			//if err == sql.ErrConnDone {
			//	fmt.Println("Erro de ErrConnDone")
			//}
			log.Println(err.Error())
			log.Fatalf("could not scan row: %v", err)

			return nil, err
		}
		approvals = append(approvals, approval)
	}

	return approvals, nil
}
