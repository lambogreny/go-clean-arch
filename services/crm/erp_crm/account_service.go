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

func AccountService(clientId string) error {
	/*
	*Crm : Todas as funções que operam no banco do CRM
	*Erp : Todas as funções que operam no banco do ERP
	 */
	log.Println("Início da transação do procedimento que leva dados de ERP para o CRM")

	////Chama a função que retorna as duas conexões
	dbCrmConn, dbErpConn, ownerCrm, connError := crm.ServicesDatabases(clientId)

	if connError != nil {
		return connError
	}

	//Criando o repositório
	repoCrm := crmRepository.NewCfrRepositoryDbErp(dbCrmConn)
	repoErp := crmRepository.NewCfrRepositoryDbErp(dbErpConn)

	//repoErp := crmRepository.NewCfrRepositoryDbErp(dbErpConn)
	//repoCrm := crmRepository.NewCfrRepositoryDbErp(dbCrmConn)

	usecaseCrm := cfr.NewProcessAccount(repoCrm)
	usecaseErp := cfr.NewProcessAccount(repoErp)

	fmt.Println(usecaseCrm)
	//fmt.Println(usecaseErp)

	//// ----------------------------------------------------------- Selecionando todas as linhas ------------------------------------------------------------//
	//
	data, err := usecaseErp.UseCaseSelect()
	if err != nil {
		utils.LogDatabase(clientId, "ACCOUNT", "S", "SELECT", true, err.Error())
		return err
	}
	fmt.Println("Dados de retorno : ", data)

	//// ----------------------------------------------------------- Percorrendo as linhas e definindo as ações ------------------------------------------------------------//

	for i, x := range data {
		log.Println("Registro : ", i, " // id : ", x.Id)
		switch helpers.String(x.Tipo) {
		case "I":
			IErr := AccountInsertWithCheck(usecaseCrm, usecaseErp, x, ownerCrm, helpers.ExtraInfo{Tipo: helpers.String(x.Tipo)})

			if IErr != nil {
				switch {
				//Para casos de duplicate key, apenas loga e continua o loop
				case strings.Contains(IErr.Error(), "duplicate key"):
					log.Println("Cai no duplicate key")
					utils.LogDatabase(clientId, "CFR", "I", helpers.String(x.Id), true, IErr.Error())
					continue
				default:
					log.Println("Cai no erro default")
					utils.LogDatabase(clientId, "CFR", "I", helpers.String(x.Id), true, IErr.Error())
					return IErr
				}
			}

		case "U":
			UErr := AccountInsertWithCheck(usecaseCrm, usecaseErp, x, ownerCrm, helpers.ExtraInfo{Tipo: helpers.String(x.Tipo)})

			if UErr != nil {
				switch {
				//Para casos de duplicate key, apenas loga e continua o loop
				case strings.Contains(UErr.Error(), "duplicate key"):
					log.Println("Cai no duplicate key")
					utils.LogDatabase(clientId, "CFR", "I", helpers.String(x.Id), true, UErr.Error())
					continue
				default:
					log.Println("Cai no erro default")
					utils.LogDatabase(clientId, "CFR", "I", helpers.String(x.Id), true, UErr.Error())
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

func AccountInsertWithCheck(usecaseCrm *cfr.ProcessAccount, usecaseErp *cfr.ProcessAccount, x cfr2.Cfr, crmOwner string, extra helpers.ExtraInfo) error {
	log.Println("CASO : ACCOUNT INSERT WITH CHECK")

	checkUpdate, err := usecaseCrm.UseCaseCheckUpdateCrm(helpers.String(x.Id), crmOwner)

	if err != nil {
		log.Println("Erro de checagem")
		return err
	}

	switch checkUpdate {
	case true:
		log.Println("Chequei o registro e cai no Update")
		UpdateErr := usecaseCrm.UseCaseUpdate(x, crmOwner)

		if UpdateErr != nil {
			return UpdateErr
		}

		deleteErr := usecaseErp.UseCaseDelete(helpers.String(x.Id), helpers.String(x.Tipo))

		if deleteErr != nil {
			return deleteErr
		}
	case false:
		log.Println("Chequei o registro e cai no insert!")

		//#TODO Fazendo o insert
		InsertErr := usecaseCrm.UseCaseInsert(x, crmOwner)

		if InsertErr != nil {
			return InsertErr
		}

		deleteErr := usecaseErp.UseCaseDelete(helpers.String(x.Id), helpers.String(x.Tipo))

		if deleteErr != nil {
			return deleteErr
		}

	}

	log.Println("O resultado do checkUpdate é : ", checkUpdate)

	return nil
}

//Não será utilizado
func AccountUpdate(usecaseCrm *cfr.ProcessAccount, usecaseErp *cfr.ProcessAccount, x cfr2.Cfr, crmOwner string, extra helpers.ExtraInfo) error {
	log.Println("CASO : ACCOUNT UPDATE")

	UpdateErr := usecaseCrm.UseCaseUpdate(x, crmOwner)

	if UpdateErr != nil {
		return UpdateErr
	}
	deleteErr := usecaseErp.UseCaseDelete(helpers.String(x.Id), helpers.String(x.Tipo))

	if deleteErr != nil {
		return deleteErr
	}

	return nil
}
