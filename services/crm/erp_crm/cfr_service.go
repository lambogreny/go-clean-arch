package erp_crm

import (
	"fmt"
	crmRepository "github.com/augusto/imersao5-esquenta-go/adapter/repository/crm"
	cfr2 "github.com/augusto/imersao5-esquenta-go/entity/crm/cfr"
	"github.com/augusto/imersao5-esquenta-go/services/crm"
	"github.com/augusto/imersao5-esquenta-go/usecase/crm/cfr"
	"github.com/augusto/imersao5-esquenta-go/utils"
	"github.com/augusto/imersao5-esquenta-go/utils/helpers"
	"log"
)

func CfrService(clientId string) error {
	/*
	*Crm : Todas as funções que operam no banco do CRM
	*Erp : Todas as funções que operam no banco do ERP
	 */
	log.Println("Início da transação do procedimento de integração com a CFR")

	//Chama a função que retorna as duas conexões
	dbCrmConn, dbErpConn, ownerCrm, connError := crm.ServicesDatabases(clientId)

	//fmt.Println("O dbCrm é : ", dbCrmConn)
	//fmt.Println("O dbErpConn é : ", dbErpConn.Stats())

	if connError != nil {
		return connError
	}

	//Criando o repositório
	repoCrm := crmRepository.NewCfrRepositoryDbErp(dbCrmConn)
	repoErp := crmRepository.NewCfrRepositoryDbErp(dbErpConn)

	//repoErp := crmRepository.NewCfrRepositoryDbErp(dbErpConn)
	//repoCrm := crmRepository.NewCfrRepositoryDbErp(dbCrmConn)

	usecaseCrm := cfr.NewProcessCfr(repoCrm)
	usecaseErp := cfr.NewProcessCfr(repoErp)

	// ----------------------------------------------------------- Selecionando todas as linhas ------------------------------------------------------------//

	data, err := usecaseCrm.UseCaseSelect(ownerCrm)
	if err != nil {
		return err
	}
	//fmt.Println(data)

	// ----------------------------------------------------------- Percorrendo as linhas e definindo as ações ------------------------------------------------------------//

	for _, x := range data {
		switch helpers.String(x.Tipo) {
		case "I":
			IErr := CfrInsertWithCheck(usecaseCrm, usecaseErp, x, ownerCrm)
			if IErr != nil {
				utils.LogDatabase(clientId, "CFR", "I", helpers.String(x.Id), true, IErr.Error())
				return IErr
			}
			//fmt.Println("Inserir, porém checando o Ti9")
		case "U":
			//fmt.Println("Atualizar")
		}

	}
	return nil
}

func CfrInsertWithCheck(usecaseCrm *cfr.ProcessCfr, usecaseErp *cfr.ProcessCfr, x cfr2.Account, crmOwner string) error {
	//Checando o se deve realizar o update
	checkUpdate, err := usecaseErp.UserCaseCheckUpdateErp(helpers.String(x.Id))
	if err != nil {
		return err
	}

	switch checkUpdate {
	case true:
		fmt.Println("Atualização")
		err := usecaseErp.UseCaseUpdate(x, crmOwner)

		if err != nil {
			return err
		}
	case false:
		fmt.Println("Insert")
	}
	fmt.Println(checkUpdate)
	return nil
}
