package erp_crm

import (
	crmRepository "github.com/augusto/imersao5-esquenta-go/adapter/repository/crm"
	"github.com/augusto/imersao5-esquenta-go/services/crm"
	"github.com/augusto/imersao5-esquenta-go/usecase/crm/cfr"
	"log"
)

func CfrService(clientId string) error {
	log.Println("Início da transação do procedimento de integração com a CFR")

	//Chama a função que retorna as duas conexões
	dbCrmConn, dbErpConn, _, connError := crm.ServicesDatabases(clientId)

	if connError != nil {
		return connError
	}

	//Criando o repositório
	repoErp := crmRepository.NewCfrRepositoryDbErp(dbErpConn)
	//repoCrm := crmRepository.NewPrdRepositoryDbErp(dbCrmConn)

	usecaseErp := cfr.NewProcessCfr(repoErp)
	//usecaseCrm := prd.NewProcessPrd(repoCrm)

	data, err := usecaseErp.UseCaseSelect()
	return nil
}
