package erp_crm

import (
	"fmt"
	"log"
	"strings"

	pedido2 "github.com/augusto/imersao5-esquenta-go/entity/crm/pedido"

	crmRepository "github.com/augusto/imersao5-esquenta-go/adapter/repository/crm"
	"github.com/augusto/imersao5-esquenta-go/services/crm"
	"github.com/augusto/imersao5-esquenta-go/usecase/crm/pedido"
	"github.com/augusto/imersao5-esquenta-go/utils"
	"github.com/augusto/imersao5-esquenta-go/utils/helpers"
)

// -------------------------------------------------------- CRM -> ERP ------------------------------------------------------------//
func PedidoService(clientId string) error {
	log.Println("Inicio da transação dos procedimento de integração dos pedidos")

	//Chama a função que retorna as duas conexões
	dbCrmConn, _, ownerCrm, connError := crm.ServicesDatabases(clientId)

	if connError != nil {
		return connError
	}

	//Criando o repositório
	repoCrm := crmRepository.NewPedidoRepositoryDbErp(dbCrmConn)

	//Criando o caso de uso
	useCaseCrm := pedido.NewProcessPedido(repoCrm)

	// ----------------------------------------------------------- Selecionando e criando os payloads ------------------------------------------------------------//

	payloads, err := useCaseCrm.UseCaseSelect(ownerCrm, helpers.ExtraInfo{})

	if err != nil {
		utils.LogDatabase(clientId, "PEDIDO", "I", "SELECT PEDIDO", true, err.Error())
		return err
	}

	// -----------------------------------------------------------Com todos os payloads já montados, chamar a api ------------------------------------------------------------//

	//Recuperando a url base do cliente para a api de entrada de pedidos
	basePedidoUrl, basePedidoErr := crm.GetPedidoUrl(clientId)

	if basePedidoErr != nil {
		return basePedidoErr
	}

	callApiErr := useCaseCrm.UseCaseCallApi(payloads, helpers.ExtraInfo{Base_url: basePedidoUrl, Owner: ownerCrm})

	if callApiErr != nil {
		return callApiErr
	}

	return nil
}

// -------------------------------------------------------- ERP -> CRM : CPV ------------------------------------------------------------//

func QuoteService(clientId string) error {
	/*
	*Crm : Todas as funções que operam no banco do CRM
	*Erp : Todas as funções que operam no banco do ERP
	 */

	log.Println("Início do serviço de pegar dados da CPV e levar para a Quote")
	dbCrmConn, dbErpConn, ownerCrm, connError := crm.ServicesDatabases(clientId)

	if connError != nil {
		return connError
	}

	repoCrm := crmRepository.NewPedidoRepositoryDbErp(dbCrmConn)
	repoErp := crmRepository.NewPedidoRepositoryDbErp(dbErpConn)

	useCaseCrm := pedido.NewProcessPedido(repoCrm)
	useCaseErp := pedido.NewProcessPedido(repoErp)

	// fmt.Println(repoCrm, repoErp)

	//// ----------------------------------------------------------- Selecionando todas as linhas ------------------------------------------------------------//

	data, err := useCaseErp.UseCaseSelectCpv()

	if err != nil {
		utils.LogDatabase(clientId, "CPV", "S", "SELECT", true, err.Error())
		return err
	}

	// fmt.Println("Dados de retorno : ", data)

	//// ----------------------------------------------------------- Percorrendo as linhas e definindo as ações ------------------------------------------------------------//

	for i, x := range data {
		log.Println("Registro : ", i, " // id : ", x.Cliente)

		switch helpers.String(x.Tipo) {
		case "I":
			IErr := CpvInsertWithCheck(useCaseCrm, useCaseErp, x, ownerCrm, helpers.ExtraInfo{})

			if IErr != nil {
				switch {
				//Para casos de duplicate key, apenas loga e continua o loop
				case strings.Contains(IErr.Error(), "duplicate key"):
					log.Println("Cai no duplicate key")
					utils.LogDatabase(clientId, "CPV", "I", helpers.String(x.Numero), true, IErr.Error())
					continue
				default:
					log.Println("Cai no erro default")
					utils.LogDatabase(clientId, "CPV", "I", helpers.String(x.Numero), true, IErr.Error())
					return IErr
				}
			}
		case "U":
			UErr := CpvInsertWithCheck(useCaseCrm, useCaseErp, x, ownerCrm, helpers.ExtraInfo{})

			if UErr != nil {
				switch {
				//Para casos de duplicate key, apenas loga e continua o loop
				case strings.Contains(UErr.Error(), "duplicate key"):
					log.Println("Cai no duplicate key")
					utils.LogDatabase(clientId, "CPV", "I", helpers.String(x.Numero), true, UErr.Error())
					continue
				default:
					log.Println("Cai no erro default")
					utils.LogDatabase(clientId, "CPV", "I", helpers.String(x.Numero), true, UErr.Error())
					return UErr
				}
			}
		}

		utils.LogDatabase(clientId, "CPV", helpers.String(x.Tipo), helpers.String(x.Numero), false, "")

	}

	return nil
}

