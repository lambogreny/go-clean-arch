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
	"time"
)

func CfrService(clientId string) error {
	/*
	*Crm : Todas as funções que operam no banco do CRM
	*Erp : Todas as funções que operam no banco do ERP
	 */
	log.Println("Início da transação do procedimento que leva os dados do CRM para o ERP")

	//Chama a função que retorna as duas conexões
	dbCrmConn, dbErpConn, ownerCrm, connError := crm.ServicesDatabases(clientId)

	if connError != nil {
		return connError
	}

	//Criando o repositório
	repoCrm := crmRepository.NewCfrRepositoryDbErp(dbCrmConn)
	repoErp := crmRepository.NewCfrRepositoryDbErp(dbErpConn)

	usecaseCrm := cfr.NewProcessCfr(repoCrm)
	usecaseErp := cfr.NewProcessCfr(repoErp)

	// ----------------------------------------------------------- Selecionando todas as linhas ------------------------------------------------------------//

	data, err := usecaseCrm.UseCaseSelect(ownerCrm)
	if err != nil {
		return err
	}
	fmt.Println("Dados de retorno : ", data)

	// ----------------------------------------------------------- Percorrendo as linhas e definindo as ações ------------------------------------------------------------//

	for i, x := range data {
		log.Println("Registro : ", i, " // id : ", x.Id)
		switch helpers.String(x.Tipo) {
		case "I":
			IErr := CfrInsertWithCheck(usecaseCrm, usecaseErp, x, ownerCrm, helpers.ExtraInfo{Tipo: helpers.String(x.Tipo)})

			if IErr != nil {
				switch {
				//Para casos de duplicate key, apenas loga e continua o loop
				case strings.Contains(IErr.Error(), "duplicate key"):
					log.Println("Cai no duplicate key")
					utils.LogDatabase(clientId, "CFR", "I", helpers.String(x.Id), true, IErr.Error())
					continue
				default:
					log.Println("Cai no erro default")
					return IErr
				}
			}

		case "U":
			UErr := CfrInsertWithCheck(usecaseCrm, usecaseErp, x, ownerCrm, helpers.ExtraInfo{Tipo: helpers.String(x.Tipo)})

			if UErr != nil {
				switch {
				//Para casos de duplicate key, apenas loga e continua o loop
				case strings.Contains(UErr.Error(), "duplicate key"):
					log.Println("Cai no duplicate key")
					utils.LogDatabase(clientId, "CFR", "I", helpers.String(x.Id), true, UErr.Error())
					continue
				default:
					log.Println("Cai no erro default")
					return UErr
				}
			}
		}

		utils.LogDatabase(clientId, "CFR", helpers.String(x.Tipo), helpers.String(x.Id), false, "")

		//Para não derrubar o banco
		time.Sleep(1 * time.Second)

	}
	return nil
}

func CfrInsertWithCheck(usecaseCrm *cfr.ProcessCfr, usecaseErp *cfr.ProcessCfr, x cfr2.Account, crmOwner string, extra helpers.ExtraInfo) error {
	log.Println("CASO : CFR INSERT WITH CHECK")
	//Checando o se deve realizar o update
	checkUpdate, err := usecaseErp.UserCaseCheckUpdateErp(helpers.String(x.Id))
	if err != nil {
		log.Println("Cai no erro de checagem!")
		return err
	}

	switch checkUpdate {
	case true:
		log.Println("Chequei o registro e cai no update")
		UpdateErr := usecaseErp.UseCaseUpdate(x, crmOwner)
		if UpdateErr != nil {
			return UpdateErr
		}
		deleteErr := usecaseCrm.UseCaseDelete(helpers.String(x.Id), crmOwner, extra.Tipo)

		if deleteErr != nil {
			return deleteErr
		}

	case false:
		log.Println("Chequei o registro e cai no Insert")
		InsertErr := usecaseErp.UseCaseInsert(x, crmOwner)

		if InsertErr != nil {
			return InsertErr
		}

		deleteErr := usecaseCrm.UseCaseDelete(helpers.String(x.Id), crmOwner, extra.Tipo)

		if deleteErr != nil {
			return deleteErr
		}
	}

	return nil
}

//Não está sendo utilizado
func CfrUpdate(usecaseCrm *cfr.ProcessCfr, usecaseErp *cfr.ProcessCfr, x cfr2.Account, crmOwner string, extra helpers.ExtraInfo) error {
	log.Println("CASO : CFR UPDATE")

	UpdateErr := usecaseErp.UseCaseUpdate(x, crmOwner)
	if UpdateErr != nil {
		return UpdateErr
	}
	deleteErr := usecaseCrm.UseCaseDelete(helpers.String(x.Id), crmOwner, extra.Tipo)

	if deleteErr != nil {
		return deleteErr
	}

	return nil
}
