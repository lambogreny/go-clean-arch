package erp_crm

import (
	"fmt"
	crmRepository "github.com/augusto/imersao5-esquenta-go/adapter/repository/crm"
	"github.com/augusto/imersao5-esquenta-go/services/crm"
	"github.com/augusto/imersao5-esquenta-go/usecase/crm/prd"
	"log"
)

func PrdService(clientId string) error {
	log.Println("Inicio da transação dos procedimento de integração da PRD")

	//Chama a função que retorna as duas conexões
	dbCrmConn, dbErpConn, _, connError := crm.ServicesDatabases(clientId)

	if connError != nil {
		return connError
	}

	//Criando o repositório
	repoErp := crmRepository.NewPrdRepositoryDbErp(dbErpConn)
	repoCrm := crmRepository.NewPrdRepositoryDbErp(dbCrmConn)

	usecaseErp := prd.NewProcessPrd(repoErp)
	usecaseCrm := prd.NewProcessPrd(repoCrm)

	data, err := usecaseErp.UseCaseSelect()

	if err != nil {
		return err
	}

	for _, x := range data {

		//Checando o se deve realizar o update
		checkUpdate, err := usecaseCrm.UseCaseCheckUpdateCrm(x.Codigo_produto)

		if err != nil {
			return err
		}

		switch checkUpdate {
		case true:
			fmt.Println("Irei realizar a operação de update!")
			err := usecaseCrm.UseCaseUpdate(x)
			if err != nil {
				return err
			}

			delete := usecaseErp.UseCaseDelete(x.Codigo_produto, x.Tipo)
			if delete != nil {
				return delete
			}

		case false:
			//#TODO Implementar o insert (Preciso de dados para teste)
			fmt.Println("Irei realizar a operação de insert!")
		}

	}

	return nil
}