func CpvInsertWithCheck(usecaseCrm *pedido.ProcessPedido, usecaseErp *pedido.ProcessPedido, x pedido2.Cpv, crmOwner string, extra helpers.ExtraInfo) error {
	log.Println("CASO : ACCOUNT INSERT WITH CHECK")

	checkUpdate, err := usecaseCrm.UseCaseCheckUpdateCrm(helpers.String(x.Numero), crmOwner)

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

		deleteErr := usecaseErp.UseCaseDelete(helpers.String(x.Numero), helpers.String(x.Tipo))

		if deleteErr != nil {
			return deleteErr
		}

	case false:
		log.Println("Chequei o registro e cai no insert!")

		InsertErr := usecaseCrm.UseCaseInsert(x, crmOwner)

		if InsertErr != nil {
			return InsertErr
		}

		deleteErr := usecaseErp.UseCaseDelete(helpers.String(x.Numero), helpers.String(x.Tipo))

		if deleteErr != nil {
			return deleteErr
		}

	}
	return nil
}

// -------------------------------------------------------- ERP -> CRM : IPV ------------------------------------------------------------//

func QuoteItemService(clientId string) error {
	/*
	*Crm : Todas as funções que operam no banco do CRM
	*Erp : Todas as funções que operam no banco do ERP
	 */

	log.Println("Início do serviço de pegar dados da IPV e levar para a QuoteItem")
	dbCrmConn, dbErpConn, ownerCrm, connError := crm.ServicesDatabases(clientId)

	if connError != nil {
		return connError
	}

	repoCrm := crmRepository.NewPedidoRepositoryDbErp(dbCrmConn)
	repoErp := crmRepository.NewPedidoRepositoryDbErp(dbErpConn)

	useCaseCrm := pedido.NewProcessPedido(repoCrm)
	useCaseErp := pedido.NewProcessPedido(repoErp)

	fmt.Println(useCaseCrm, useCaseErp)

	//// ----------------------------------------------------------- Selecionando todas as linhas ------------------------------------------------------------//

	data, err := useCaseErp.UseCaseSelectIpv()

	if err != nil {
		utils.LogDatabase(clientId, "IPV", "S", "SELECT", true, err.Error())
		return err
	}

	fmt.Println("Dados de retorno : ", data)

	//// ----------------------------------------------------------- Percorrendo as linhas e definindo as ações ------------------------------------------------------------//

	for i, x := range data {
		log.Println("Registro : ", i, " // id : ", x.Pk)

		switch helpers.String(x.Tipo) {
		case "I":
			IErr := IpvInsertWithCheck(useCaseCrm, useCaseErp, x, ownerCrm, helpers.ExtraInfo{})

			if IErr != nil {
				switch {
				//Para casos de duplicate key, apenas loga e continua o loop
				case strings.Contains(IErr.Error(), "duplicate key"):
					log.Println("Cai no duplicate key")
					utils.LogDatabase(clientId, "CPV", "I", helpers.String(x.Pk), true, IErr.Error())
					continue
				default:
					log.Println("Cai no erro default")
					utils.LogDatabase(clientId, "CPV", "I", helpers.String(x.Pk), true, IErr.Error())
					return IErr
				}
			}
		case "U":
			UErr := IpvInsertWithCheck(useCaseCrm, useCaseErp, x, ownerCrm, helpers.ExtraInfo{})

			if UErr != nil {
				switch {
				//Para casos de duplicate key, apenas loga e continua o loop
				case strings.Contains(UErr.Error(), "duplicate key"):
					log.Println("Cai no duplicate key")
					utils.LogDatabase(clientId, "CPV", "I", helpers.String(x.Pk), true, UErr.Error())
					continue
				default:
					log.Println("Cai no erro default")
					utils.LogDatabase(clientId, "CPV", "I", helpers.String(x.Pk), true, UErr.Error())
					return UErr
				}
			}
		}

		utils.LogDatabase(clientId, "IPV", helpers.String(x.Tipo), helpers.String(x.Pk), false, "")
	}

	return nil
}

func IpvInsertWithCheck(usecaseCrm *pedido.ProcessPedido, usecaseErp *pedido.ProcessPedido, x pedido2.Ipv, crmOwner string, extra helpers.ExtraInfo) error {
	log.Println("CASO : IPV INSERT WITH CHECK")

	checkUpdate, err := usecaseCrm.UseCaseCheckUpdateCrmIpv(helpers.String(x.Pk), crmOwner)

	if err != nil {
		log.Println("Erro de checagem")
		return err
	}

	switch checkUpdate {
	case true:
		log.Println("Chequei o registro e cai no Update")
		UpdateErr := usecaseCrm.UseCaseUpdateIpv(x, crmOwner)

		if UpdateErr != nil {
			return UpdateErr
		}

		// deleteErr := usecaseErp.UseCaseDelete(helpers.String(x.Numero), helpers.String(x.Tipo))

		// if deleteErr != nil {
		// 	return deleteErr
		// }

	case false:
		log.Println("Chequei o registro e cai no insert!")

		// InsertErr := usecaseCrm.UseCaseInsert(x, crmOwner)

		// if InsertErr != nil {
		// 	return InsertErr
		// }

		// deleteErr := usecaseErp.UseCaseDelete(helpers.String(x.Numero), helpers.String(x.Tipo))

		// if deleteErr != nil {
		// 	return deleteErr
		// }

	}
	return nil
}
