package erp_crm

import (
	crmRepository "github.com/augusto/imersao5-esquenta-go/adapter/repository/crm"
	"github.com/augusto/imersao5-esquenta-go/services/crm"
	"github.com/augusto/imersao5-esquenta-go/usecase/crm/pedido"
	"github.com/augusto/imersao5-esquenta-go/utils"
	"github.com/augusto/imersao5-esquenta-go/utils/helpers"
	"log"
)

/*
	Implementar ainda:
	* #TODO Chamada na api do André
	* #TODO Delete na sincroniza
*/
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
