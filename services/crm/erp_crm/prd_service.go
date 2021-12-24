package erp_crm

import (
	crmRepository "github.com/augusto/imersao5-esquenta-go/adapter/repository/crm"
	prd2 "github.com/augusto/imersao5-esquenta-go/entity/crm/prd"
	"github.com/augusto/imersao5-esquenta-go/services/crm"
	"github.com/augusto/imersao5-esquenta-go/usecase/crm/prd"
	"github.com/augusto/imersao5-esquenta-go/utils"
	"log"
)

func PrdService(clientId string) error {
	log.Println("Inicio da transação dos procedimento de integração da PRD")

	//Chama a função que retorna as duas conexões
	dbCrmConn, dbErpConn, crmOwner, connError := crm.ServicesDatabases(clientId)

	if connError != nil {
		return connError
	}

	//Criando o repositório
	repoErp := crmRepository.NewPrdRepositoryDbErp(dbErpConn)
	repoCrm := crmRepository.NewPrdRepositoryDbErp(dbCrmConn)

	usecaseErp := prd.NewProcessPrd(repoErp)
	usecaseCrm := prd.NewProcessPrd(repoCrm)

	data, err := usecaseErp.UseCaseSelect()

	//Logando testes
	//utils.LogDatabase("PRD", "INFO", "123", false, "")

	if err != nil {
		return err
	}

	for _, x := range data {

		switch x.Tipo {
		case "I":
			err := InsertWithCheck(usecaseCrm, usecaseErp, x, crmOwner)
			if err != nil {
				utils.LogDatabase(clientId, "PRD", "I", x.Codigo_produto, true, err.Error())
				return err
			}

		case "U":
			err := Update(usecaseCrm, usecaseErp, x, crmOwner)
			if err != nil {
				utils.LogDatabase(clientId, "PRD", "I", x.Codigo_produto, true, err.Error())
				return err
			}
		}
		utils.LogDatabase(clientId, "PRD", x.Tipo, x.Codigo_produto, false, "")
	}

	return nil
}

func InsertWithCheck(usecaseCrm *prd.ProcessPrd, usecaseErp *prd.ProcessPrd, x prd2.Prd, crmOwner string) error {

	//Checando o se deve realizar o update
	checkUpdate, err := usecaseCrm.UseCaseCheckUpdateCrm(x.Codigo_produto)

	if err != nil {
		return err
	}

	switch checkUpdate {
	case true:
		err := usecaseCrm.UseCaseUpdate(x, crmOwner)
		if err != nil {
			return err
		}

		delete := usecaseErp.UseCaseDelete(x.Codigo_produto, x.Tipo)
		if delete != nil {
			return delete
		}

	case false:
		err := usecaseCrm.UseCaseInsert(x, crmOwner)

		if err != nil {
			return err
		}

		delete := usecaseErp.UseCaseDelete(x.Codigo_produto, x.Tipo)
		if delete != nil {
			return delete
		}
	}

	return nil
}

func Update(usecaseCrm *prd.ProcessPrd, usecaseErp *prd.ProcessPrd, x prd2.Prd, crmOwner string) error {

	err := usecaseCrm.UseCaseUpdate(x, crmOwner)
	if err != nil {
		return err
	}

	delete := usecaseErp.UseCaseDelete(x.Codigo_produto, x.Tipo)
	if delete != nil {
		return delete
	}

	return nil
}
