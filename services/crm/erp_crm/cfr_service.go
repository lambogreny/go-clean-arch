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
	"strings"
)

func CfrService(clientId string) error {
	/*
	*Crm : Todas as funções que operam no banco do CRM
	*Erp : Todas as funções que operam no banco do ERP
	 */
	log.Println("Início da transação do procedimento de integração com a CFR")

	//Chama a função que retorna as duas conexões
	dbCrmConn, dbErpConn, ownerCrm, connError := crm.ServicesDatabases(clientId)

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
				switch {
				//Para casos de duplicate key, apenas loga e continua o loop
				case strings.Contains(IErr.Error(), "duplicate key"):
					utils.LogDatabase(clientId, "CFR", "I", helpers.String(x.Id), true, IErr.Error())
				default:
					return IErr
				}
			}

		case "U":
			UErr := CfrUpdate(usecaseCrm, usecaseErp, x, ownerCrm)

			if UErr != nil {
				utils.LogDatabase(clientId, "CFR", "U", helpers.String(x.Id), true, UErr.Error())
				return UErr
			}
		}

		utils.LogDatabase(clientId, "CFR", helpers.String(x.Tipo), helpers.String(x.Id), false, "")

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

		UpdateErr := usecaseErp.UseCaseUpdate(x, crmOwner)
		if UpdateErr != nil {
			return UpdateErr
		}
		deleteErr := usecaseErp.UseCaseDelete(helpers.String(x.Id), crmOwner)

		if deleteErr != nil {
			return deleteErr
		}

	case false:
		fmt.Println("Insert")
		InsertErr := usecaseErp.UseCaseInsert(x, crmOwner)

		if InsertErr != nil {
			return InsertErr
		}
	}

	return nil
}

func CfrUpdate(usecaseCrm *cfr.ProcessCfr, usecaseErp *cfr.ProcessCfr, x cfr2.Account, crmOwner string) error {

	UpdateErr := usecaseErp.UseCaseUpdate(x, crmOwner)
	if UpdateErr != nil {
		return UpdateErr
	}

	return nil
}
